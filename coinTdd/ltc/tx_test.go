package ltc

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/shengdoushi/base58"
	"strconv"
	"testing"
)

func TestUtxoIndex(t *testing.T)  {
	var b bytes.Buffer
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], 2)
	b.Write(buf[:])

	fmt.Println(b)
	fmt.Println(hex.EncodeToString(b.Bytes()))
}

func TestHash32(t *testing.T) {
	// 인풋 트랜잭션을  더블 sha256
	inputTxid := "010000000126c07ece0bce7cda0ccd14d99e205f118cde27e83dd75da7b141fe487b5528fb000000008b48304502202b7e37831273d74c8b5b1956c23e79acd660635a8d1063d413c50b218eb6bc8a022100a10a3a7b5aaa0f07827207daf81f718f51eeac96695cf1ef9f2020f21a0de02f01410452684bce6797a0a50d028e9632be0c2a7e5031b710972c2a3285520fb29fcd4ecfb5fc2bf86a1e7578e4f8a305eeb341d1c6fc0173e5837e2d3c7b178aade078ffffffff02b06c191e010000001976a9143564a74f9ddb4372301c49154605573d7d1a88fe88ac00e1f505000000001976a914010966776006953d5567439e5e39f86a0d273bee88ac00000000"
	//inputTxid := "ec7f7c9ae9519d6559a0f2c3b2e540d397d7e9f1"
	byteTx := []byte(inputTxid)
	fmt.Println(byteTx)

	//var b bytes.Buffer
	//var buf [2]byte
	//binary.LittleEndian.PutUint32(buf[:], hash2)
	//b.Reset()
	//b.Write([]byte(inputTxid))
	//
	//fmt.Println(len(inputTxid))


	//hash := []byte(inputTxid)
	hash, err := hex.DecodeString(inputTxid)
	if err != nil {
		fmt.Println("!!!!!!!!")
		fmt.Println(err)
	}

	fmt.Println("***********************************************")
	fmt.Println(hash)
	fmt.Println("***********************************************")

	s := sha256.New()
	s.Reset()
	s.Write(hash)
	inputHash1 := s.Sum(nil)

	s.Reset()
	s.Write(inputHash1)
	inputHash2 := s.Sum(nil)

	fmt.Println(inputHash2)
	fmt.Println(hex.EncodeToString(inputHash2))
	fmt.Println(hex.EncodeToString(inputHash1))
}

func TestBtcDeAddr(t *testing.T)  {
	addr := "3PFW7kS1qpKL4pTYMfw4AWRyrVhqdgcqBs"

	alphabet  := base58.BitcoinAlphabet
	byteAddr, err := base58.Decode(addr, alphabet)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(byteAddr)
	fmt.Println(len(byteAddr))

	fmt.Println(hex.EncodeToString(byteAddr))

	s := sha256.New()
	s.Reset()
	s.Write(byteAddr)
	inputHash1 := s.Sum(nil)

	s.Reset()
	s.Write(inputHash1)
	inputHash2 := s.Sum(nil)

	fmt.Println("***********************************************")
	fmt.Println(hex.EncodeToString(inputHash2))
}


func TestReverse(t *testing.T) {
	txid := "3dbf1a4c5b7d1693798933022dc8185348a404af96d4e423c405f7190259c516"

	byteTx := []byte(txid)

	fmt.Println(len(txid))
	fmt.Println(len(byteTx))
	fmt.Println(byteTx)

	fmt.Println("***********************************************")

	//binary.Uvarint()
	//version := 0x02000000
	//
	//qq := hex.DecodedLen(version)
	//fmt.Println(qq)

	var b bytes.Buffer
	//var buf [4]byte

	b.Write(byteTx[:])
	fmt.Println(b)
	fmt.Println(hex.EncodeToString(b.Bytes()))
}

func TestEndian(t *testing.T) {
	value := hex.DecodedLen(10000)
	fmt.Println(value)
	s4 := strconv.FormatUint(100000, 16)
	fmt.Println(s4)

	var b bytes.Buffer
	var buf [8]byte
	binary.LittleEndian.PutUint32(buf[:], 0x186a0)
	b.Write(buf[:])

	fmt.Println(b)
	fmt.Println(hex.EncodeToString(b.Bytes()))
}


func TestHash160(t *testing.T) {
	data := "ec7f7c9ae9519d6559a0f2c3b2e540d397d7e9f1"

	tt := []byte(data)

	hash, err := hex.DecodeString(data)
	if err != nil {
		fmt.Println("!!!!!!!!")
		fmt.Println(err)
	}
	fmt.Println(hash)

	qq := btcutil.Hash160(tt)
	fmt.Println(qq)
	fmt.Println(hex.EncodeToString(qq))
	fmt.Println(len(hex.EncodeToString(qq)))
}

func TestByte(t *testing.T) {
	var b bytes.Buffer
	//var q [1]byte
	//q[0] = 17
	//b.Write(q[:])
	//fmt.Println(hex.EncodeToString(b.Bytes()))

	inputCount := byte(17)
	var q [1]byte
	q[0] = inputCount
	b.Reset()
	b.Write(q[:])
	fmt.Println(hex.EncodeToString(b.Bytes()))
}

func TestLen(t *testing.T) {
	q := "76a91421b2ed1347b21a1046e5b6577fb5244752f7ece888ac"
	fmt.Println(len(q))
}

func TestLtcTxCreator(t *testing.T) {
	// version set
	version := uint32(2)
	var b bytes.Buffer
	var versionBuf [4]byte
	binary.LittleEndian.PutUint32(versionBuf[:], version)
	b.Write(versionBuf[:])
	versionTx := hex.EncodeToString(b.Bytes())
	fmt.Println(versionTx)

	// input tx count
	inputCount := 1
	var inputCountByte [1]byte
	inputCountByte[0] = byte(inputCount)
	b.Reset()
	b.Write(inputCountByte[:])
	inputCountTx := hex.EncodeToString(b.Bytes())
	fmt.Println(inputCountTx)

	// first txid
	inputTxid := "1025c758203a9f4aab1ca4b9faa772acaf162faf6d5c0a312f6fec3b8c5e84ff"
	inputTxidArr := make([]string, 0)
	for i:=0; i<len(inputTxid)/2; i++ {
		inputTxidArr = append(inputTxidArr, inputTxid[i*2 : i*2+2])
	}
	firstInputTx := ""
	for i:=len(inputTxidArr); i>0; i-- {
		firstInputTx += inputTxidArr[i-1]
	}
	fmt.Println(firstInputTx)

	//first txid output index
	inputIndex := 2
	var inIndexByte [4]byte
	inIndexByte[0] = byte(inputIndex)
	b.Reset()
	b.Write(inIndexByte[:])
	inputIndexTx := hex.EncodeToString(b.Bytes())
	fmt.Println(inputIndexTx)


	// 임시 scriptsig 길이 ==> 00
	tempScriptSigLen := "00"
	fmt.Println(tempScriptSigLen)

	// 4byte sequence field ==> ffffffff
	sequenceField := "ffffffff"
	fmt.Println(sequenceField)


	// output tx count
	outputCount := 1
	var outputCountByte [1]byte
	outputCountByte[0] = byte(outputCount)
	b.Reset()
	b.Write(outputCountByte[:])
	outputCountTx := hex.EncodeToString(b.Bytes())
	fmt.Println(outputCountTx)

	// amount 8byte satoshis
	amount := 12.345
	value := uint64(amount * 100000000)
	var valueByte [8]byte
	binary.LittleEndian.PutUint64(valueByte[:], value)
	b.Reset()
	b.Write(valueByte[:])
	valueTx := hex.EncodeToString(b.Bytes())
	fmt.Println(valueTx)

	// output addr => hash160 => scriptpubkey
	scriptPubKey := ""

	//outputAddr := "LTCPodtg91HsKf5xB1MQ5L13CkcSfd2rCV"
	//outputAddr := "MTpp5c4DpxMihE3EMu86vrbjfWPhvsnG3q"
	outputAddr := "ltc1qttqv5gk8rg57a32u8yurrdnlfa0sfff6ymkh6qlkznckgya7ftkqctch4r"
	var deOutAddrStr string
	if outputAddr[:4] == "ltc1" {
		_, decoded, err := bech32.Decode(outputAddr)
		if err != nil {
			fmt.Println("Error:", err)
		}
		length := len(decoded)
		decoded, _ = hex.DecodeString(hex.EncodeToString(decoded[1:length-4]))
		deBeAddr, err := bech32.ConvertBits(decoded, 5, 8, false)
		deOutAddrStr = hex.EncodeToString(deBeAddr)
	} else {
		alphabet  := base58.BitcoinAlphabet
		byteAddr, err := base58.Decode(outputAddr, alphabet)
		if err != nil {
			fmt.Println(err)
		}
		length := len(byteAddr)
		deOutAddrStr = hex.EncodeToString(byteAddr[1:length-4])
		fmt.Println("outputAddr : ", deOutAddrStr)
	}

	switch {
		case outputAddr[:4] == "ltc1" :
			// ltc bech32 address
			scriptPubKey = "0020" + deOutAddrStr + "4aec"
			break
		case outputAddr[:1] == "1" || outputAddr[:1] == "L" :
			// op_dup = 76, op_hash160 = a9, script_length = 14, op_equalverify = 88, op_checksig = ac
			scriptPubKey = "76a914" + deOutAddrStr + "88ac"
			break
		case outputAddr[:1] == "3" || outputAddr[:1] == "M" :
			// op_hash160 = a9, script_length = 14, op_equal = 87
			// M is LTC address
			scriptPubKey = "a914" + deOutAddrStr + "87"
			break
		case outputAddr[:1] == "b" :
			// 0
			scriptPubKey = "00" + deOutAddrStr
			break
	}
	// output scriptpubkey length(0x) => before outAddrTx
	scriptPubKeyLen := strconv.FormatInt(int64(len(scriptPubKey)/2), 16)
	outAddrTx := scriptPubKeyLen + scriptPubKey
	fmt.Println(outAddrTx)

	// 4byte lockTime ==> 00000000
	lockTime := "00000000"
	fmt.Println(lockTime)

	fmt.Println("rawTx : ", versionTx + inputCountTx + firstInputTx + inputIndexTx + tempScriptSigLen + sequenceField + outputCountTx + valueTx + outAddrTx + lockTime)
}

func TestBechLTC(t *testing.T) {
	//outputAddr := "bitcoincash:qqd8r9a36plmm9xgjsd3pkpuv0mgs4mjzuy4lme4s5"
	//outputAddr := "ltc1:qttqv5gk8rg57a32u8yurrdnlfa0sfff6ymkh6qlkznckgya7ftkqctch4r"
	//byteAddr, prefix, addrType, err := utils.CheckDecodeCashAddress(outputAddr)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(prefix)
	//fmt.Println(addrType)
	//deOutAddrStr := hex.EncodeToString(byteAddr)
	//fmt.Println(deOutAddrStr)
	//encoded := "bc1pw508d6qejxtdg4y5r3zarvary0c5xw7kw508d6qejxtdg4y5r3zarvary0c5xw7k7grplx"
	encoded := "ltc1qttqv5gk8rg57a32u8yurrdnlfa0sfff6ymkh6qlkznckgya7ftkqctch4r"
	hrp, decoded, err := bech32.Decode(encoded)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Show the decoded data.
	fmt.Println("Decoded human-readable part:", hrp)
	fmt.Println("Decoded Data:", hex.EncodeToString(decoded))
	length := len(decoded)
	decoded, _ = hex.DecodeString(hex.EncodeToString(decoded[1:length-4]))

	qqqq, err := bech32.ConvertBits(decoded, 5, 8, false)
	fmt.Println(hex.EncodeToString(qqqq))

	fmt.Println("Decoded Data:", hex.EncodeToString(decoded[1:length-4]))

}