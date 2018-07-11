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
	cells "github.com/ipfn/go-ipfn-cells"
	"github.com/ipfn/go-ipfn-cells/chainops"
	"github.com/ipfn/go-ipfn-cells/synaptic"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	keypair "github.com/ipfn/go-ipfn-keypair"
)

// assignOp - Assign power.
func assignOp(state State) (res State, err error) {
	if !state.Head().IsGenesis() {
		return nil, errors.New("cannot assign on non-zero height")
	}
	quantity, err := uint64Op(state.Op().Child(0))
	if err != nil {
		return
	}
	assignee, err := cidOp(state.Op().Child(1))
	if err != nil {
		return
	}
	state.Store().Set(assignee, quantity)
	logger.Infow("Assign Operation", "op", state.Op())
	return state, nil
}

// delegateOp - Delegates power.
func delegateOp(state State) (res State, err error) {
	if state.Op().ChildrenSize() == 0 {
		return nil, errors.New("DelegateOp: no children")
	}
	pk, ok := state.Op().Context().Value(pkCtxKey).(*btcec.PublicKey)
	if !ok {
		return nil, errors.New("DelegateOp: no signature")
	}
	// TODO: currently only self-delegation is possible and allowed
	// if state.Op().ChildrenSize() == 1 || state.Op().Child(1).OpCode() == chainops.OpNonce {
	// }
	quantity, err := uint64Op(state.Op().Child(0))
	if err != nil {
		return
	}
	logger.Infow("Delegate Power Operation",
		"cid", keypair.CID(pk).String(),
		"quantity", quantity,
	)

	op := NewCell(state.Op().Context(), state.Op(), cells.Op(chainops.OpTransfer,
		chainops.Pubkey(pk),
		synaptic.String("key"),
		synaptic.String("value"),
	))
	return state.WithOp(op), nil
}

type ctxKey string

const pkCtxKey ctxKey = "btcec.PublicKey"

// signatureOp - Verifies signature of parent operation.
func signatureOp(state State) (res State, err error) {
	cid := state.Op().Parent().Child(0).CID().Bytes()
	sigOp := state.Op().Memory()
	// verifies signature and recovers public key
	pk, _, err := btcec.RecoverCompact(btcec.S256(), sigOp, cid)
	if err != nil {
		return nil, fmt.Errorf("SignatureOp: invalid signature: %v", err)
	}
	// set public key in context
	// TODO: measure impact of this and make sure it doesn't leak
	ctx := context.WithValue(state.Op().Context(), pkCtxKey, pk)
	return state.WithOp(NewCell(ctx, state.Op(), chainops.Pubkey(pk))), nil
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
	// execute signature op
	sigStack, err := signatureOp(state.WithOp(sigOp))
	if err != nil {
		return
	}
	// operation to execute
	execOp := state.Op().ExecChild(0)
	ctx := sigStack.Op().Context()
	return Unwind(state.WithOp(execOp.WithContext(ctx)))
}
