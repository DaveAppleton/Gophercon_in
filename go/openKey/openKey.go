package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"golang.org/x/crypto/ssh/terminal"
)

// getPassword gets a password without it being cached in stdin
func getPassword() (password string) {
	fmt.Print("Enter password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err == nil {
		password = strings.TrimSpace(string(bytePassword))
	}
	return
}

func getOwnerTx(keyStore string, pass string) (oTx *bind.TransactOpts, err error) {
	file, err := os.Open(keyStore)
	if err != nil {
		return
	}
	return bind.NewTransactor(file, pass)
}

func main() {
	var key string
	flag.StringVar(&key, "key", "", " path to key")
	flag.Parse()
	ownerTx, err := getOwnerTx(key, getPassword())
	// END OMIT
	if err != nil {
		log.Fatal("Error ", err)
	}
	fmt.Println(ownerTx.From.Hex())
}
