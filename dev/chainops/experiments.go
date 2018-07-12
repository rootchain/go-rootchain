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

package chainops

import (
	base32check "github.com/ipfn/go-base32check"
	cells "github.com/ipfn/go-ipfn-cells"
	multihash "github.com/multiformats/go-multihash"
	"github.com/rootchain/go-rootchain/dev/synaptic"
)

// NewMultihash - Creates multihash binary cell.
func NewMultihash(mh multihash.Multihash) *cells.BinaryCell {
	return cells.New(OpMultihash, []byte(mh))
}

// NewID New- Creates new uint64 cell.
func NewID(num cells.ID) *cells.BinaryCell {
	return cells.New(OpID, num.Bytes())
}

// NewIDFromString - Creates new uint64 cell.
func NewIDFromString(body string) *cells.BinaryCell {
	return NewID(cells.NewIDFromString(body))
}

// ParseID - Creates new uint64 cell from string by parsing it.
func ParseID(src string) (_ *cells.BinaryCell, err error) {
	id, err := base32check.CheckDecodeString(src)
	if err != nil {
		return
	}
	return cells.New(synaptic.OpUint64, id), nil
}

// MustParseID - Creates new uint64 cell from string.
func MustParseID(str string) *cells.BinaryCell {
	c, err := ParseID(str)
	if err != nil {
		panic(err)
	}
	return c
}

// // ParseAddress - Parses short address from string.
// func ParseAddress(src string) (_ *cells.BinaryCell, err error) {
// 	addr, err := address.ParseAddress(src)
// 	if err != nil {
// 		return
// 	}
// 	bytes, err := addr.Marshal()
// 	if err != nil {
// 		return
// 	}
// 	return cells.New(OpAddress, bytes), nil
// }

// // MustParseAddress - Parses short address or panics.
// func MustParseAddress(src string) *cells.BinaryCell {
// 	c, err := ParseAddress(src)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return c
// }
