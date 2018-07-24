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
	"time"

	"github.com/ipfn/go-ipfn-cells"
)

// BlockHeader - Block header.
type BlockHeader struct {
	// Height - Block height.
	Height uint64 `json:"height,omitempty"`
	// Timestamp - Block time.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// HeadHash - Head content ID.
	HeadHash *cells.CID `json:"head_hash,omitempty"`
	// PrevHash - Previous block head hash.
	PrevHash *cells.CID `json:"prev_hash,omitempty"`
	// ExecHash - Block execution hash.
	ExecHash *cells.CID `json:"exec_hash,omitempty"`
	// StateHash - State trie hash.
	StateHash *cells.CID `json:"state_hash,omitempty"`
	// SignedHash - Signed head hash.
	SignedHash *cells.CID `json:"signed_hash,omitempty"`
}
