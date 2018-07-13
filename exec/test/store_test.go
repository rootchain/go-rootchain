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

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/gogo/protobuf/proto"
	cells "github.com/ipfn/go-ipfn-cells"
	keypair "github.com/ipfn/go-ipfn-keypair"
	ipfsdb "github.com/rootchain/go-ipfs-db"
	"github.com/rootchain/go-rootchain/exec"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *T) {
	store := initStore(nil)

	for index := 0; index < 10; index++ {
		txt := proto.EncodeVarint(uint64(index))
		key, _ := cells.SumCID(keypair.CIDPrefix, txt)
		val := uint64(1000 + index)
		err := store.Update(key, val)
		assert.Empty(t, err)
	}

	root, err := store.Commit()
	assert.Empty(t, err)
	assert.Equal(t, "z45oqTRuBFTptGFFjUFdFDTR48ETK8MkZkM184QYsAzFxSSK8mZ", root.String())

	store = initStore(root)

	for index := 0; index < 10; index++ {
		txt := proto.EncodeVarint(uint64(index))
		key, _ := cells.SumCID(keypair.CIDPrefix, txt)
		val, err := store.Get(key)
		assert.Empty(t, err)
		assert.Equal(t, uint64(1000+index), val)
		err = store.Update(key, val)
		assert.Empty(t, err)
	}

	nr, err := store.Commit()
	assert.Empty(t, err)
	assert.Equal(t, root.String(), nr.String())
	assert.Equal(t, uint64(10045), store.Total())
}

func BenchmarkStoreUpdate(b *B) {
	b.StopTimer()
	store := initStore(nil)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		txt := proto.EncodeVarint(uint64(i))
		val := uint64(i)
		key, _ := cells.SumCID(keypair.CIDPrefix, txt)
		store.Update(key, val)
	}
}

func BenchmarkStoreGet(b *B) {
	b.StopTimer()
	store := initStore(nil)
	txt := proto.EncodeVarint(uint64(1000))
	val := uint64(1000)
	key, _ := cells.SumCID(keypair.CIDPrefix, txt)
	store.Update(key, val)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		store.Get(key)
	}
}

func initStore(root *cells.CID) exec.Store {
	store, _ := exec.NewStore(root, localTrieDB)
	return store
}

var (
	localStore  = ethdb.NewMemDatabase()
	ipfsStore   = ipfsdb.Wrap(localStore)
	localTrieDB = trie.NewDatabase(ipfsStore)
)

func init() {
	// disable IPFS store in Travis CI tests
	if os.Getenv("CI") != "" || os.Getenv("TRAVIS") != "" {
		localTrieDB = trie.NewDatabase(localStore)
	}
}
