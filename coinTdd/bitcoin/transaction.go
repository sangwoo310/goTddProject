package bitcoin

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func CreateRawTransaction() string {
	// 인풋 트랜잭션을  더블 sha256
	inputTxid := "010000000126c07ece0bce7cda0ccd14d99e205f118cde27e83dd75da7b141fe487b5528fb000000008b48304502202b7e37831273d74c8b5b1956c23e79acd660635a8d1063d413c50b218eb6bc8a022100a10a3a7b5aaa0f07827207daf81f718f51eeac96695cf1ef9f2020f21a0de02f01410452684bce6797a0a50d028e9632be0c2a7e5031b710972c2a3285520fb29fcd4ecfb5fc2bf86a1e7578e4f8a305eeb341d1c6fc0173e5837e2d3c7b178aade078ffffffff02b06c191e010000001976a9143564a74f9ddb4372301c49154605573d7d1a88fe88ac00e1f505000000001976a914010966776006953d5567439e5e39f86a0d273bee88ac00000000"

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

	// output index 설정
	//qq := uint8(1)
	outputIndex := []byte(strconv.Itoa(1))
	fmt.Println(outputIndex)
	fmt.Println(string(outputIndex))

	return ""
}

