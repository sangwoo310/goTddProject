package eos

import (
	"encoding/hex"
	"encoding/json"
	"os"

	"fmt"
	"testing"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/token"
)



///*
///**
//	struct model
//*/
//type Name string
//type AccountName Name
//type PermissionName Name
//type ActionName Name
//type HexBytes []byte
//
//// Action
//type Action struct {
//	Account       AccountName       `json:"account"`
//	Name          ActionName        `json:"name"`
//	Authorization []PermissionLevel `json:"authorization,omitempty"`
//	ActionData
//}
//
//type ActionData struct {
//	HexData  HexBytes    `json:"hex_data,omitempty"`
//	Data     interface{} `json:"data,omitempty" eos:"-"`
//	abi      []byte      // TBD: we could use the ABI to decode in obj
//	toServer bool
//}
//
//type PermissionLevel struct {
//	Actor      AccountName    `json:"actor"`
//	Permission PermissionName `json:"permission"`
//}
//
//func NewActionData(obj interface{}) ActionData {
//	return ActionData{
//		HexData:  []byte{},
//		Data:     obj,
//		toServer: true,
//	}
//}
//
//
//
//
//type JSONTime struct {
//	time.Time
//}
//type Varuint32 uint32
//
//type TransactionHeader struct {
//	Expiration     JSONTime `json:"expiration"`
//	RefBlockNum    uint16   `json:"ref_block_num"`
//	RefBlockPrefix uint32   `json:"ref_block_prefix"`
//
//	MaxNetUsageWords Varuint32 `json:"max_net_usage_words"`
//	MaxCPUUsageMS    uint8     `json:"max_cpu_usage_ms"`
//	DelaySec         Varuint32 `json:"delay_sec"` // number of secs to delay, making it cancellable for that duration
//}
//
//type Transaction struct { // WARN: is a `variant` in C++, can be a SignedTransaction or a Transaction.
//	TransactionHeader
//
//	ContextFreeActions []*Action    `json:"context_free_actions"`
//	Actions            []*Action    `json:"actions"`
//	Extensions         []*Extension `json:"transaction_extensions"`
//}
//
//type Extension struct {
//	Type uint16
//	Data HexBytes
//}
//
//
//
//
//type Int64 int64
//type Symbol struct {
//	Precision uint8
//	Symbol    string
//
//	// Caching of symbol code if it was computed once
//	symbolCode uint64
//}
//
//type Asset struct {
//	Amount Int64
//	Symbol
//}
//
//
//
//type Checksum256 []byte
//type CompressionType uint8
//
//type TxOptions struct {
//	ChainID          Checksum256 // If specified, we won't hit the API to fetch it
//	HeadBlockID      Checksum256 // If provided, don't hit API to fetch it.  This allows offline transaction signing.
//	MaxNetUsageWords uint32
//	DelaySecs        uint32
//	MaxCPUUsageMS    uint8 // If you want to override the CPU usage (in counts of 1024)
//	//ExtraKCPUUsage uint32 // If you want to *add* some CPU usage to the estimated amount (in counts of 1024)
//	Compress CompressionType
//}
//
////func (opts *TxOptions) FillFromChain(api *API) error {
////	if opts == nil {
////		return errors.New("TxOptions should not be nil, send an object")
////	}
////
////	if opts.HeadBlockID == nil || opts.ChainID == nil {
////		info, err := api.cachedGetInfo()
////		if err != nil {
////			return err
////		}
////
////		if opts.HeadBlockID == nil {
////			opts.HeadBlockID = info.HeadBlockID
////		}
////		if opts.ChainID == nil {
////			opts.ChainID = info.ChainID
////		}
////	}
////
////	return nil
////}
//
///**
//	registerAction
// */
//
////var RegisteredActions = map[AccountName]map[ActionName]reflect.Type{}
////func RegisterAction(accountName AccountName, actionName ActionName, obj interface{}) {
////	// TODO: lock or som'th.. unless we never call after boot time..
////	if RegisteredActions[accountName] == nil {
////		RegisteredActions[accountName] = make(map[ActionName]reflect.Type)
////	}
////	RegisteredActions[accountName][actionName] = reflect.TypeOf(obj)
////}
//
//func AN(in string) AccountName    { return AccountName(in) }
//func ActN(in string) ActionName   { return ActionName(in) }
//func PN(in string) PermissionName { return PermissionName(in) }
//
////var AN = AN
////var PN = PN
////var ActN = ActN
//
///**
//	transfer code
// */
//func NewTransfer(from, to AccountName, quantity Asset, memo string) *Action {
//	qq := &Action{
//		Account:       AN("eosio.token"),
//		Name:          ActN("transfer"),
//		Authorization: []PermissionLevel{
//			{Actor: from, Permission: PN("active")},
//		},
//		ActionData:    NewActionData(Transfer{
//			From:     from,
//			To:       to,
//			Quantity: quantity,
//			Memo:     memo,
//		}),
//	}
//	return qq
//	//return &Action{
//	//	Account: AN("eosio.token"),
//	//	Name:    ActN("transfer"),
//	//	Authorization: []eos.PermissionLevel{
//	//		{Actor: from, Permission: PN("active")},
//	//	},
//	//	ActionData: eos.NewActionData(Transfer{
//	//		From:     from,
//	//		To:       to,
//	//		Quantity: quantity,
//	//		Memo:     memo,
//	//	}),
//	//}
//}
//
//// Transfer represents the `transfer` struct on `eosio.token` contract.
//type Transfer struct {
//	From     AccountName `json:"from"`
//	To       AccountName `json:"to"`
//	Quantity Asset       `json:"quantity"`
//	Memo     string          `json:"memo"`
//}

//
///**
//	transaction code
// */
//func NewTransaction(actions []*Action, opts *TxOptions) *Transaction {
//	if opts == nil {
//		opts = &TxOptions{}
//	}
//
//	tx := &Transaction{Actions: actions}
//	tx.Fill(opts.HeadBlockID, opts.DelaySecs, opts.MaxNetUsageWords, opts.MaxCPUUsageMS)
//	return tx
//}
//
//func (tx *Transaction) Fill(headBlockID Checksum256, delaySecs, maxNetUsageWords uint32, maxCPUUsageMS uint8) {
//	tx.setRefBlock(headBlockID)
//
//	if tx.ContextFreeActions == nil {
//		tx.ContextFreeActions = make([]*Action, 0, 0)
//	}
//	if tx.Extensions == nil {
//		tx.Extensions = make([]*Extension, 0, 0)
//	}
//
//	tx.MaxNetUsageWords = Varuint32(maxNetUsageWords)
//	tx.MaxCPUUsageMS = maxCPUUsageMS
//	tx.DelaySec = Varuint32(delaySecs)
//
//	tx.SetExpiration(30 * time.Second)
//}
//
//func (tx *Transaction) SetExpiration(in time.Duration) {
//	tx.Expiration = JSONTime{time.Now().UTC().Add(in)}
//}
//
//func (tx *Transaction) setRefBlock(blockID []byte) {
//	if len(blockID) == 0 {
//		return
//	}
//	tx.RefBlockNum = uint16(binary.BigEndian.Uint32(blockID[:4]))
//	tx.RefBlockPrefix = binary.LittleEndian.Uint32(blockID[8:16])
//}
//
//
///**
//   transaction sample code
//*/
//
//func TestSample(t *testing.T) {
//	//from := AccountName(Name("fromAddr"))
//	from := AN("fromAddr")
//	to := AN("toAddr")
//	quantity := Asset{
//		Amount: 100,
//		Symbol: Symbol{
//			Precision:  0,
//			Symbol:     "EOS",
//			symbolCode: 0,
//		},
//	}
//	memo := "memo"
//
//	txOpts := &TxOptions{}
//	txOpts = nil
//	//if err := txOpts.FillFromChain(api); err != nil {
//	//	panic(fmt.Errorf("filling tx opts: %s", err))
//	//}
//
//	tx := NewTransaction([]*Action{NewTransfer(from, to, quantity, memo)}, txOpts)
//
//	fmt.Println(tx)
//	fmt.Println(tx.Actions[0])
//
//
//
//
//	signedTx, packedTx, err := api.SignTransaction(tx, txOpts.ChainID, eos.CompressionNone)
//	if err != nil {
//		panic(fmt.Errorf("sign transaction: %s", err))
//	}
//
//	content, err := json.MarshalIndent(signedTx, "", "  ")
//	if err != nil {
//		panic(fmt.Errorf("json marshalling transaction: %s", err))
//	}
//
//	fmt.Println(string(content))
//	fmt.Println()
//
//	response, err := api.PushTransaction(packedTx)
//	if err != nil {
//		panic(fmt.Errorf("push transaction: %s", err))
//	}
//
//	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(response.Processed.ID))
//}
//*/


func getAPIURL() string {
	apiURL := os.Getenv("EOS_GO_API_URL")
	if apiURL != "" {
		return apiURL
	}

	return "https://mainnet.eoscanada.com"
}

func readPrivateKey() string {
	// Right now, the key is read from an environment variable, it's an example after all.
	// In a real-world scenario, would you probably integrate with a real wallet or something similar
	envName := "EOS_GO_PRIVATE_KEY"
	//privateKey := os.Getenv(envName)
	//privateKey := "5HqShoQr8Scf9z1qBpnSnse1xyNKqUNfrsbb9GX5n5Ln7C6VH4N"
	privateKey := "5KKh6vJPGmztXFeJNFQidCjZjtKVuTBRJg21VD9xT25pxNEyuat"		// 진짜 pk 사용주의!!!
	if privateKey == "" {
		panic(fmt.Errorf("private key environment variable %q must be set", envName))
	}

	return privateKey
}

//pk := "5HqShoQr8Scf9z1qBpnSnse1xyNKqUNfrsbb9GX5n5Ln7C6VH4N"
func TestTx(t *testing.T) {
	api := eos.New(getAPIURL())

	keyBag := &eos.KeyBag{}
	err := keyBag.ImportPrivateKey(readPrivateKey())
	if err != nil {
		panic(fmt.Errorf("import private key: %s", err))
	}
	api.SetSigner(keyBag)

	from := eos.AccountName("t3st1ngst4g3")
	to := eos.AccountName("eosuser2")
	quantity, err := eos.NewEOSAssetFromString("0.0001 EOS")
	memo := ""

	if err != nil {
		panic(fmt.Errorf("invalid quantity: %s", err))
	}

	txOpts := &eos.TxOptions{}
	if err := txOpts.FillFromChain(api); err != nil {
		panic(fmt.Errorf("filling tx opts: %s", err))
	}

	//v := types.Interface{}
	tx := eos.NewTransaction([]*eos.Action{token.NewTransfer(from, to, quantity, memo)}, txOpts)

	qq, err := json.Marshal(tx)
	fmt.Println("qq : ", string(qq))
	fmt.Println("tx : ", tx)
	signedTx, packedTx, err := api.SignTransaction(tx, txOpts.ChainID, eos.CompressionNone)
	if err != nil {
		panic(fmt.Errorf("sign transaction: %s", err))
	}

	fmt.Println("signedTx : ", signedTx)
	fmt.Println("packedTx : ", packedTx)

	content, err := json.MarshalIndent(signedTx, "", "  ")
	if err != nil {
		panic(fmt.Errorf("json marshalling transaction: %s", err))
	}

	fmt.Println(string(content))
	fmt.Println()

	response, err := api.PushTransaction(packedTx)
	if err != nil {
		panic(fmt.Errorf("push transaction: %s", err))
	}

	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(response.Processed.ID))
}