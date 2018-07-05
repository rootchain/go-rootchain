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
			chainops.MustParseAddress("bmh8ctyulazuftasplwfs980"), // from
			chainops.MustParseAddress("b6tqv9flp7tbwp3cpsl00qmb"), // to
			synaptic.Uint64(1e6),                                  // quantity
		},
	}
)

func main() {
	w := wallet.NewDefault()

	acc1, err := w.DeriveKeyPath("default/x/test", []byte("123"))
	if err != nil {
		panic(err)
	}

	op := opcode.TODO()

	for index := 0; index < 1000; index++ {
		op.Add(txn)
	}

	pk, _ := acc1.ECPrivKey()
	// pub, _ := acc1.ECPubKey()
	s, _ := chainops.Sign(pk, txn)
	// logger.PrintJSON(s)
	b, _ := s.Marshal()
	logger.Printf("%x", b)
	logger.Printf("%d", len(b))
	// verifySignedOp(pub, err)
}

// func verifySignedOp(op *opcode.BinaryCell) (_ *opcode.BinaryCell, err error) {
// 	if op.OpCode != chainops.OpSigned || len(op.Children) < 2 {
// 		return nil, errors.New("invalid signed op")
// 	}
// 	return
// }

// func verifySignatureCell(hash []byte, op *opcode.BinaryCell) (_ *opcode.BinaryCell, err error) {
// 	if op != chainops.OpSigned || len(op.Children) < 2 {
// 		return nil, errors.New("invalid signed op")
// 	}
// }
