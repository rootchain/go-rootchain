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

	cells "github.com/ipfn/go-ipfn-cells"
	"github.com/rootchain/go-rootchain/core/chain"
)

// State - Execution stack context.
type State interface {
	// Op - Execution cell.
	Op() Cell

	// Block - Chain block.
	Block() *chain.Block

	// Store - Execution state.
	Store() Store

	// WithOp - Switches op cell.
	WithOp(Cell) State
}

// NewState - Creates new execution state.
func NewState(store Store, block *chain.Block) State {
	op := NewRoot(context.TODO(), block.Root())
	return &execState{op: op, block: block, store: store}
}

// NextState - Creates state for next block from prev state and exec root.
func NextState(state State, execRoot cells.MutableCell) (State, error) {
	block, err := state.Block().Next(execRoot)
	if err != nil {
		return nil, err
	}
	store, err := state.Store().Clone()
	if err != nil {
		return nil, err
	}
	return NewState(store, block), nil
}

type execState struct {
	op    Cell
	store Store
	block *chain.Block
}

func (s *execState) Op() Cell {
	return s.op
}

func (s *execState) Block() *chain.Block {
	return s.block
}

func (s *execState) Store() Store {
	return s.store
}

func (s *execState) WithOp(op Cell) State {
	return &execState{op: op, block: s.block, store: s.store}
}
