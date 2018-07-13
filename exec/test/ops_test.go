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

package exectest

import (
	"os"
	. "testing"

	cells "github.com/ipfn/go-ipfn-cells"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	wallet "github.com/ipfn/go-ipfn-wallet"
	"github.com/rootchain/go-rootchain/dev/chainops"
	"github.com/rootchain/go-rootchain/dev/genesis"
	"github.com/rootchain/go-rootchain/exec"
	"github.com/stretchr/testify/assert"
)

func init() {
	logger.Disable()
	wallet.SetDefaultPath(os.TempDir())
}

func TestAssignOp(t *T) {
	w := wallet.NewDefault()
	state := initState(w)
	assert.Equal(t, uint64(3e6), state.Store().Total())
	{
		key, _ := w.UnlockedDerive(wallet.MustParseKeyPath("default/x/first"))
		c, _ := key.CID()
		val, _ := state.Store().Get(c)
		assert.Equal(t, uint64(1e6), val)
	}
	{
		key, _ := w.UnlockedDerive(wallet.MustParseKeyPath("default/x/assign-op-test"))
		c, _ := key.CID()
		val, _ := state.Store().Get(c)
		assert.Equal(t, uint64(0), val)
		state, err := exec.NextState(state, chainops.NewRootOp(
			chainops.NewAssignPowerOp(0, 1000, c),
		))
		assert.Empty(t, err)
		_, err = exec.Unwind(state)
		assert.Equal(t, "AssignOp: cannot assign on non-zero height", err.Error())
	}
}

func TestDelegateOp(t *T) {
	w := wallet.NewDefault()
	state := initState(w)
	assert.Equal(t, uint64(3e6), state.Store().Total())

	// NOTE: this is hazardous it does not update exec hash but doesnt matter here
	signedOp := state.Op().Child(state.Op().ChildrenSize() - 1).(*cells.BinaryCell)
	signedOp.SetChildren([]cells.Cell{
		chainops.NewDelegatePowerOp(0, 2),
		signedOp.Child(1),
	})

	state, err := exec.Unwind(state)
	assert.Equal(t, "DelegateOp: balance 0 is not enough to delegate 2", err.Error())
}

func BenchmarkDelegateOp(b *B) {
	b.StopTimer()
	w := wallet.NewDefault()
	state := initState(w)
	key, _ := w.UnlockedDerive(wallet.MustParseKeyPath("default/x/third"))
	privKey, err := key.ECPrivKey()
	if err != nil {
		b.Fatal(err)
	}
	delegateOp := chainops.NewDelegatePowerOp(0, 1e6)
	signedOp, err := chainops.NewSignedOp(delegateOp, privKey)
	if err != nil {
		b.Fatal(err)
	}
	state, _ = exec.NextState(state, chainops.NewRootOp(signedOp))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := exec.Unwind(state)
		if err != nil {
			b.Fatal(err)
		}
	}
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
	delegateOp := chainops.NewDelegatePowerOp(0, 1)
	return chainops.NewSignedOp(delegateOp, privKey)
}

func initState(w *wallet.Wallet) exec.State {
	defer logger.Sync()
	if ok, err := w.KeyExists("default"); !ok {
		_, err := w.CreateSeed("default", []byte("123"))
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}
	_, err := w.Unlock("default", []byte("123"))
	if err != nil {
		panic(err)
	}
	config := &genesis.Config{
		Wallet:   w,
		Database: localTrieDB,
	}
	config.Assign(genesis.MustParsePowerString("default/x/first:1e6:1e6"))
	config.Assign(genesis.MustParsePowerString("default/x/second:1e6:1e6"))
	config.Assign(genesis.MustParsePowerString("default/x/third:1e6:0"))
	head, err := genesis.Init(config)
	if err != nil {
		panic(err)
	}
	store := initStore(head.State())
	return exec.NewState(store, head)
}
