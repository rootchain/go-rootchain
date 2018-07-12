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
	// Total - Total power.
	Total() uint64
	// Get - Gets value under key.
	Get(key *cells.CID) uint64
	// Set - Sets value under key.
	Set(key *cells.CID, value uint64)
	// Clone - Clones store.
	Clone() Store
}

// NewStore - Creates new mutable execution store.
func NewStore() Store {
	return &execStore{maps: make(map[string]uint64)}
}

type execStore struct {
	root  *execStore
	maps  map[string]uint64
	total uint64
}

func (s *execStore) Get(key *cells.CID) uint64 {
	return s.get(key.String())
}

func (s *execStore) Set(key *cells.CID, value uint64) {
	s.set(key.String(), value)
}

func (s *execStore) Total() uint64 {
	return s.total
}

func (s *execStore) Clone() Store {
	return &execStore{root: s, total: s.total, maps: make(map[string]uint64)}
}

func (s *execStore) set(cid string, value uint64) {
	prev := s.get(cid)
	s.total += value - prev
	s.maps[cid] = value
	logger.Infow("Store Set", "key", cid, "value", value, "total", s.total, "prev", prev)
}

func (s *execStore) get(key string) uint64 {
	// todo: make sure zero-ing doesnt mess-up `ok` value
	if value, ok := s.maps[key]; ok {
		return value
	}
	if s.root == nil {
		return 0
	}
	return s.root.get(key)
}
