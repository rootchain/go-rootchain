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
	. "testing"

	cells "github.com/ipfn/go-ipfn-cells"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	wallet "github.com/ipfn/go-ipfn-wallet"
	"github.com/rootchain/go-rootchain/dev/chainops"
	"github.com/rootchain/go-rootchain/dev/genesis"
	"github.com/stretchr/testify/assert"
)

func TestAssignOp(t *T) {
	w := wallet.NewDefault()
	state := initState(w)

	key, _ := w.UnlockedDerive(wallet.MustParseKeyPath("default/x/assign-op-test"))
	c, _ := key.CID()

	state, _ = NextState(state, chainops.NewRoot(
		chainops.NewAssignPower(0, 1000, c),
	))

	_, err := Unwind(state)
	assert.Equal(t, "AssignOp: cannot assign on non-zero height", err.Error())
}

func TestDelegateOp(t *T) {
	w := wallet.NewDefault()
	state := initState(w)

	_, err := Unwind(state)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(3e6), state.Store().Total())

	// NOTE: this is hazardous it does not update exec hash but doesnt matter here
	signedOp := state.Op().Child(state.Op().ChildrenSize() - 1).(*cells.BinaryCell)
	signedOp.SetChildren([]cells.Cell{
		chainops.NewDelegatePower(0, 2),
		signedOp.Child(1),
	})

	state, err = Unwind(state)
	assert.Equal(t, "DelegateOp: balance 0 is not enough to delegate 2", err.Error())
}

func newSignedOp(w *wallet.Wallet) (_ *cells.BinaryCell, err error) {
	key, err := w.UnlockedDerive(wallet.MustParseKeyPath("default/x/signed-op-test"))
	if err != nil {
		return
	}
	privKey, err := key.ECPrivKey()
	if err != nil {
		return
	}
	delegateOp := chainops.NewDelegatePower(0, 1)
	return chainops.NewSignOperation(delegateOp, privKey)
}

func initState(w *wallet.Wallet) State {
	defer logger.Sync()
	_, err := w.Unlock("default", []byte("123"))
	if err != nil {
		panic(err)
	}

	config := &genesis.Config{Wallet: w}

	assignPaths := []string{
		"default/x/first:1e6:1e6",
		"default/x/second:1e6:1e6",
		"default/x/third:1e6:0",
	}

	for _, path := range assignPaths {
		power, err := genesis.ParsePowerString(path)
		if err != nil {
			panic(err)
		}
		config.Assign(power)
	}

	head, err := genesis.Init(config)
	if err != nil {
		panic(err)
	}

	return NewState(NewStore(), head)
}
