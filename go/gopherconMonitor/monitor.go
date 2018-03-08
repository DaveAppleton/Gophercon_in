package main

import (
	"fmt"
	"log"

	"wrapper"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var host = "/Users/daveappleton/Library/Ethereum/rinkeby/geth.ipc"

//

func main() {
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatal(err)
	}
	// start OMIT
	contract := common.HexToAddress("0x7F3A19eDC1fFBd3fC2b4fc978b0d8f8C19316186")
	gopherconFilterer, err := wrapper.NewGopherConFilterer(contract, client)
	if err != nil {
		log.Fatal(err)
	}
	start := uint64(1875695)
	b := bind.WatchOpts{Start: &start}
	thumpChannel := make(chan *wrapper.GopherConThump, 50)
	_, err = gopherconFilterer.WatchThump(&b, thumpChannel)
	if err != nil {
		log.Fatal(err)
	}
	for {
		x := <-thumpChannel
		fmt.Println("Thump(", x.Who.Hex(), ",", x.Count, ")")
	}
	// end OMIT
}
