package bitcoin

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/eoscanada/eos-go/btcsuite/btcutil/base58"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ripemd160"
	"log"
	"testing"
)

func TestPubKeyToAddress(t *testing.T) {
	pubKey := "0242e7d8af9aa664e7f050bab81a0087151d55d80653368f7758c9c4c0f40b6c3a"
	//pubKey := "0250863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b2352"
	pubKeyByte, err := hex.DecodeString(pubKey)
	assert.NoError(t, err)
	log.Println("pubKeyByte : ", pubKeyByte)

	// one sha256
	s := sha256.New()
	s.Reset();
	_, err = s.Write(pubKeyByte)
	assert.NoError(t, err)
	hash1 := s.Sum(nil)
	log.Println("hash1 : ", hash1)
	log.Println("hash1 : ", hex.EncodeToString(hash1))

	r := ripemd160.New()
	r.Reset();
	_, err = r.Write(hash1)
	assert.NoError(t, err)
	ripemd160Addr := r.Sum(nil)
	log.Println("ripemd160Addr : ", ripemd160Addr)
	log.Println("ripemd160Addr : ", hex.EncodeToString(ripemd160Addr))

	deB58 := make([]byte, 0)
	deB58 = append(deB58, byte(0))
	deB58 = append(deB58, ripemd160Addr...)
	log.Println(deB58)
	log.Println("ripemdVersionAdd : ", hex.EncodeToString(deB58))

	// checksum sha256 1
	s.Reset();
	_, err = s.Write(deB58)
	//_, err = s.Write(ripemd160Addr)
	assert.NoError(t, err)
	checkSha1 := s.Sum(nil)
	log.Println("checkSha1 : ", hex.EncodeToString(checkSha1))

	// double sha256 2
	s.Reset()
	_, err = s.Write(checkSha1)
	assert.NoError(t, err)
	checkSha2 := s.Sum(nil)
	log.Println("checkSha2 : ",  checkSha2)
	log.Println("checkSha2 : ", hex.EncodeToString(checkSha2))
	log.Println("cheksum : ", checkSha2[:len(checkSha2)-4])
	log.Println("cheksum : ", hex.EncodeToString(checkSha2[:4]))

	deB58 = append(deB58, checkSha2[:4]...)
	addr := base58.Encode(deB58)
	log.Println("addr : ",  addr)
	//log.Println("addr : ", hex.EncodeToString(addr))

}
