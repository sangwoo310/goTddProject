package ethereum

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
	"testing"
)
/*
const AddressLength = 20

type EthTx struct {
	AccountNonce uint64         `json:"nonce"    gencodec:"required"`
	Price        *big.Int       `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64         `json:"gas"      gencodec:"required"`
	Recipient    *Address 		`json:"to"       rlp:"nil"` // nil means contract creation
	Amount       *big.Int       `json:"value"    gencodec:"required"`
	Payload      []byte         `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	//Hash *common.Hash `json:"hash" rlp:"-"`
}

type txdataMarshaling struct {
	AccountNonce uint64          `json:"nonce"    gencodec:"required"`
	Price        *big.Int        `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64          `json:"gas"      gencodec:"required"`
	//Recipient    *common.Address `json:"to"       rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"    gencodec:"required"`
	Payload      []byte          `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
	// This is only used when marshaling to JSON.
	//Hash *common.Hash `json:"hash" rlp:"-"`
}

type Address [AddressLength]byte

type encbuf struct {
	str     []byte      // string data, contains everything except list headers
	lheads  []*listhead // all list headers
	lhsize  int         // sum of sizes of all encoded list headers
	sizebuf []byte      // 9-byte auxiliary buffer for uint encoding
}

type listhead struct {
	offset int // index of this header in string data
	size   int // total size of encoded data (including list headers)
}

type txdata struct {
	AccountNonce uint64          `json:"nonce"    gencodec:"required"`
	Price        *big.Int        `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64          `json:"gas"      gencodec:"required"`2
	//Recipient    *common.Address `json:"to"       rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"    gencodec:"required"`
	Payload      []byte          `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	//Hash *common.Hash `json:"hash" rlp:"-"`
}*/
/*
tt := txdata{
	AccountNonce: nonce,
	Recipient:    to,
	Payload:      data,
	Amount:       new(big.Int),
	GasLimit:     gasLimit,
	Price:        new(big.Int),
	V:            new(big.Int),
	R:            new(big.Int),
	S:            new(big.Int),
}
*/
/*
func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}

func FromHex(s string) []byte {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return Hex2Bytes(s)
}

func HexToAddress(s string) Address { return BytesToAddress(FromHex(s)) }

func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

func TestRlp(t *testing.T)  {
	addr := HexToAddress("0x687422eEA2cB73B5d3e424bA5456b782919AFc85")
	fmt.Println(addr)

	e := EthTx{
		AccountNonce: 0,
		Price:        big.NewInt(1000000),
		GasLimit:     21000,
		Recipient:	  &addr,
		Amount:       big.NewInt(1000000000000000000
		Payload:      nil,
		//V:            big.NewInt(0x1c),
		//R:            big.NewInt(0x668ed6500efd75df7cb9c9b9d8152292a75453ec2d11030b0eec42f6a7ace602),
		//S:            big.NewInt(0x3efcbbf4d53e0dfa4fde6c6d9a73221418652abd66dff7fdd78b8fcc28b9fbf),
	}
	//w := encbuf{}
	//err := rlp.Encode(w,&e)
	fmt.Println(e)
}
*/
func TestQqqqqq(t *testing.T) {
	qq := Tx()
	fmt.Println(qq)
}

func TestRlpDecode(t *testing.T) {
	unsignedTx := "e9808203e8825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a764000080808080"
	byteData, err := hex.DecodeString(unsignedTx)
	if err != nil {
		log.Panic(err)
	}

	var q interface{}
	if err := rlp.DecodeBytes(byteData, &q); err != nil {
		log.Panic(err)
	}
	fmt.Println(q)
	fmt.Println()

	data, err := json.Marshal(q)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(hex.EncodeToString(data))
	fmt.Println(string(data))
}