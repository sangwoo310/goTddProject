package tron

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/shengdoushi/base58"
	"testing"
)

func TestAddressForm(t *testing.T) {
	// bas64 -> double sha256 4byte -> base64 + double sha256[4:] -> base58encoding
	addr := "419e98e60b6bad2fa2544169e0dbb9a15288501127"
	hexAddr, _ := hex.DecodeString(addr)

	s := sha256.New()
	s.Reset()
	_, err := s.Write(hexAddr)
	if err != nil {
		fmt.Println(err)
	}
	hash1 := s.Sum(nil)

	s.Reset()
	_, err = s.Write(hash1)
	if err != nil {
		fmt.Println(err)
	}
	hash2 := s.Sum(nil)

	checkSum := hash2[:4]
	checkAddr := hex.EncodeToString(checkSum)

	lastAddr := addr + checkAddr
	hexLast, _ := hex.DecodeString(lastAddr)

	alphabet  := base58.BitcoinAlphabet
	qq := base58.Encode(hexLast, alphabet)
	fmt.Println(qq)
}

func TestAddressForm2(t *testing.T) {
	addr := "TGPnfvkVkUdyCWmrVGnQ9jFW5Ca1S7XH6h"
	//hexAddr, _ := hex.DecodeString(addr)

	alphabet  := base58.BitcoinAlphabet
	qq, _ := base58.Decode(addr, alphabet)
	fmt.Println(hex.EncodeToString(qq))
}

func TestCreateTx(t *testing.T) {

}