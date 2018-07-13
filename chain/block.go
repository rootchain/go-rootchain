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

	"github.com/btcsuite/btcd/btcec"
)

// Block - Chain block structure.
type Block struct {
	header     *BlockHeader
	opsRoot    cells.MutableCell
	sigRoot    cells.MutableCell
	signatures [][]byte
}

// NewBlock - Creates new block structure.
func NewBlock(index uint64, prevHash *cells.CID, opsRoot cells.MutableCell) (block *Block, err error) {
	if opsRoot.OpCode() != chainops.OpRoot {
		return nil, fmt.Errorf("invalid root opcode %s", opsRoot.OpCode())
	}
	if prevHash == nil && index > 0 {
		return nil, fmt.Errorf("prev hash cannot be empty with index %d", index)
	}
	execCID := opsRoot.CID()
	header, err := NewBlockHeader(index, prevHash, execCID)
	if err != nil {
		return
	}
	block = &Block{
		header:  header,
		opsRoot: opsRoot,
		sigRoot: cells.Op(chainops.OpSigned, opsRoot),
	}
	err = block.calcHeader()
	if err != nil {
		return
	}
	return
}

// Prev - Returns previous CID.
func (block *Block) Prev() *cells.CID {
	return block.header.Prev
}

// Height - Returns state index.
func (block *Block) Height() uint64 {
	return block.header.Height
}

// Head - Returns head CID.
func (block *Block) Head() *cells.CID {
	return block.header.Head
}

// Signed - Returns signed head CID.
func (block *Block) Signed() *cells.CID {
	return block.header.Signed
}

// Root - Returns root operation.
func (block *Block) Root() cells.Cell {
	return block.opsRoot
}

// Signatures - Returns state signatures.
func (block *Block) Signatures() [][]byte {
	return block.signatures
}

// Exec - Adds operation to execute.
func (block *Block) Exec(op cells.Cell) {
	block.opsRoot.AddChildren(op)
	block.reset()
}

// NOps - Returns amount of operations.
func (block *Block) NOps() int {
	return block.opsRoot.ChildrenSize()
}

// IsGenesis - Returns true on zero height.
func (block *Block) IsGenesis() bool {
	return block.Height() == 0
}

// SetStateHash - Sets state hash. Resets head hash.
func (block *Block) SetStateHash(c *cells.CID) {
	block.header.SetStateHash(c)
}

// Next - Returns next state including given ops.
func (block *Block) Next(root cells.MutableCell) (*Block, error) {
	if root.ChildrenSize() == 0 {
		return nil, errors.New("cannot produce state with zero operations")
	}
	return NewBlock(block.Height()+1, block.Head(), root)
}

// Sign - Signs state with given private key.
// Computes new signed header hash.
func (block *Block) Sign(key *btcec.PrivateKey) (_ cells.Cell, err error) {
	block.header.EnsureHead()
	sigOp, err := chainops.NewSignatureOp(block.Head().Bytes(), key)
	if err != nil {
		return
	}
	block.sigRoot.AddChildren(sigOp)
	block.signatures = append(block.signatures, sigOp.Memory())
	block.header.Signed = nil
	return
}

type stateJSON struct {
	BlockHeader *BlockHeader         `json:"header,omitempty"`
	ExecOps     []*cells.CellPrinter `json:"exec_ops,omitempty"`
	Signatures  [][]byte             `json:"signatures,omitempty"`
}

// MarshalJSON - Marshals state to JSON.
func (block *Block) MarshalJSON() ([]byte, error) {
	if err := block.calcHeader(); err != nil {
		return nil, err
	}
	return json.Marshal(stateJSON{
		BlockHeader: block.header,
		ExecOps:     cells.NewChildrenPrinter(block.opsRoot),
		Signatures:  block.signatures,
	})
}

func (block *Block) reset() {
	block.header.Time = time.Now()
	block.header.Head = nil
	block.header.Exec = nil
	block.header.Signed = nil
	block.signatures = nil
}

func (block *Block) calcHeader() (err error) {
	if block.header.Exec == nil {
		block.header.SetExecHash(block.opsRoot.CID())
	}
	if block.header.Head == nil {
		if err := block.header.EnsureHead(); err != nil {
			return err
		}
	}
	if block.header.Signed == nil && len(block.signatures) > 0 {
		body, err := block.sigRoot.Marshal()
		if err != nil {
			return err
		}
		block.header.Signed, err = cells.SumCID(SignedPrefix, body)
		if err != nil {
			return err
		}
	}
	return
}
