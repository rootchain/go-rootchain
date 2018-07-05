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

package chain

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ipfn/go-ipfn-cmd-util/logger"
	"github.com/rootchain/go-rootchain/cells"
	"github.com/rootchain/go-rootchain/cells/chainops"

	"github.com/btcsuite/btcd/btcec"
)

// State - Chain state structure.
type State struct {
	header     *Header
	opsRoot    cells.MutableCell
	sigRoot    cells.MutableCell
	signatures [][]byte
}

// NewState - Creates new state structure.
func NewState(index uint64, prevHash *cells.CID, execOps []cells.Cell) (_ *State, err error) {
	if prevHash == nil && index > 0 {
		return nil, fmt.Errorf("prev hash cannot be empty with index %d", index)
	}
	opsRoot := cells.Root(execOps)
	execCID := opsRoot.CID()
	// execCID, err := opsRoot.CID()
	// if err != nil {
	// 	return
	// }
	header, err := NewHeader(index, prevHash, execCID)
	if err != nil {
		return
	}
	sigRoot := chainops.Signed(opsRoot)
	return &State{
		header:  header,
		opsRoot: opsRoot,
		sigRoot: sigRoot,
	}, nil
}

// Head - Returns head CID.
func (state *State) Head() *cells.CID {
	if state.header.Head == nil {
		state.header.Head = state.opsRoot.CID()
	}
	return state.header.Head
}

// Signed - Returns signed head CID.
func (state *State) Signed() *cells.CID {
	return state.header.Signed
}

// Prev - Returns previous CID.
func (state *State) Prev() *cells.CID {
	return state.header.Prev
}

// Height - Returns state index.
func (state *State) Height() uint64 {
	return state.header.Height
}

// Root - Returns root operation.
func (state *State) Root() cells.Cell {
	return state.opsRoot
}

// Signatures - Returns state signatures.
func (state *State) Signatures() [][]byte {
	return state.signatures
}

// Exec - Adds operation to execute.
func (state *State) Exec(op cells.Cell) {
	state.opsRoot.AddChild(op)
	state.reset()
}

func (state *State) reset() {
	state.header.Head = nil
	state.header.Signed = nil
}

// Next - Returns next state including given ops.
func (state *State) Next(exec []cells.Cell) (*State, error) {
	if len(exec) == 0 {
		return nil, errors.New("cannot produce state with zero operations")
	}
	return NewState(state.Height()+1, state.Head(), exec)
}

// Sign - Signs state with given private key.
// Computes new signed header hash.
func (state *State) Sign(key *btcec.PrivateKey) (_ cells.Cell, err error) {
	// BUG(crackcomm): proper fucking signature xD
	sigOp, err := chainops.SignBytes(state.Head().Bytes(), key)
	if err != nil {
		return
	}
	state.sigRoot.AddChild(sigOp)
	logger.Printf("sigRoot: %#v", cells.NewPrinter(sigOp).String())
	state.signatures = append(state.signatures, sigOp.Memory())
	body, err := cells.Marshal(state.sigRoot)
	if err != nil {
		return
	}
	state.header.Signed, err = cells.SumCID(SignedPrefix, body)
	return
}

type stateJSON struct {
	Header     *Header              `json:"header,omitempty"`
	ExecOps    []*cells.CellPrinter `json:"exec_ops,omitempty"`
	Signatures [][]byte             `json:"signatures,omitempty"`
}

// MarshalJSON - Marshals state to JSON.
func (state *State) MarshalJSON() ([]byte, error) {
	return json.Marshal(stateJSON{
		Header:     state.header,
		ExecOps:    cells.NewChildrenPrinter(state.opsRoot),
		Signatures: state.signatures,
	})
}