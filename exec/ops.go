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

package exec

import (
	"context"
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	keypair "github.com/ipfn/go-ipfn-keypair"
	"github.com/rootchain/go-rootchain/dev/chainops"
)

// assignOp - Assign power to public key hash or another address.
// It's only possible to assign power in genesis block.
func assignOp(state State) (res State, err error) {
	if state.Op().ChildrenSize() != 2 {
		return nil, errors.New("AssignOp: two arguments are required")
	}
	if !state.Block().IsGenesis() {
		return nil, errors.New("AssignOp: cannot assign on non-zero height")
	}
	quantity, err := uint64Op(state.Op().Child(0))
	if err != nil {
		return
	}
	assignee, err := cidOp(state.Op().Child(1))
	if err != nil {
		return
	}
	logger.Debugw("Assign Operation", "op", state.Op())
	state.Store().Set(assignee, quantity)
	return state, nil
}

// delegateOp - Delegates power.
func delegateOp(state State) (res State, err error) {
	// TODO: will change on delegation to others
	if state.Op().ChildrenSize() != 1 && state.Op().ChildrenSize() != 2 {
		return nil, errors.New("DelegateOp: two arguments are required")
	}
	pk, ok := state.Op().Context().Value(pkCtxKey).(*btcec.PublicKey)
	if !ok {
		return nil, errors.New("DelegateOp: no signature")
	}
	quantity, err := uint64Op(state.Op().Child(0))
	if err != nil {
		return
	}
	// TODO: currently only self-delegation is possible and allowed
	if state.Op().ChildrenSize() == 2 && state.Op().Child(1).OpCode() != chainops.OpNonce {
		return nil, errors.New("DelegateOp: second argument is not nonce")
	}
	addr := keypair.CID(pk)
	balance := state.Store().Get(addr)
	logger.Debugw("Delegate Power Operation",
		"source", addr.String(),
		"quantity", quantity,
		"balance", balance,
		"valid", quantity <= balance,
	)
	if quantity > balance {
		return nil, fmt.Errorf("DelegateOp: balance %d is not enough to delegate %d", balance, quantity)
	}
	return state, nil
}

type ctxKey string

const pkCtxKey ctxKey = "btcec.PublicKey"

// signatureOp - Verifies signature of parent operation.
func signatureOp(state State) (res State, err error) {
	if state.Op().Parent().OpCode() != chainops.OpSigned {
		return nil, errors.New("SignatureOp: not child of signed op")
	}
	hash := state.Op().Parent().Child(0).CID().Bytes()
	// verifies if signature is not malformed and recovers public key
	sig := state.Op().Memory()
	pk, _, err := btcec.RecoverCompact(btcec.S256(), sig, hash[6:])
	if err != nil {
		return nil, fmt.Errorf("SignatureOp: signature malformed: %v", err)
	}
	// set public key in context
	// TODO: measure impact of this and make sure it doesn't leak
	ctx := context.WithValue(state.Op().Context(), pkCtxKey, pk)
	// return public key cell
	cell := chainops.NewPubkey(pk)
	return state.WithOp(NewRoot(ctx, cell)), nil
}

// signedOp - Performs signature-verified operation.
func signedOp(state State) (res State, err error) {
	// TODO(crackcomm): implement multisig
	if state.Op().ChildrenSize() > 2 {
		return nil, errors.New("SignedOp: multisig not implemented")
	}
	if state.Op().ChildrenSize() != 2 {
		return nil, errors.New("SignedOp: two children are required")
	}
	// signature to verify
	sigOp := state.Op().ExecChild(1)
	// verify it's signature to explicitly execute it
	if sigOp.OpCode() != chainops.OpSignature {
		return nil, fmt.Errorf("SignedOp: expected a signature but got %s", sigOp.OpCode())
	}
	// execute signature op
	// it needs to verify against operation to execute
	sigState, err := signatureOp(state.WithOp(sigOp))
	if err != nil {
		return
	}
	// operation to execute
	signedContext := sigState.Op().Context()
	signedExecOp := state.Op().ExecChild(0).WithContext(signedContext)
	return Unwind(state.WithOp(signedExecOp))
}
