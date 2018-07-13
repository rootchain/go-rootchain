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
	"strconv"
	"strings"

	cells "github.com/ipfn/go-ipfn-cells"
	wallet "github.com/ipfn/go-ipfn-wallet"
)

// Distribution - Genesis power distribution.
type Distribution struct {
	// Address - Usually public key hash.
	Address *cells.CID

	// WalletKeyPath - Wallet key path.
	WalletKeyPath *wallet.KeyPath

	// AssignQuantity - Quantity of power to assign.
	AssignQuantity uint64

	// DelegateQuantity - Quantity of power to delegate.
	// Only possible if wallet key path is set.
	DelegateQuantity uint64
}

// ParsePowerString - Parses power distribution string.
//
// Following formats are allowed:
//   - `<wallet-key-path>:<assign-power>:<delegate-power>`
//   - `<pubkeyhash-addr>:<assign-power>`
//
// If delegate value is `-1` it's equal to assign power.
func ParsePowerString(keyPath string) (res *Distribution, err error) {
	split := strings.Split(keyPath, ":")
	if len(split) < 2 || len(split) > 3 {
		return nil, fmt.Errorf("invalid key:power:delegated format: %q", keyPath)
	}
	res = new(Distribution)
	if strings.HasPrefix(split[0], "zFNSc") && !strings.Contains(split[0], "/") && len(split) == 2 {
		res.Address, err = cells.DecodeCID(split[0])
	} else {
		res.WalletKeyPath, err = wallet.ParseKeyPath(split[0])
	}
	if err != nil {
		return nil, err
	}
	assign, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
		return nil, err
	}
	res.AssignQuantity = uint64(assign)
	if len(split) == 3 {
		delegate, err := strconv.ParseFloat(split[2], 64)
		if err != nil {
			return nil, err
		}
		if delegate == -1 {
			delegate = assign
		}
		res.DelegateQuantity = uint64(delegate)
	}
	return
}

// MustParsePowerString - Parses power distribution string. Panics on error.
func MustParsePowerString(keyPath string) (res *Distribution) {
	res, err := ParsePowerString(keyPath)
	if err != nil {
		panic(err)
	}
	return
}
