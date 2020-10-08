package module

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
	"io"
	"log"
	"math/big"
)

type Transaction struct {
	data txdata
	// caches
	//hash atomic.Value
	//size atomic.Value
	//from atomic.Value
}

type txdata struct {
	AccountNonce uint64          `json:"nonce"    gencodec:"required"`
	Price        *big.Int        `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64          `json:"gas"      gencodec:"required"`
	Recipient    *Address `json:"to"       rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"    gencodec:"required"`
	Payload      []byte          `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	//Hash *common.Hash `json:"hash" rlp:"-"`
}

//type txdata struct {
//	AccountNonce uint64          `json:"nonce"    `
//	Price        *big.Int        `json:"gasPrice" `
//	GasLimit     uint64          `json:"gas"      `
//	Recipient    *Address 		 `json:"to"       ` // nil means contract creation
//	Amount       *big.Int        `json:"value"    `
//	Payload      []byte          `json:"input"    `
//
//	// Signature values
//	V *big.Int `json:"v"`
//	R *big.Int `json:"r"`
//	S *big.Int `json:"s"`
//
//	// This is only used when marshaling to JSON.
//	//Hash *common.Hash `json:"hash" rlp:"-"`
//}

func NewTransaction(nonce uint64, to Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *Transaction {
	return newTransaction(nonce, &to, amount, gasLimit, gasPrice, data)
}

func newTransaction(nonce uint64, to *Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *Transaction {
	//if len(data) > 0 {
	//	data = common.CopyBytes(data)
	//}
	d := txdata{
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
	if amount != nil {
		d.Amount.Set(amount)
	}
	if gasPrice != nil {
		d.Price.Set(gasPrice)
	}

	return &Transaction{data: d}
}



type Transactions []*Transaction

func (s Transactions) GetRlp(i int) []byte {
	enc, _ := rlp.EncodeToBytes(s[i])
	return enc
}

func (s Transactions) DecodeRlp(b []byte) {
	var q interface{}
	fmt.Println("passed data : ",b)
	if err := rlp.DecodeBytes(b, q); err != nil {
		fmt.Println("!!!!!!!!!!!!!err")
		log.Panic(err)
	}
	fmt.Println(q)
}


func (tx *Transaction) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &tx.data)
}
