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

package address

import (
	"errors"
	"fmt"
	"hash/crc32"
	"math"

	cid "gx/ipfs/QmapdYm1b22Frv3k17fqrBYTFRxwiaVJkB299Mfn33edeB/go-cid"

	"github.com/gogo/protobuf/proto"
	base32check "github.com/ipfn/go-base32check"
	"github.com/ipfn/ipfn/go/opcode"
)

// Address - Short address with extra checksum.
type Address struct {
	// ID - Cell id.
	ID opcode.ID `json:"id,omitempty"`

	// CID - Content ID.
	CID *opcode.CID `json:"cid,omitempty"`

	// Extra - Extra checksum.
	Extra uint16 `json:"extra,omitempty"`
}

// ParseAddress - Parses short address from string.
func ParseAddress(body string) (addr *Address, err error) {
	addr = new(Address)
	err = addr.UnmarshalString(body)
	return
}

// MustParseAddress - Parses short address or panics.
func MustParseAddress(src string) (addr *Address) {
	addr, err := ParseAddress(src)
	if err != nil {
		panic(err)
	}
	return
}

// NewAddress - Creates address from bytes.
func NewAddress(bytes []byte) (addr *Address) {
	addr = new(Address)
	addr.SetBytes(bytes)
	return
}

// ToBytes - Creates address from bytes.
func ToBytes(src string) (body []byte, err error) {
	addr, err := ParseAddress(src)
	if err != nil {
		return
	}
	return addr.Marshal()
}

// FromCID - Creates address from content identifier.
func FromCID(c *opcode.CID) (addr *Address) {
	addr = new(Address)
	addr.SetCID(c)
	return
}

// CidToShort - Creates short address from content identifier.
func CidToShort(c *opcode.CID) (addr *Address) {
	addr = new(Address)
	addr.SetBytes(c.Bytes())
	return
}

// IsShortAddress - Returns true if there is no cid available, only short address.
func (addr *Address) IsShortAddress() bool {
	return addr.CID == nil
}

// String - Returns short address in string format.
func (addr *Address) String() string {
	body, err := addr.Marshal()
	if err != nil {
		panic(err)
	}
	body = base32check.CheckEncode(body)
	return string(append([]byte{'b'}, body...))
}

// SetCID - Sets address from cid.
func (addr *Address) SetCID(c *opcode.CID) {
	bytes := c.Bytes()
	addr.ID = opcode.NewID(bytes)
	addr.Extra = uint16(math.Ceil(math.Sqrt(float64(uint64(addr.ID) % uint64(crc32.ChecksumIEEE(bytes))))))
	addr.CID = c
	return
}

// SetBytes - Sets address from bytes.
func (addr *Address) SetBytes(bytes []byte) {
	addr.ID = opcode.NewID(bytes)
	addr.Extra = uint16(math.Ceil(math.Sqrt(float64(uint64(addr.ID) % uint64(crc32.ChecksumIEEE(bytes))))))
}

// Marshal - Marshals address as byte array.
func (addr *Address) Marshal() (_ []byte, err error) {
	buff := proto.NewBuffer(nil)
	if err := buff.EncodeVarint(uint64(addr.ID)); err != nil {
		return nil, err
	}
	if err := buff.EncodeVarint(uint64(addr.Extra)); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// Unmarshal - Unmarshals address from byte array.
func (addr *Address) Unmarshal(body []byte) (err error) {
	buff := proto.NewBuffer(body)
	id, err := buff.DecodeVarint()
	if err != nil {
		return err
	}
	addr.ID = opcode.ID(id)
	checksum, err := buff.DecodeVarint()
	if err != nil {
		return err
	}
	if checksum > math.MaxUint32 {
		return errors.New("checksum too big")
	}
	addr.Extra = uint16(checksum)
	return
}

// MarshalJSON - Marshals address as JSON.
func (addr *Address) MarshalJSON() ([]byte, error) {
	if addr.CID != nil {
		return []byte(fmt.Sprintf("%q", addr.CID.String())), nil
	}
	return []byte(fmt.Sprintf("%q", addr.String())), nil
}

// UnmarshalJSON - Unmarshals address from JSON.
func (addr *Address) UnmarshalJSON(body []byte) (err error) {
	if len(body) < 2 {
		return errors.New("invalid address")
	}
	body = body[1 : len(body)-1]
	if len(body) == 0 {
		return
	}
	return addr.UnmarshalString(string(body))
}

// UnmarshalString - Unmarshals address from string.
func (addr *Address) UnmarshalString(body string) (err error) {
	if len(body) <= 1 {
		return errors.New("address too short")
	}
	if body[0] == 'z' {
		c, err := cid.Parse(body)
		if err != nil {
			return err
		}
		addr.SetCID(opcode.WrapCID(c))
		return nil
	}
	if body[0] != 'b' {
		return fmt.Errorf("invalid codec %x", body[0])
	}
	// remove 'b' byte
	body = body[1:]
	decoded, err := base32check.CheckDecodeString(body)
	if err != nil {
		return
	}
	return addr.Unmarshal(decoded)
}
