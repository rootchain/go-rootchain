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

import "github.com/rootchain/go-rootchain/chain"

// State - Execution stack context.
type State interface {
	// Op - Execution cell.
	Op() Cell

	// Head - Chain head state.
	Head() *chain.State

	// Store - Execution state.
	Store() Store

	// WithOp - Switches op cell.
	WithOp(Cell) State
}

// NewState - Creates new execution state.
func NewState(head *chain.State, store Store, op Cell) State {
	return &execState{op: op, head: head, store: store}
}

type execState struct {
	op    Cell
	head  *chain.State
	store Store
}

func (s *execState) Op() Cell {
	return s.op
}

func (s *execState) Head() *chain.State {
	return s.head
}

func (s *execState) Store() Store {
	return s.store
}

func (s *execState) WithOp(op Cell) State {
	return NewState(s.head, s.store, op)
}
