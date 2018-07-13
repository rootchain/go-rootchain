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
	. "testing"

	"github.com/ipfn/go-base32check"

	"github.com/ipfn/go-ipfn-cells"
	"github.com/rootchain/go-rootchain/dev/chainops"
	"github.com/rootchain/go-rootchain/dev/synaptic"
	"github.com/stretchr/testify/assert"
)

var (
	genesisEnc = "8s00zsc00bljv0rvcy03kb9kuw6490dy5qpx82rh3brxb22qx08r6098bfx653av2vuf57zlmv0p7070ss7s0"
)

func TestBinaryCell(t *T) {
	c, _ := cells.DecodeCID("zFNScYMH8j9JuHxLR5KLsNP528LuLBi7ToCrh9tmdLb85pWno5Bg")
	allocOp := cells.Op(chainops.OpAssignPower,
		chainops.NewCIDOp(c),
		synaptic.Uint64(1e6))

	var head string
	block, err := NewBlock(0, nil, chainops.NewRootOp(allocOp))
	assert.Empty(t, err)
	assert.Equal(t, uint64(0), block.Height())
	// assert.Equal(t, "zFSec2XVAw1qbBFm7rFFV81U8UwGqBLuV7SGze8vPyQbziy7zbku", block.Head().String())
	assert.Equal(t, "zFuncm69C7pUHeGn1cTt5i8YHKsRxyuLrU4PJwMFXF75jvseJdAK", block.header.Exec.String())
	head = block.Head().String()

	body, err := block.Root().Marshal()
	assert.Equal(t, genesisEnc, base32check.EncodeToString(body))

	block, err = block.Next(chainops.NewRootOp(allocOp))
	assert.Empty(t, err)
	assert.Equal(t, uint64(1), block.Height())
	assert.Equal(t, head, block.Prev().String())
	// assert.Equal(t, "zFSec2XVBdAsns1xHLkvMhcJnvtK4tchNG8DBhiXcZd4kUqQbGVa", block.Head().String())
	// assert.Equal(t, "zFSec2XVGN5saHLfhnwm3s5TRXTPNGGcbggyfXSYLP97nxn6HStJ", block.Header.Exec.String())
	head = block.Head().String()

	block, err = block.Next(chainops.NewRootOp(allocOp))
	assert.Empty(t, err)
	assert.Equal(t, uint64(2), block.Height())
	assert.Equal(t, head, block.Prev().String())
	// assert.Equal(t, "zFSec2XV3e6uoSd8sTcBM717xMweVCBQFBoMQc4Qy3JGtYMtu34E", block.Head().String())
	// assert.Equal(t, "zFSec2XVGN5saHLfhnwm3s5TRXTPNGGcbggyfXSYLP97nxn6HStJ", block.Header.Exec.String())
}
