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
)

// Store - Execution store.
type Store interface {
	// Get - Gets value under key.
	Get(key *cells.CID) uint64
	// Set - Sets value under key.
	Set(key *cells.CID, value uint64)
	// Total - Gets total amount stored.
	Total() uint64
}

// NewStore - Creates new mutable execution store.
func NewStore() Store {
	return &execStore{maps: make(map[string]uint64)}
}

type execStore struct {
	root  *Store
	maps  map[string]uint64
	total uint64
}

func (s *execStore) Get(key *cells.CID) uint64 {
	return s.maps[key.String()]
}

func (s *execStore) Set(key *cells.CID, value uint64) {
	cid := key.String()
	prev := s.maps[cid]
	s.total += value - prev
	s.maps[cid] = value
	// logger.Infow("Store Set", "key", key, "value", value, "total", s.total, "prev", prev)
}

func (s *execStore) Total() uint64 {
	return s.total
}
