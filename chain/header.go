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
	"fmt"
	"time"

	"github.com/ipfn/ipfn/go/opcode"
)

// Header - State header structure.
type Header struct {
	// Height - State height.
	Height uint64 `json:"height,omitempty"`

	// Time - State time.
	Time time.Time `json:"timestamp,omitempty"`

	// Head - Head content ID.
	Head *opcode.CID `json:"head_hash,omitempty"`

	// Prev - Previous state hash.
	Prev *opcode.CID `json:"prev_hash,omitempty"`

	// Exec - State execution hash.
	Exec *opcode.CID `json:"exec_hash,omitempty"`

	// State - State trie hash.
	State *opcode.CID `json:"state_hash,omitempty"`

	// Signed - Signed hash.
	Signed *opcode.CID `json:"signed_hash,omitempty"`
}

// NewHeader - Creates new state header structure.
func NewHeader(index uint64, prevHash *opcode.CID, execCID *opcode.CID) (hdr *Header, err error) {
	if prevHash == nil && index > 0 {
		return nil, fmt.Errorf("prev hash cannot be empty with index %d", index)
	}
	hdr = &Header{
		Height: index,
		Time:   time.Now(),
		Exec:   execCID,
		Prev:   prevHash,
	}
	hdr.Head, err = NewHeadCID(hdr)
	if err != nil {
		return nil, err
	}
	// BUG(crackcomm): fucking state trie hash?!
	hdr.State, err = opcode.SumCID(StateTriePrefix, hdr.Head.Bytes())
	if err != nil {
		return
	}
	return
}
