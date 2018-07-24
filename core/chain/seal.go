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

	"github.com/btcsuite/btcd/btcec"
	"github.com/ipfn/go-ipfn-cells"
	"github.com/rootchain/go-rootchain/dev/chainops"
	"github.com/rootchain/go-rootchain/dev/contents"
)

// SignedBlock - Signed chain block.
type SignedBlock struct {
	rootOp     cells.MutableCell
	signedHash *cells.CID
	signatures [][]byte

	*Block
}

// NewSignedBlock - Creates new signed block.
func NewSignedBlock(block *Block) *SignedBlock {
	return &SignedBlock{
		rootOp: cells.Op(chainops.OpSigned, block.opsRoot),
		Block:  block,
	}
}

// SignedHash - Returns signed head CID.
func (block *SignedBlock) SignedHash() *cells.CID {
	return block.signedHash
}

// Signatures - Returns state signatures.
func (block *SignedBlock) Signatures() [][]byte {
	return block.signatures
}

// Sign - Signs state with given private key.
// Computes new signed header hash.
func (block *SignedBlock) Sign(key *btcec.PrivateKey) (err error) {
	err = block.Block.EnsureHead()
	if err != nil {
		return
	}
	sigOp, err := chainops.NewSignatureOp(block.headHash.Bytes(), key)
	if err != nil {
		return
	}
	block.signedHash = nil
	block.signatures = append(block.signatures, sigOp.Memory())
	block.rootOp.AddChildren(sigOp)
	return
}

// MarshalJSON - Marshals state to JSON.
func (block *SignedBlock) MarshalJSON() ([]byte, error) {
	header, err := block.Header()
	if err != nil {
		return nil, err
	}
	return json.Marshal(struct {
		BlockHeader *BlockHeader         `json:"header,omitempty"`
		ExecOps     []*cells.CellPrinter `json:"exec_ops,omitempty"`
		Signatures  [][]byte             `json:"signatures,omitempty"`
	}{
		BlockHeader: header,
		ExecOps:     cells.NewChildrenPrinter(block.opsRoot),
		Signatures:  block.signatures,
	})
}

// Header - Creates block header with signed hash.
func (block *SignedBlock) Header() (header *BlockHeader, err error) {
	header, err = block.Block.Header()
	if err != nil {
		return
	}
	if block.signedHash != nil || len(block.signatures) == 0 {
		return
	}
	body, err := block.rootOp.Marshal()
	if err != nil {
		return
	}
	header.SignedHash, err = cells.SumCID(contents.SignedPrefix, body)
	return
}
