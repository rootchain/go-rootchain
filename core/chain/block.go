// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chain

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ipfn/go-ipfn-cells"
	"github.com/rootchain/go-rootchain/dev/chainops"
	"github.com/rootchain/go-rootchain/dev/contents"
)

// Block - Chain block.
type Block struct {
	height    uint64
	timestamp time.Time
	headHash  *cells.CID
	prevHash  *cells.CID
	execHash  *cells.CID
	stateHash *cells.CID
	opsRoot   cells.MutableCell
}

// NewBlock - Creates new block.
func NewBlock(index uint64, prevCID *cells.CID, opsRoot cells.MutableCell) (*Block, error) {
	if opsRoot.OpCode() != chainops.OpRoot {
		return nil, fmt.Errorf("invalid root opcode %s", opsRoot.OpCode())
	}
	if prevCID == nil && index > 0 {
		return nil, fmt.Errorf("prev hash cannot be empty with index %d", index)
	}
	return &Block{
		height:    index,
		timestamp: time.Now(),
		prevHash:  prevCID,
		opsRoot:   opsRoot,
	}, nil
}

// Height - Returns state index.
func (block *Block) Height() uint64 {
	return block.height
}

// PrevHash - Returns previous CID.
func (block *Block) PrevHash() *cells.CID {
	return block.prevHash
}

// HeadHash - Returns head CID.
func (block *Block) HeadHash() *cells.CID {
	return block.headHash
}

// StateHash - Returns state CID.
func (block *Block) StateHash() *cells.CID {
	return block.stateHash
}

// SetStateHash - Sets state hash. Resets head hash and signatures.
func (block *Block) SetStateHash(c *cells.CID) {
	block.stateHash = c
}

// Root - Returns root operation.
func (block *Block) Root() cells.Cell {
	return block.opsRoot
}

// Exec - Adds operation to execute.
func (block *Block) Exec(op cells.Cell) {
	block.headHash = nil
	block.execHash = nil
	block.stateHash = nil
	block.opsRoot.AddChildren(op)
}

// ExecSize - Returns amount of operations.
func (block *Block) ExecSize() int {
	return block.opsRoot.ChildrenSize()
}

// IsGenesis - Returns true on zero height.
func (block *Block) IsGenesis() bool {
	return block.Height() == 0
}

// Next - Returns next state including given ops.
func (block *Block) Next(opsRoot cells.MutableCell) (*Block, error) {
	if opsRoot.ChildrenSize() == 0 {
		return nil, errors.New("cannot produce state with zero operations")
	}
	return NewBlock(block.height+1, block.headHash, opsRoot)
}

// Seal - Creates signed block.
func (block *Block) Seal() *SignedBlock {
	return NewSignedBlock(block)
}

// Header - Creates block header.
func (block *Block) Header() (*BlockHeader, error) {
	if err := block.EnsureHead(); err != nil {
		return nil, err
	}
	return &BlockHeader{
		Height:    block.height,
		Timestamp: block.timestamp,
		HeadHash:  block.headHash,
		PrevHash:  block.prevHash,
		ExecHash:  block.execHash,
		StateHash: block.stateHash,
	}, nil
}

// EnsureHead - Calculates head.
func (block *Block) EnsureHead() error {
	if block.height != 0 && block.prevHash == nil {
		return errors.New("cannot compute head w/o prev hash")
	}
	if block.stateHash == nil {
		return errors.New("cannot compute head w/o state hash")
	}
	if block.timestamp.IsZero() {
		return errors.New("cannot compute head w/o timestamp")
	}
	if block.execHash == nil {
		body, err := block.opsRoot.Marshal()
		if err != nil {
			return err
		}
		block.execHash, err = cells.SumCID(contents.OperationTriePrefix, body)
		if err != nil {
			return err
		}
	}
	if block.headHash == nil {
		body, err := chainops.NewHeaderOp(block.height, block.prevHash, block.execHash, block.stateHash, block.timestamp).Marshal()
		if err != nil {
			return err
		}
		block.headHash, err = cells.SumCID(contents.HeaderPrefix, body)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON - Marshals state to JSON.
func (block *Block) MarshalJSON() ([]byte, error) {
	header, err := block.Header()
	if err != nil {
		return nil, err
	}
	return json.Marshal(struct {
		BlockHeader *BlockHeader         `json:"header,omitempty"`
		ExecOps     []*cells.CellPrinter `json:"exec_ops,omitempty"`
	}{
		BlockHeader: header,
		ExecOps:     cells.NewChildrenPrinter(block.opsRoot),
	})
}
