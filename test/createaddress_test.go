package test

import (
	"encoding/hex"
	"fmt"
	"github.com/LwwL-123/go-substrate-crypto/ss58"
	"testing"
)

func Test_CreateAddress(t *testing.T) {
	pub := "a69958eee5de0cb8fb250eba9c4b4ab1675468e68e49a5ebcac22fa9340fe938"
	pubBytes, _ := hex.DecodeString(pub)
	fmt.Println(pubBytes)
	//pubBytes = append([]byte{0xff},pubBytes...)
	address, err := ss58.Encode(pubBytes, ss58.SubstratePrefix)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(address)
	addresses := "5Fq9MpKxdjzCWEHHtqZ6rdYkKUtW4qwmJV4VHwKBan2hxRyL"
	b, err := ss58.Decode(addresses)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(b[1:]))
	fmt.Println(b)
}

func Test_DecodeChainXAddress(t *testing.T) {

}
