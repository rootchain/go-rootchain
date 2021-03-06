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
	"fmt"

	"github.com/rootchain/go-rootchain/dev/chainops"
)

// Unwind - Unwinds execution of cells.
func Unwind(state State) (State, error) {
	c := state.Op()
	switch c.OpCode() {
	case chainops.OpRoot:
		count := c.ChildrenSize()
		for index := 0; index < count; index++ {
			cell := c.ExecChild(index)
			exec := state.WithOp(cell)
			_, err := Unwind(exec)
			if err != nil {
				return nil, err
			}
		}
		return state, nil
	case chainops.OpSigned:
		return signedOp(state)
	case chainops.OpAssignPower:
		return assignOp(state)
	case chainops.OpDelegatePower:
		return delegateOp(state)
	}
	return nil, fmt.Errorf("unimplemented unwind for %s", c.OpCode())
}
