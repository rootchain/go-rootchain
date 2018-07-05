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
	"encoding/json"

	"github.com/gogo/protobuf/proto"
)

// Cell - Operation cell interface.
type Cell interface {
	// CID - Operation CID.
	CID() *CID

	// OpCode - Operation ID.
	OpCode() ID

	// Memory - Operation memory.
	Memory() []byte

	// Child - Child cell by index.
	Child(int) Cell

	// ChildrenSize - Amount of children.
	ChildrenSize() int
}

// MutableCell - Mutable cell interface.
type MutableCell interface {
	Cell

	// AddChild - Adds children.
	AddChild(Cell)

	// SetOpCode - Sets operation ID.
	SetOpCode(ID)

	// SetMemory - Set operation memory.
	SetMemory([]byte)

	// SetChildren - Set operation children.
	SetChildren([]Cell)
}

// BinaryCell - Binary representation of cell.
type BinaryCell struct {
	opCode   ID     `json:"op,omitempty"`
	memory   []byte `json:"value,omitempty"`
	children []Cell `json:"ops,omitempty"`

	cid  *CID
	body []byte
}

// CID - Computes marshaled cell cid.
func (cell *BinaryCell) CID() (_ *CID) {
	if cell.cid != nil {
		return cell.cid
	}
	body, err := cell.Marshal()
	if err != nil {
		panic(err)
	}
	cell.cid, err = SumCID(CellPrefix, body)
	if err != nil {
		panic(err)
	}
	return cell.cid
}

// OpCode - Operation ID.
func (cell *BinaryCell) OpCode() ID {
	return cell.opCode
}

// Memory - Operation memory.
func (cell *BinaryCell) Memory() []byte {
	return cell.memory
}

// Child - Child cell by index.
func (cell *BinaryCell) Child(n int) Cell {
	if len(cell.children) <= n {
		return nil
	}
	return cell.children[n]
}

// ChildrenSize - Amount of children.
func (cell *BinaryCell) ChildrenSize() int {
	return len(cell.children)
}

// AddChild - Appends new children operation.
func (cell *BinaryCell) AddChild(child Cell) {
	cell.children = append(cell.children, child)
}

// SetOpCode - Sets operation ID.
func (cell *BinaryCell) SetOpCode(opCode ID) {
	cell.opCode = opCode
}

// SetMemory - Set operation memory.
func (cell *BinaryCell) SetMemory(memory []byte) {
	cell.memory = memory
}

// SetChildren - Set operation children.
func (cell *BinaryCell) SetChildren(children []Cell) {
	cell.children = children
}

// MarshalJSON - Marshals cell as JSON.
func (cell *BinaryCell) MarshalJSON() (_ []byte, err error) {
	type jsonCell struct {
		OpCode   ID             `json:"op,omitempty"`
		Memory   []byte         `json:"value,omitempty"`
		Children []*CellPrinter `json:"ops,omitempty"`
	}
	children := make([]*CellPrinter, len(cell.children))
	for n, child := range cell.children {
		children[n] = NewPrinter(child)
	}
	return json.Marshal(jsonCell{
		OpCode:   cell.opCode,
		Memory:   cell.memory,
		Children: children,
	})
}

// Marshal - Marshals cell as byte array.
func (cell *BinaryCell) Marshal() (_ []byte, err error) {
	if cell.body != nil {
		return cell.body, nil
	}
	buff := proto.NewBuffer(nil)
	err = marshal(cell, buff)
	if err != nil {
		return
	}
	cell.body = buff.Bytes()
	return cell.body, nil
}

// Unmarshal - Unmarshals cell from byte array.
func (cell *BinaryCell) Unmarshal(body []byte) (err error) {
	cell.body = body
	return unmarshal(cell, proto.NewBuffer(body))
}

// Checksum - Computes marshalled xxhash64 of cell content id.
func (cell *BinaryCell) Checksum() (_ ID, err error) {
	// hash, err := cell.CID()
	// if err != nil {
	// 	return
	// }
	// return NewID(hash.Bytes()), nil
	return NewID(cell.CID().Bytes()), nil
}
