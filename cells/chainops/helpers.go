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
	"github.com/btcsuite/btcd/btcec"
	base32check "github.com/ipfn/go-base32check"
	"github.com/ipfn/go-ipfn-keypair"
	multihash "github.com/multiformats/go-multihash"
	"github.com/rootchain/go-rootchain/cells"
	"github.com/rootchain/go-rootchain/cells/synaptic"
	"github.com/rootchain/go-rootchain/dev/address"
)

// ID - Creates new uint64 cell.
func ID(num cells.ID) *cells.BinaryCell {
	return cells.New(OpID, num.Bytes())
}

// IDFromString - Creates new uint64 cell.
func IDFromString(body string) *cells.BinaryCell {
	return ID(cells.NewIDFromString(body))
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
func MustParseID(str string) (c *cells.BinaryCell) {
	c, err := ParseID(str)
	if err != nil {
		panic(err)
	}
	return
}

// ParseAddress - Parses short address from string.
func ParseAddress(src string) (_ *cells.BinaryCell, err error) {
	addr, err := address.ParseAddress(src)
	if err != nil {
		return
	}
	bytes, err := addr.Marshal()
	if err != nil {
		return
	}
	return cells.New(OpAddress, bytes), nil
}

// MustParseAddress - Parses short address or panics.
func MustParseAddress(src string) (c *cells.BinaryCell) {
	c, err := ParseAddress(src)
	if err != nil {
		panic(err)
	}
	return
}

// CID - Creates CID binary cell.
func CID(c *cells.CID) (_ *cells.BinaryCell) {
	if c == nil {
		return cells.Op(OpCID)
	}
	return cells.New(OpCID, c.Bytes())
}

// Multihash - Creates multihash binary cell.
func Multihash(mh multihash.Multihash) *cells.BinaryCell {
	return cells.New(OpMultihash, []byte(mh))
}

// Sign - Signs binary cell and creates signed operation.
func Sign(op *cells.BinaryCell, pk *btcec.PrivateKey) (_ *cells.BinaryCell, err error) {
	body, err := cells.Marshal(op)
	if err != nil {
		return
	}
	sig, err := SignBytes(body, pk)
	if err != nil {
		return
	}
	return cells.Op(OpSigned, op, sig), nil
}

// SignBytes - Signs bytes and creates signature operation.
func SignBytes(body []byte, pk *btcec.PrivateKey) (_ *cells.BinaryCell, err error) {
	sig, err := btcec.SignCompact(btcec.S256(), pk, body, false)
	if err != nil {
		return
	}
	return Signature(sig), nil
}

// Signature - Creates signature binary cell.
func Signature(sig []byte) (_ *cells.BinaryCell) {
	return cells.New(OpSignature, sig)
}

// Signed - Creates signed binary cell.
func Signed(op *cells.BinaryCell, signatures ...cells.Cell) *cells.BinaryCell {
	ops := append(cells.Ops(op), signatures...)
	return cells.Op(OpSigned, ops...)
}

// Pubkey - Creates public key cell.
func Pubkey(pubkey *btcec.PublicKey) (c *cells.BinaryCell) {
	return PubkeyBytes(pubkey.SerializeCompressed())
}

// PubkeyBytes - Creates public key cell.
func PubkeyBytes(pubkey []byte) (c *cells.BinaryCell) {
	return cells.New(OpPubkey, pubkey)
}

// Genesis - Creates genesis operation.
func Genesis() (c *cells.BinaryCell) {
	return cells.Op(OpGenesis)
}

// AssignPower - Creates assign power operation.
func AssignPower(nonce cells.ID, quantity uint64, pubkey []byte) (c *cells.BinaryCell) {
	return cells.Op(OpAssignPower,
		// cells.New(OpNonce, nonce.Bytes()),
		synaptic.Uint64(quantity),
		PubkeyBytes(pubkey))
}

// AssignPowerAddr - Creates assign power operation.
func AssignPowerAddr(nonce cells.ID, quantity uint64, addr *cells.CID) (c *cells.BinaryCell) {
	return cells.Op(OpAssignPower,
		// cells.New(OpNonce, nonce.Bytes()),
		synaptic.Uint64(quantity),
		CID(addr))
}

// DelegatePower - Creates delegate power operation.
func DelegatePower(nonce cells.ID, quantity uint64, pubkeys ...[]byte) (c *cells.BinaryCell) {
	c = cells.Op(OpAssignPower,
		// cells.New(OpNonce, nonce.Bytes()),
		synaptic.Uint64(quantity))
	if len(pubkeys) > 0 {
		for _, pubkey := range pubkeys {
			c.AddChild(PubkeyBytes(pubkey))
		}
	}
	return
}

// PubkeyToAddr - Creates public key hash cell from public key.
func PubkeyToAddr(bytes []byte) *cells.BinaryCell {
	c, err := cells.SumCID(keypair.CIDPrefix, bytes)
	if err != nil {
		panic(err)
	}
	return cells.New(OpPubkeyAddr, c.Bytes())
}
