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
	. "testing"

	"github.com/ipfn/go-ipfn-cells"

	wallet "github.com/ipfn/go-ipfn-wallet"
	"github.com/stretchr/testify/assert"
)

func TestParsePowerString(t *T) {
	expected := &Distribution{
		AssignQuantity:   1e6,
		DelegateQuantity: 1e6,
		WalletKeyPath:    wallet.MustParseKeyPath("default/x/first"),
	}
	parsed, err := ParsePowerString("default/x/first:1e6:1e6")
	assert.Empty(t, err)
	assert.Equal(t, expected, parsed)
}

func TestParsePowerString_Address(t *T) {
	addr, _ := cells.DecodeCID("zFNScYMHAiBJPJcT9ri2cNwKTJfitbRcE2PSfMKQj272C6A5Fs25")
	parsed, err := ParsePowerString("zFNScYMHAiBJPJcT9ri2cNwKTJfitbRcE2PSfMKQj272C6A5Fs25:1e6")
	assert.Empty(t, err)
	assert.Equal(t, parsed, &Distribution{
		AssignQuantity: 1e6,
		Address:        addr,
	})
	parsed, err = ParsePowerString("zFNScYMHA:1e6")
	assert.Empty(t, err)
	assert.Equal(t, parsed, &Distribution{
		AssignQuantity: 1e6,
		WalletKeyPath:  wallet.MustParseKeyPath("zFNScYMHA"),
	})
}

func TestParsePowerString_Chain(t *T) {
	addr, _ := cells.DecodeCID("zFSec2XV8EGiAKaBS8xL8meisn9KyKVwY7UprRH2x51SBgPXKGZM")
	parsed, err := ParsePowerString("zFSec2XV8EGiAKaBS8xL8meisn9KyKVwY7UprRH2x51SBgPXKGZM:1e6")
	assert.Empty(t, err)
	assert.Equal(t, parsed, &Distribution{
		AssignQuantity: 1e6,
		Address:        addr,
	})
	parsed, err = ParsePowerString("zFSec2XV8EG:1e6")
	assert.Empty(t, err)
	assert.Equal(t, parsed, &Distribution{
		AssignQuantity: 1e6,
		WalletKeyPath:  wallet.MustParseKeyPath("zFSec2XV8EG"),
	})
}
