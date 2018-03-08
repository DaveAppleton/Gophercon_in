package main

import (
	"fmt"
	"log"
	"math/big"

	"wrapper"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var host = "https://radically-improved-albacore.quiknode.io/03f376cf-defe-49e9-a563-98859c7adf1e/6lXh94wSJ_N_wjKsCpkwiw==/"

func main() {
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatal(err)
	}
	// start OMIT
	contract := common.HexToAddress("0x7F3A19eDC1fFBd3fC2b4fc978b0d8f8C19316186")
	reader, err := wrapper.NewGopherCon(contract, client)
	if err != nil {
		log.Fatal(err)
	}
	numAddresses, _ := reader.NumTargets(nil)
	for i := int64(0); i < numAddresses.Int64(); i++ {
		addr, err := reader.HitList(nil, big.NewInt(i))
		if err != nil {
			log.Fatal(err)
		}
		hits, err := reader.Hits(nil, addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(addr.Hex(), hits)
	}
	// end OMIT
}
