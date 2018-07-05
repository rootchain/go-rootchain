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
	"encoding/json"

	"github.com/ipfn/go-ipfn-cmd-util/logger"
	"github.com/ipfn/ipfn/go/opcode"
	"github.com/ipfn/ipfn/go/opcode/chainops"
	"github.com/ipfn/ipfn/go/opcode/synaptic"
	"github.com/ipfn/ipfn/go/wallet"
)

var (
	txn = &opcode.BinaryCell{
		OpCode: chainops.OpTransfer,
		Children: []*opcode.BinaryCell{
			chainops.MustParseAddress("beqpdfdhq87dkncb"),       // from
			chainops.MustParseAddress("bnx37fk4wmxur3j0puapwv"), // to
			synaptic.Uint64(1e6),                                // quantity
		},
	}

	// sig( cid( txn ) )
	// signature of above transaction
	signature = &opcode.BinaryCell{
		OpCode: chainops.OpSignature,
		Children: []*opcode.BinaryCell{
			synaptic.MustParseBigInt("7f3aa6a10adbc7bcfe554c6a8d1e0e1870518d53c011d166c56350038254f3ba"), // r
			synaptic.MustParseBigInt("0720fdfe61c98924e04651b21beb972564754404f2b58a5686c02ecad0192316"), // s
			synaptic.Uint64(123), // v
		},
	}
)

func main() {
	w := wallet.NewDefault()

	acc1, err := w.DeriveKey(wallet.NewKeyPath("default", "acc1", true), []byte("123"))
	if err != nil {
		panic(err)
	}

	pk, _ := acc1.ECPrivKey()
	sig, _ := pk.Sign([]byte("test"))

	logger.Printf("r: %x", sig.R.Bytes())
	logger.Printf("s: %x", sig.S.Bytes())

	signedTxn, _ := chainops.Sign(pk, txn)

	b, _ := json.MarshalIndent(*signedTxn, "", "  ")
	logger.Printf("%s", b)
	b, _ = signedTxn.Marshal()
	logger.Printf("%x", b)
}
