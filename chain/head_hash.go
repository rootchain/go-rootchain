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
	"github.com/rootchain/go-rootchain/cells"
	"github.com/rootchain/go-rootchain/cells/chainops"
	"github.com/rootchain/go-rootchain/cells/synaptic"
)

// NewHeadCID - Computes header cid.
func NewHeadCID(hdr *Header) (_ *cells.CID, err error) {
	body, err := cells.Marshal(hdr.Cell())
	if err != nil {
		return
	}
	return cells.SumCID(HeaderPrefix, body)
}

// Cell - Creates header binary cell.
func (hdr *Header) Cell() *cells.BinaryCell {
	return cells.Op(chainops.OpHeader,
		synaptic.Uint64(hdr.Height),
		synaptic.Uint64(uint64(hdr.Time.Unix())),
		chainops.CID(hdr.Prev),
		chainops.CID(hdr.Exec),
	)
}
