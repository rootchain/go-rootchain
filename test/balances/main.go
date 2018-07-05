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

package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
)

func main() {
	memdb := ethdb.NewMemDatabase()
	db := trie.NewDatabase(memdb)
	t, _ := trie.New(common.Hash{}, db)
	t.Update([]byte("123"), []byte("lol"))
	log.Println(t.Hash().String())
	t.Update([]byte("123"), []byte("lol123"))
	log.Println(t.Hash().String())
	t.Update([]byte("123"), []byte("lol"))
	log.Println(t.Hash().String())
	t.Update([]byte("1234"), []byte("lol"))
	t.Update([]byte("12345"), []byte("ldol"))
	for index := 0; index < 11; index++ {
		t.Update([]byte(fmt.Sprintf("12346666%d", index)), []byte(fmt.Sprintf("sdf%d", index)))
	}
	t.Update([]byte(fmt.Sprintf("12346666%d", 6)), []byte(fmt.Sprintf("sdf%d", 6)))
	t.Update([]byte(fmt.Sprintf("12346666%d", 4)), []byte(fmt.Sprintf("sdf%d", 4)))
	t.Update([]byte(fmt.Sprintf("12346666%d", 5)), []byte(fmt.Sprintf("sdf%d", 5)))
	t.Update([]byte("12347"), []byte("lold"))
	log.Println(t.Hash().String())
	t.Update([]byte("123"), []byte("lo4l"))
	root, err := t.Commit(func(leaf []byte, parent common.Hash) error {
		log.Printf("leaf: 0x%x", leaf)
		return nil
	})
	if err != nil {
		panic(err)
	}
	log.Printf("root: %s", root.String())
	t.Update([]byte("123"), []byte("lol"))
	log.Println(t.Hash().String())
	log.Println(root.String())
	t2, err := trie.New(root, db)
	if err != nil {
		panic(err)
	}
	t2.Update([]byte("123"), []byte("lol"))
	log.Println(t2.Hash().String())
	log.Printf("12345: %s", t2.Get([]byte("12345")))
}
