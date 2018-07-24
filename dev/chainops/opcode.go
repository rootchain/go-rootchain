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

import "github.com/ipfn/go-ipfn-cells"

const (
	// OpRoot - Offset of chain operation code (60).
	OpRoot cells.ID = 0x3c

	// OpMultihash - Multihash native operation code.
	OpMultihash = OpRoot + 1

	// OpID - ID operation.
	OpID = OpRoot + 2

	// OpCID - Contentn ID native operation code.
	OpCID = OpRoot + 3

	// OpHeader - Header operation.
	OpHeader = OpRoot + 4

	// OpNonce - Nonce op (noop).
	OpNonce = OpRoot + 5

	// OpClaim - Address claim operation.
	OpClaim = OpRoot + 6

	// OpAssignPower - Allocation of power operation.
	OpAssignPower = OpRoot + 7

	// OpDelegatePower - Investment of power operation.
	OpDelegatePower = OpRoot + 8

	// OpPubkey - Public key operation.
	OpPubkey = OpRoot + 9

	// OpPubkeyAddr - Public key hash operation.
	OpPubkeyAddr = OpRoot + 10

	// OpSignature - Signature operation.
	OpSignature = OpRoot + 11

	// OpSigned - Signed operation.
	OpSigned = OpRoot + 12

	// OpAddress - Address native operation code.
	OpAddress = OpRoot + 13

	// OpTransfer - Transfer of an asset.
	OpTransfer = OpRoot + 14

	// OpNoop - Signalizes operation finish.
	OpNoop = OpRoot + 15
)

func init() {
	cells.Register(OpID, "id")
	cells.Register(OpRoot, "root")
	cells.Register(OpHeader, "header")
	cells.Register(OpAssignPower, "assign_power")
	cells.Register(OpDelegatePower, "delegate_power")
	cells.Register(OpSignature, "signature")
	cells.Register(OpPubkey, "pubkey")
	cells.Register(OpPubkeyAddr, "pubkey_addr")
	cells.Register(OpSigned, "signed")
	cells.Register(OpAddress, "address")
	cells.Register(OpMultihash, "multihash")
	cells.Register(OpCID, "cid")
	cells.Register(OpClaim, "claim")
	cells.Register(OpTransfer, "transfer")
	cells.Register(OpNonce, "nonce")
	cells.Register(OpNoop, "noop")
}
