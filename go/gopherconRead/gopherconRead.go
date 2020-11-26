package main

import (
	"fmt"
	"log"

	"com.xxx/wrapper"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var host = "https://radically-improved-albacore.quiknode.io/03f376cf-defe-49e9-a563-98859c7adf1e/6lXh94wSJ_N_wjKsCpkwiw==/"

func main() {
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatal(err)
	}
	// long OMIT
	contract := common.HexToAddress("0x7F3A19eDC1fFBd3fC2b4fc978b0d8f8C19316186")
	gopherconFilterer, err := wrapper.NewGopherConFilterer(contract, client)
	if err != nil {
		log.Fatal(err)
	}
	b := bind.FilterOpts{Start: 1875695}
	thumpIterator, err := gopherconFilterer.FilterThump(&b)
	if err != nil {
		log.Fatal(err)
	}
	// start OMIT
	for thumpIterator.Next() {
		address := thumpIterator.Event.Who
		count := thumpIterator.Event.Count
		fmt.Println("Thump(", address.Hex(), ",", count, ")")
	}
	// end OMIT
}
