package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/crypto"
)

//start main OMIT

func PrintSig(test string) {
	h := crypto.Keccak256Hash([]byte(test))
	fmt.Println(test, common.Bytes2Hex(h[:4]))
}

func main() {
	PrintSig("transfer(address,uint256)")
	PrintSig("hitMe()")
}

//end main OMIT
