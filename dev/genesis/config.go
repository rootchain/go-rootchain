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
	wallet "github.com/ipfn/go-ipfn-wallet"
)

// Config - Genesis config.
type Config struct {
	// Wallet - Wallet to use for genesis.
	// Default wallet is used if empty.
	Wallet *wallet.Wallet

	// Power - Initial power distribution.
	Power []*Distribution
}

// Assign - Assign power distribution.
func (config *Config) Assign(power *Distribution) {
	config.Power = append(config.Power, power)
}

// WalletKeys - Unique wallet keys requiring unlocking.
func (config *Config) WalletKeys() (keys []string) {
	for _, alloc := range config.Power {
		if alloc.WalletKeyPath == nil {
			continue
		}
		keys = appendIfMissing(keys, alloc.WalletKeyPath.SeedName)
	}
	return keys
}

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
