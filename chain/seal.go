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

// Package chain implements IPFN rootchain.
package chain

import "math"

// SealPower - Returns amount of power required for sealing a block.
func SealPower(delegated uint64) uint64 {
	return uint64(math.Ceil(float64(delegated) * 0.51))
}

// MinDelegatedPower - Returns amount of minimum power delegated.
func MinDelegatedPower(total uint64) uint64 {
	return uint64(math.Ceil(float64(total) * 0.42))
}
