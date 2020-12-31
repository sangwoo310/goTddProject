package main

import (
	"encoding/hex"
	"github.com/eoscanada/eos-go/btcsuite/btcutil/base58"
	"log"
	"testing"
)

func TestKeyPair(t *testing.T)  {
	qq := "5JQBx9Uzd1sTwxqPqqCV884zpnMmo53f61jiLaM73XLXE2B71Lh"

	ww := base58.Decode(qq)

	log.Println(hex.EncodeToString(ww))

	qqq := []byte{42, 74, 94, 158, 04, 93, 158, 213, 48, 87, 77, 251, 69, 243, 234, 185, 24, 176, 186, 175, 202, 95, 184, 44, 175, 252, 157, 50, 243, 38, 158, 142}

	str := hex.EncodeToString(qqq)
	log.Println(str)


	tt := 6<<3
	log.Println(tt)
}

