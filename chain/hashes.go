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
	mh "gx/ipfs/QmPnFwZ2JXKnXgMw8CdBPxn7FWh6LLdjUjxV1fKHuJnkr8/go-multihash"
	cid "gx/ipfs/QmapdYm1b22Frv3k17fqrBYTFRxwiaVJkB299Mfn33edeB/go-cid"

	"github.com/ipfn/go-ipfn-cells"
	"github.com/rootchain/go-rootchain/dev/chainops"
	"github.com/rootchain/go-rootchain/dev/contents"
	"github.com/rootchain/go-rootchain/dev/synaptic"
)

var (
	// HeaderPrefix - Header CID prefix.
	HeaderPrefix = cid.Prefix{
		Version:  1,
		Codec:    contents.ChainHeader,
		MhType:   mh.KECCAK_256,
		MhLength: 32,
	}

	// SignedPrefix - Signed header CID prefix.
	SignedPrefix = cid.Prefix{
		Version:  1,
		Codec:    contents.ChainSigned,
		MhType:   mh.KECCAK_256,
		MhLength: 32,
	}

	// StateTriePrefix - State trie CID prefix.
	StateTriePrefix = cid.Prefix{
		Version:  1,
		Codec:    cid.EthStateTrie,
		MhType:   mh.KECCAK_256,
		MhLength: 32,
	}
)

// NewHeadCID - Computes header cid.
func NewHeadCID(hdr *BlockHeader) (_ *cells.CID, err error) {
	body, err := hdr.Cell().Marshal()
	if err != nil {
		return
	}
	return cells.SumCID(HeaderPrefix, body)
}

// Cell - Creates header binary cell.
func (hdr *BlockHeader) Cell() *cells.BinaryCell {
	return cells.Op(chainops.OpHeader,
		synaptic.Uint64(hdr.Height),
		synaptic.Uint64(uint64(hdr.Time.Unix())),
		chainops.NewCID(hdr.Prev),
		chainops.NewCID(hdr.Exec),
		chainops.NewCID(hdr.State),
	)
}
