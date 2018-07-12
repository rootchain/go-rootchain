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

	"github.com/ipfn/go-ipfn-cells"
)

// BlockHeader - Block header structure.
type BlockHeader struct {
	// Height - Block height.
	Height uint64 `json:"height,omitempty"`

	// Time - Block time.
	Time time.Time `json:"timestamp,omitempty"`

	// Head - Head content ID.
	Head *cells.CID `json:"head_hash,omitempty"`

	// Prev - Previous block head hash.
	Prev *cells.CID `json:"prev_hash,omitempty"`

	// Exec - Block execution hash.
	Exec *cells.CID `json:"exec_hash,omitempty"`

	// State - State trie hash.
	State *cells.CID `json:"state_hash,omitempty"`

	// Signed - Signed hash.
	Signed *cells.CID `json:"signed_hash,omitempty"`
}

// NewBlockHeader - Creates new state header structure.
func NewBlockHeader(index uint64, prevHash *cells.CID, execCID *cells.CID) (hdr *BlockHeader, err error) {
	if prevHash == nil && index > 0 {
		return nil, fmt.Errorf("prev hash cannot be empty with index %d", index)
	}
	hdr = &BlockHeader{
		Height: index,
		Time:   time.Now(),
		Prev:   prevHash,
	}
	err = hdr.SetExec(execCID)
	if err != nil {
		return
	}
	// BUG(crackcomm): fucking state trie hash?!
	hdr.State, err = cells.SumCID(StateTriePrefix, hdr.Head.Bytes())
	if err != nil {
		return
	}
	return
}

// SetExec - Sets exec and head hash.
func (hdr *BlockHeader) SetExec(c *cells.CID) error {
	hdr.Exec = c
	hdr.Signed = nil
	return hdr.calcHead()
}

func (hdr *BlockHeader) calcHead() (err error) {
	hdr.Head, err = NewHeadCID(hdr)
	return
}
