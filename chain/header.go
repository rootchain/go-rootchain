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
	"errors"
	"fmt"
	"time"

	"github.com/rootchain/go-rootchain/dev/chainops"

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
func NewBlockHeader(index uint64, prevCID *cells.CID, execCID *cells.CID) (hdr *BlockHeader, err error) {
	if prevCID == nil && index > 0 {
		return nil, fmt.Errorf("prev hash cannot be empty with index %d", index)
	}
	return &BlockHeader{
		Height: index,
		Time:   time.Now(),
		Prev:   prevCID,
		Exec:   execCID,
	}, nil
}

// SetExecHash - Sets exec hash. Resets exec, state and head hash.
func (hdr *BlockHeader) SetExecHash(c *cells.CID) {
	hdr.Exec = c
	hdr.State = nil
	hdr.Head = nil
	hdr.Signed = nil
}

// SetStateHash - Sets state hash. Resets head hash.
func (hdr *BlockHeader) SetStateHash(c *cells.CID) {
	hdr.Head = nil
	hdr.State = c
	hdr.Signed = nil
}

// EnsureHead - Calculates head.
func (hdr *BlockHeader) EnsureHead() error {
	if hdr.Height != 0 && hdr.Prev == nil {
		return errors.New("cannot compute head w/o prev hash")
	}
	if hdr.Exec == nil {
		return errors.New("cannot compute head w/o exec hash")
	}
	if hdr.State == nil {
		return errors.New("cannot compute head w/o state hash")
	}
	if hdr.Time.IsZero() {
		return errors.New("cannot compute head w/o timestamp")
	}
	if hdr.Head == nil {
		hdr.Head = chainops.NewHeaderOp(hdr.Height, hdr.Prev, hdr.Exec, hdr.State, hdr.Time).CID()
	}
	return nil
}
