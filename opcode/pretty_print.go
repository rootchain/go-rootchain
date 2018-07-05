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

package opcode

import (
	"bytes"
	"fmt"
	"strings"

	cid "gx/ipfs/QmapdYm1b22Frv3k17fqrBYTFRxwiaVJkB299Mfn33edeB/go-cid"

	"github.com/gogo/protobuf/proto"
)

// CellPrinter - Cell pretty printer.
type CellPrinter struct {
	*BinaryCell
}

// NewPrinter - Creates new cell pretty printer.
func NewPrinter(cell *BinaryCell) *CellPrinter {
	return &CellPrinter{BinaryCell: cell}
}

// NewPrinters - Creates new cell pretty printer.
func NewPrinters(cells []*BinaryCell) (res []*CellPrinter) {
	for _, cell := range cells {
		res = append(res, NewPrinter(cell))
	}
	return
}

// MarshalJSON - Marshals cell as JSON.
func (p *CellPrinter) MarshalJSON() (_ []byte, err error) {
	return prettyPrint(p.BinaryCell), nil
}

// String - Prints to string.
func (p *CellPrinter) String() string {
	return string(prettyPrint(p.BinaryCell))
}

func prettyPrint(cell *BinaryCell) (_ []byte) {
	buff := bytes.NewBuffer(nil)
	buff.WriteByte('"')
	writeStringScript(cell, buff)
	buff.WriteByte('"')
	return buff.Bytes()
}

func writeStringScript(cell *BinaryCell, buff *bytes.Buffer) {
	buff.WriteString(strings.ToUpper(fmt.Sprintf("OP_%s", cell.OpCode)))
	if len(cell.Memory) > 0 {
		writePrettyMemory(cell, buff)
	}
	if len(cell.Children) == 0 {
		return
	}
	buff.WriteString(" [ ")
	for _, child := range cell.Children {
		writeStringScript(child, buff)
		buff.WriteByte(' ')
	}
	buff.WriteString("]")
}

func writePrettyMemory(cell *BinaryCell, buff *bytes.Buffer) {
	buff.WriteByte(' ')
	switch cell.OpCode {
	case 31, 62, 75: // uint64 or id or nonce
		i, _ := proto.DecodeVarint(cell.Memory)
		buff.WriteString(fmt.Sprintf("%d", i))
	case 63, 70: // cid or pubkey addr
		c, _ := cid.Cast(cell.Memory)
		buff.WriteString(c.String())
	default:
		buff.WriteString(fmt.Sprintf("0x%x", cell.Memory))
	}
}
