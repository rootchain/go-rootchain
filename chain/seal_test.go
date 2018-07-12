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

	"github.com/stretchr/testify/assert"
)

func TestSealPower(t *T) {
	var total uint64 = 3
	minimum := MinDelegatedPower(total)
	assert.Equal(t, uint64(2), minimum)
	assert.Equal(t, uint64(2), SealPower(minimum))
	total = 4
	minimum = MinDelegatedPower(total)
	assert.Equal(t, uint64(2), minimum)
	assert.Equal(t, uint64(2), SealPower(minimum))
	assert.Equal(t, uint64(2), SealPower(3))
	assert.Equal(t, uint64(3), SealPower(4))
}
