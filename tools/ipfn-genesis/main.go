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
	"github.com/ipfn/ipfn/go/chain"
	"github.com/ipfn/ipfn/go/opcode"
	"github.com/ipfn/ipfn/go/opcode/chainops"
	"github.com/ipfn/ipfn/go/opcode/synaptic"
)

func main() {
	genesisOp := &opcode.BinaryCell{OpCode: chainops.OpGenesis}

	allocOp := &opcode.BinaryCell{
		OpCode: chainops.OpAllocation,
		Children: []*opcode.BinaryCell{
			chainops.MustParseAddress("b7dlu9ahtazhar30psm4sqlc"),
			synaptic.Uint64(1e14), // 100m of base units
		},
	}

	claimOp := &opcode.BinaryCell{
		OpCode: chainops.OpClaim,
		Children: []*opcode.BinaryCell{
			chainops.MustParseAddress("b7dlu9ahtazhar30psm4sqlc"),
			// (here: op.pubkey)
		},
	}

	state, _ := chain.NewState(0, nil, opcode.Ops(genesisOp, allocOp, claimOp))
	logger.PrintJSON(state)

	b, _ := chainops.MustParseAddress("b7dlu9ahtazhar30psm4sqlc").Marshal()
	logger.Printf("%d", len(b))
	logger.Printf("%d", len(state.Header.StateHash.Bytes()))

	// txn := &opcode.BinaryCell{
	// 	OpCode: chainops.OpTransfer,
	// 	Children: []*opcode.BinaryCell{
	// 		chainops.NewIDFromString("test"),
	// 		synaptic.NewUint64(1e6), // quantity
	// 	},
	// }
	// state, _ = chain.NewState(1, state.Header.StateHash, opcode.Ops(txn))
	// logger.PrintJSON(state)
}
