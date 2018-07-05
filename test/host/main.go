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
	"context"
	"fmt"

	"github.com/ipfn/go-ipfn-cmd-util/logger"
	"github.com/ipfn/ipfn/go/host"
	"github.com/ipfn/ipfn/go/wallet"
	libp2p "github.com/libp2p/go-libp2p"
)

func main() {
	// The context governs the lifetime of the libp2p node
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := wallet.NewDefault()

	acc1, err := w.DeriveKeyPath("default/x/test1", []byte("123"))
	if err != nil {
		panic(err)
	}

	h, err := host.New(ctx,
		host.KeyPair(acc1),
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/9200"),
	)

	fmt.Printf("Hello World, my hosts ID is %s\n", h.ID().Pretty())

	acc2, err := w.DeriveKeyPath("default/x/test2", []byte("123"))
	if err != nil {
		panic(err)
	}

	h2, err := host.New(ctx,
		host.KeyPair(acc2),
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/9000"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello World, my second hosts ID is %s\n", h2.ID().Pretty())
	pk, err := h2.RecoverPublicKey()
	if err != nil {
		panic(err)
	}
	if pk == nil {
		logger.Print("not recovered!")
	} else {
		fmt.Printf("Hello World, my second hosts public key is %x (recovered from id)\n", pk.SerializeCompressed())
	}

	select {}
}
