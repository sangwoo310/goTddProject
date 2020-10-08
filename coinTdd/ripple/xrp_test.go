package ripple

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

const (
	from         = "rsTwerzJEGiKh7WjJcC3Q7776D4eGvDXPz"
	to           = "rG5AB117rJ7e2MZGKE4XfaVK5BdyHBxcSm"
	from_private = "c1dd3d6ba77aaa58a15f706e54c4d1dd59b2f61fdb621fd2dd80e204f0eaa2dc"
	to_private   = "c57891a6f2212dd312a12cb9323e69b6ad8a0faaf8435ca533876a7c12b80ae8"
)
/*
func TestGoXrpTest(t *testing.T) {
	Account := "rsTwerzJEGiKh7WjJcC3Q7776D4eGvDXPz"
	Amount := "30"
	Destination := "rG5AB117rJ7e2MZGKE4XfaVK5BdyHBxcSm"
	Fee := "12"
	Flags := 2147483648
	last := uint32(13313150)
	Sequence := 1
	//TxnSignature := "3045022100D59891D15129AFA2297506207AF14A97C2C236C690BA5E167E84BC070CA3774202203F80DFC3D8965AA4705940B9233ED8570812557F1E9DC011DEAF47DC2AE8BD58"
	SigningPubKey := "028C35EEA94EE7FA9C8485426E164159330BA2453368F399669D5110009F270EE9"

	fromAccount, _ := data.NewAccountFromAddress(Account)
	toAccount, _ := data.NewAccountFromAddress(Destination)
	amount, _ := data.NewAmount(Amount + "/XRP")
	fee, _ := data.NewValue(Fee, true)
	flags := data.TransactionFlag(Flags)
	//tSig, _ := hex.DecodeString(TxnSignature)
	//txnSign := data.VariableLength(tSig)
	signPubKey := data.PublicKey{}
	pk, _ := hex.DecodeString(SigningPubKey)
	copy(signPubKey[:], pk)

	txn := data.TxBase{
		TransactionType:    data.PAYMENT,
		Account:            *fromAccount,
		LastLedgerSequence: &last,
		Flags:              &flags,
		Sequence:           uint32(Sequence),
		//TxnSignature:       &txnSign,
		Fee:           *fee,
		SigningPubKey: &signPubKey,
	}
	payment := data.Payment{
		TxBase:      txn,
		Amount:      *amount,
		Destination: *toAccount,
	}
	fmt.Println(payment)

	//res, err := client.makeTxBlob(&payment)
	//if err != nil {
	//	t.Error("gen blob err: ", err)
	//}
	//t.Log("tx blog: ", res)
}
*/

func TestSerialize(t *testing.T) {
	type Pays struct {
		Currency	string
		Issuer		string
		Value		string
	}

	type XrpJson struct {
		Account			string
		Expiration		uint64
		Fee				string
		Flags			uint64
		OfferSequence	uint64
		Sequence		uint64
		SigningPubKey	string
		TakerGets		string
		TakerPays		Pays
		TransactionType	string
		TxnSignature	string
		Hash			string
	}

	st := XrpJson{
		Account:         "rMBzp8CgpE441cp5PVyA9rpVV7oT8hP3ys",
		Expiration:      595640108,
		Fee:             "10",
		Flags:           524288,
		OfferSequence:   1752791,
		Sequence:		 1752792,
		SigningPubKey:   "03EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3",
		TakerGets:       "15000000000",
		TransactionType: "OfferCreate",
		TxnSignature:    "30440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C",
		Hash:            "73734B611DDA23D3F5F62E20A173B78AB8406AC5015094DA53F53D39B9EDB06C",
	}
	st.TakerPays.Value = "7072.8"
	st.TakerPays.Currency = "USD"
	st.TakerPays.Issuer = "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"

	qq, _ := json.Marshal(st)
	fmt.Println(hex.EncodeToString(qq))
}

func TestSQQQ(t *testing.T) {
	hexStr := "Payment"
	//qq,_ := hex.DecodeString(hexStr)
	//fmt.Println(string(qq))
	byteStr := [] byte(hexStr)
	fmt.Println(byteStr)
	fmt.Println(hex.EncodeToString(byteStr))
}