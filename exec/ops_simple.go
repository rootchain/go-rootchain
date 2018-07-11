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

	"github.com/gogo/protobuf/proto"

	cells "github.com/ipfn/go-ipfn-cells"
	"github.com/ipfn/go-ipfn-cells/chainops"
	"github.com/ipfn/go-ipfn-cells/synaptic"
)

func cidOp(cell cells.Cell) (*cells.CID, error) {
	if err := verifyOpCode(chainops.OpCID, cell); err != nil {
		return nil, err
	}
	return cells.ParseCID(cell.Memory())
}

func uint64Op(cell cells.Cell) (n uint64, _ error) {
	if err := verifyOpCode(synaptic.OpUint64, cell); err != nil {
		return 0, err
	}
	n, _ = proto.DecodeVarint(cell.Memory())
	return
}

func verifyOpCode(op cells.ID, cell cells.Cell) error {
	if cell.OpCode() != op {
		return fmt.Errorf("invalid cid %s", cell.OpCode())
	}
	return nil
}
