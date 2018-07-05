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

package sign

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"golang.org/x/crypto/sha3"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	keypair "github.com/ipfn/go-ipfn-keypair"
	wallet "github.com/ipfn/go-ipfn-wallet"
)

var (
	keyPath  string
	filePath string
)

func init() {
	SignCmd.PersistentFlags().StringVarP(&keyPath, "key-path", "k", "default", "wallet key path (<wallet>/<x|m>/<path>)")
	SignCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "path of file to sign")
}

// SignCmd - Key sign command.
var SignCmd = &cobra.Command{
	Use:   "sign [content]",
	Short: "Sign with ECDSA",
	Example: `  $ rcx wallet sign -w example -xd mnemonic '{"value": "0xd"}'
  $ rcx wallet sign -d m/44'/138'/0'/0/0 '{"value": "0xd"}'
  $ rcx wallet sign -w example -d m/44'/138'/0'/0/0 '{"value": "0xd"}'`,
	Long: `Signs content with key derived from wallet.

Reads console stdin on empty arguments and -f file path flag.

Default wallet name used is "default".`,
	Annotations: map[string]string{"category": "wallet"},
	Args:        checkSignArgs,
	Run:         cmdutil.WrapCommand(HandleSignCmd),
}

// HandleSignCmd - Handles key sign command.
func HandleSignCmd(cmd *cobra.Command, args []string) (err error) {
	acc, err := deriveWallet()
	if err != nil {
		return
	}
	logger.Print("Reading content from stdin…")
	body, err := readContent(args)
	if err != nil {
		return
	}
	priv, err := acc.ECPrivKey()
	if err != nil {
		return
	}
	hash := sha3.Sum512(body)
	signature, err := priv.Sign(hash[:])
	if err != nil {
		return
	}
	pubKey, err := acc.ECPubKey()
	if err != nil {
		return
	}
	if !signature.Verify(hash[:], pubKey) {
		return errors.New("Cannot verify signature")
	}
	c, err := acc.CID()
	if err != nil {
		return
	}
	sigBytes := signature.Serialize()
	logger.Print()
	logger.Printf("Signature CID:  %s", c)
	logger.Printf("Signature hex:  %x", sigBytes)
	logger.Printf("Signature hash: %x", sha3.Sum512(sigBytes))
	return
}

func readContent(args []string) (body []byte, err error) {
	if filePath != "" {
		return ioutil.ReadFile(filePath)
	}
	if len(args) > 0 {
		return []byte(strings.Join(args, " ")), nil
	}
	return ioutil.ReadAll(os.Stdin)
}

func checkSignArgs(cmd *cobra.Command, args []string) (err error) {
	if keyPath == "" {
		return errors.New("derivation path cannot be empty")
	}
	return nil
}

func deriveWallet() (_ *keypair.KeyPair, err error) {
	path, err := wallet.ParseKeyPath(keyPath)
	if err != nil {
		return
	}
	return wallet.PromptDeriveKey(path)
}
