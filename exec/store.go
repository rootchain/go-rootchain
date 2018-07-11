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
	cells "github.com/ipfn/go-ipfn-cells"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
)

// Store - Execution store.
type Store interface {
	// Set - Sets value under key.
	Set(key *cells.CID, value uint64)
}

// NewStore - Creates new mutable execution store.
func NewStore() Store {
	return &execStore{}
}

type execStore struct {
	root *Store
}

func (s *execStore) Set(key *cells.CID, value uint64) {
	logger.Infow("Store Set", "key", key, "value", value)
}
