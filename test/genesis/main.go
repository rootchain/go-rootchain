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

package main

import (
	"io/ioutil"

	"github.com/ipfn/go-ipfn-cmd-util/logger"
	"github.com/ipfn/ipfn/go/chain"
	"github.com/ipfn/ipfn/go/keypair"
	"github.com/ipfn/ipfn/go/opcode"
	"github.com/ipfn/ipfn/go/opcode/chainops"
	"github.com/ipfn/ipfn/go/opcode/synaptic"
	"github.com/ipfn/ipfn/go/wallet"
)

func main() {
	w := wallet.NewDefault()

	acc1, err := w.DeriveKeyPath("default/x/test", []byte("123"))
	if err != nil {
		panic(err)
	}

	privkey1, _ := acc1.ECPrivKey()

	acc2, err := w.DeriveKeyPath("default/x/test2", []byte("123"))
	if err != nil {
		panic(err)
	}

	privkey2, _ := acc2.ECPrivKey()

	acc3, err := w.DeriveKeyPath("default/x/test3", []byte("123"))
	if err != nil {
		panic(err)
	}

	// privkey3, _ := acc3.ECPrivKey()

	// pub, _ := acc1.ECPubKey()
	// s, _ := chainops.Sign(pk, txn)
	// key := newKey()

	// signedGenesis := signedOp(genesisOp, key)

	allocOp := chainops.AssignPower(0, 1e7, acc1.Serialize())
	alloc2Op := chainops.AssignPower(1, 1e7, acc2.Serialize())
	alloc3Op := chainops.AssignPower(2, 1e7, acc3.Serialize())

	delegateOp, _ := chainops.Sign(opcode.Op(chainops.OpDelegatePower, synaptic.Uint64(1e7)), privkey1)
	delegate2Op, _ := chainops.Sign(opcode.Op(chainops.OpDelegatePower, synaptic.Uint64(1e7)), privkey2)

	c, _ := opcode.SumCID(chain.HeaderPrefix, []byte("randomdata"))

	state, err := chain.NewState(0, nil, opcode.Ops(
		chainops.Genesis(), allocOp, alloc2Op, alloc3Op, delegateOp, delegate2Op)) //, claimOp))
	if err != nil {
		logger.Error(err)
	}

	state.Sign(privkey1)
	state.Sign(privkey2)

	logger.PrintJSON(state) //.Head().String())

	body, err := state.Root().Marshal()
	if err != nil {
		panic(err)
	}

	if err = ioutil.WriteFile("block.cell", body, 0666); err != nil {
		panic(err)
	}

	signature := state.Signatures()[1]
	if err = ioutil.WriteFile("signature.cell", signature, 0666); err != nil {
		panic(err)
	}

	// transfer := &opcode.BinaryCell{
	// 	OpCode: chainops.OpTransfer,
	// 	Children: []*opcode.BinaryCell{
	// 		synaptic.Uint64(1e6),                    // quantity
	// 		chainops.PubkeyToAddr(acc1.Serialize()), // receiver
	// 		// chainops.MustParseAddress("beqpdfdhq87dkncb"),       // from
	// 		// chainops.MustParseAddress("bnx37fk4wmxur3j0puapwv"), // to
	// 	},
	// }
	// signedTxn, _ := chainops.Sign(transfer, privkey3)

	// state, err = state.Next(opcode.Ops(signedTxn))
	// if err != nil {
	// 	logger.Error(err)
	// }
	// state.Sign(privkey1)
	// state.Sign(privkey2)

	// logger.PrintJSON(state)

	// state.Sign(privkey1)
	// state.Sign(privkey2)

	// claimOp := &opcode.BinaryCell{
	// 	OpCode: chainops.OpClaim,
	// 	Children: []*opcode.BinaryCell{
	// 		chainops.CID(state.Signed()),
	// 		chainops.ID(0),
	// 	},
	// }
	// signedClaim, _ := chainops.Sign(claimOp, privkey1)

	// state, err = state.Next(opcode.Ops(signedClaim))
	// if err != nil {
	// 	logger.Error(err)
	// }

	// state.Sign(privkey1)
	// state.Sign(privkey2)

	// logger.PrintJSON(state) //.Head().String())

	c, _ = acc1.CID()
	logger.Print(c.String())

	// c, _ = genesisOp.CID()
	// logger.Print(c.String())
}

func newKey() *keypair.KeyPair {
	acc, err := wallet.NewDefault().DeriveKey(wallet.NewKeyPath("default", "acc1", true), []byte("123"))
	if err != nil {
		panic(err)
	}
	return acc
}
