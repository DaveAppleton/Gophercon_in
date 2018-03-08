package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func GetMD5Hash(text string) {
	hasher := md5.New()
	hasher.Write([]byte(text))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
}

//start main OMIT

func GetKeccak256Hash(test string) {
	h := crypto.Keccak256Hash([]byte(test))
	fmt.Println(h.Hex())
}

func main() {
	GetKeccak256Hash("Young man, in mathematics you don't understand things. You just get used to them.")
	GetKeccak256Hash("Young man. in mathematics you don't understand things. You just get used to them.")
}

//end main OMIT
