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

package genesis

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	cells "github.com/ipfn/go-ipfn-cells"
	wallet "github.com/ipfn/go-ipfn-wallet"
	"github.com/rootchain/go-rootchain/chain"
	"github.com/rootchain/go-rootchain/dev/chainops"
)

// Init - Initializes a new chain.
func Init(config *Config) (state *chain.State, err error) {
	// set default wallet if empty
	if config.Wallet == nil {
		config.Wallet = wallet.NewDefault()
	}
	var (
		privKeys    []*btcec.PrivateKey
		assignOps   []cells.Cell
		delegateOps []cells.Cell
	)
	// derive private keys for all key paths
	var nonce cells.ID
	for _, dest := range config.Power {
		key, err := config.Wallet.UnlockedDerive(dest.WalletKeyPath)
		if err != nil {
			return nil, fmt.Errorf("wallet %s: %v", dest.WalletKeyPath.SeedName, err)
		}
		privKey, err := key.ECPrivKey()
		if err != nil {
			return nil, err
		}
		addr, err := key.CID()
		if err != nil {
			return nil, err
		}
		if dest.DelegateQuantity > 0 {
			delegateOp := chainops.DelegatePower(nonce, dest.DelegateQuantity)
			signedOp, err := chainops.Sign(delegateOp, privKey)
			if err != nil {
				return nil, err
			}
			privKeys = append(privKeys, privKey)
			delegateOps = append(delegateOps, signedOp)
			nonce++
		}
		assignOps = append(assignOps, chainops.AssignPower(0, dest.AssignQuantity, addr))
	}
	// chain exec root op
	root := chainops.Root()
	root.AddChildren(assignOps...)
	root.AddChildren(delegateOps...)
	// initialize state
	state, err = chain.NewState(0, nil, root)
	if err != nil {
		return nil, err
	}
	// sign state with private keys
	for _, key := range privKeys {
		if _, err := state.Sign(key); err != nil {
			return nil, err
		}
	}
	return
}
