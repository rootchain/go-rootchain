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
	body, err := Marshal(cell)
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
	cell.reset()
}

// SetOpCode - Sets operation ID.
func (cell *BinaryCell) SetOpCode(opCode ID) {
	cell.opCode = opCode
	cell.reset()
}

// SetMemory - Set operation memory.
func (cell *BinaryCell) SetMemory(memory []byte) {
	cell.memory = memory
	cell.reset()
}

// SetChildren - Set operation children.
func (cell *BinaryCell) SetChildren(children []Cell) {
	cell.children = children
	cell.reset()
}

// Checksum - Computes marshalled xxhash64 of cell content id.
func (cell *BinaryCell) Checksum() (_ ID, err error) {
	return NewID(cell.CID().Bytes()), nil
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

func (cell *BinaryCell) reset() {
	cell.cid = nil
	cell.body = nil
}
