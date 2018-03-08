package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func getOwnerTx(keyStore string, pass string) (oTx *bind.TransactOpts, err error) {
	file, err := os.Open(keyStore)
	if err != nil {
		return
	}
	return bind.NewTransactor(file, pass)
}

func main() {
	keyStore := "UTC--2018-03-03T18-22-00.825Z--794d7f1192e703b934cc92746c55a3d159f781fe"
	ownerTx, err := getOwnerTx(keyStore, "gophercon")
	if err != nil {
		log.Fatal("Error ", err)
	}
	fmt.Println(ownerTx.From.Hex())
}
