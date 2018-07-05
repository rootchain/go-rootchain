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

package cells

import (
	cid "gx/ipfs/QmapdYm1b22Frv3k17fqrBYTFRxwiaVJkB299Mfn33edeB/go-cid"

	multihash "github.com/multiformats/go-multihash"
	"github.com/rootchain/go-rootchain/cids"
)

// TODO - Returns empty cell.
func TODO() *BinaryCell {
	return new(BinaryCell)
}

// Ops - Returns ops.
func Ops(ops ...*BinaryCell) []*BinaryCell {
	return ops
}

// Op - Creates new binary cell.
func Op(op ID, ops ...Cell) *BinaryCell {
	return &BinaryCell{opCode: op, children: ops}
}

// RootOp - Creates new binary cell.
func RootOp(ops []Cell) *BinaryCell {
	return &BinaryCell{children: ops}
}

// New - Creates new binary cell.
func New(op ID, memory []byte) *BinaryCell {
	return &BinaryCell{
		opCode: op,
		memory: memory,
	}
}

// UnmarshalCell - Unmarshals new binary cell.
func UnmarshalCell(body []byte) (c *BinaryCell, err error) {
	c = new(BinaryCell)
	err = c.Unmarshal(body)
	return
}

// CellPrefix - Binary cell CID prefix.
var CellPrefix = cid.Prefix{
	Version:  1,
	Codec:    cids.BinaryCell,
	MhType:   multihash.KECCAK_256,
	MhLength: 32,
}
