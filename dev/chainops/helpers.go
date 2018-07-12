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
	"fmt"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ipfn/go-ipfn-cells"
	"github.com/ipfn/go-ipfn-keypair"
	"github.com/rootchain/go-rootchain/dev/synaptic"
)

// NewRoot - Creates root operation.
func NewRoot(ops ...cells.Cell) *cells.BinaryCell {
	return cells.Op(OpRoot, ops...)
}

// NewCID - Creates CID binary cell.
func NewCID(c *cells.CID) *cells.BinaryCell {
	if c == nil {
		return cells.Op(OpCID)
	}
	return cells.New(OpCID, c.Bytes())
}

// NewSignOperation - Signs binary cell and creates signed operation.
func NewSignOperation(op *cells.BinaryCell, pk *btcec.PrivateKey) (_ *cells.BinaryCell, err error) {
	hash := op.CID().Digest()
	if size := len(hash); size != 32 {
		return nil, fmt.Errorf("invalid hash length %d", size)
	}
	sig, err := SignBytes(hash, pk)
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
	return NewSignature(sig), nil
}

// NewSignature - Creates signature binary cell.
func NewSignature(sig []byte) *cells.BinaryCell {
	return cells.New(OpSignature, sig)
}

// NewSigned - Creates signed binary cell.
func NewSigned(op cells.Cell, signatures ...cells.Cell) *cells.BinaryCell {
	ops := append(cells.Ops(op), signatures...)
	return cells.Op(OpSigned, ops...)
}

// NewPubkey - Creates public key cell.
func NewPubkey(pubkey *btcec.PublicKey) *cells.BinaryCell {
	return NewPubkeyFromBytes(pubkey.SerializeCompressed())
}

// NewPubkeyFromBytes - Creates public key cell.
func NewPubkeyFromBytes(pubkey []byte) *cells.BinaryCell {
	return cells.New(OpPubkey, pubkey)
}

// NewAssignPower - Creates assign power operation.
func NewAssignPower(nonce cells.ID, quantity uint64, addr *cells.CID) *cells.BinaryCell {
	c := cells.Op(OpAssignPower, synaptic.Uint64(quantity), NewCID(addr))
	if nonce > 0 {
		c.AddChildren(NewNonce(nonce))
	}
	return c
}

// NewDelegatePower - Creates delegate power operation.
func NewDelegatePower(nonce cells.ID, quantity uint64, addrs ...*cells.CID) *cells.BinaryCell {
	c := cells.Op(OpDelegatePower, synaptic.Uint64(quantity))
	if len(addrs) > 0 {
		for _, addr := range addrs {
			c.AddChildren(NewCID(addr))
		}
	}
	if nonce > 0 {
		c.AddChildren(NewNonce(nonce))
	}
	return c
}

// NewPubkeyToAddr - Creates public key hash cell from public key.
func NewPubkeyToAddr(bytes []byte) *cells.BinaryCell {
	c, err := cells.SumCID(keypair.CIDPrefix, bytes)
	if err != nil {
		panic(err)
	}
	return cells.New(OpPubkeyAddr, c.Bytes())
}

// NewNonce - Creates new uint64 cell.
func NewNonce(nonce cells.ID) *cells.BinaryCell {
	return cells.New(OpNonce, nonce.Bytes())
}

// NewHeader - Creates new header cell.
// Following format is used: `[height][prev-hash][exec-hash][state-hash][timestamp]`.
func NewHeader(height uint64, prev, exec, state *cells.CID, t time.Time) *cells.BinaryCell {
	return cells.Op(OpHeader,
		synaptic.Uint64(height),
		NewCID(prev),
		NewCID(exec),
		NewCID(state),
		synaptic.Uint64(uint64(t.Unix())),
	)
}
