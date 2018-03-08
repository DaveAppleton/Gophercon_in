package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
)

// abigen -sol gophercon.sol -pkg main -out wrapper.go
var host = "https://radically-improved-albacore.quiknode.io/03f376cf-defe-49e9-a563-98859c7adf1e/6lXh94wSJ_N_wjKsCpkwiw==/"

func makeFilter(signature string, address string, start int64) ethereum.FilterQuery {
	startInt := big.NewInt(start)
	contractAddress := common.HexToAddress(address)
	topic := crypto.Keccak256Hash([]byte(signature))
	thumpFilter := ethereum.FilterQuery{}
	thumpFilter.Addresses = []common.Address{contractAddress}
	thumpFilter.Topics = [][]common.Hash{[]common.Hash{topic}}
	thumpFilter.FromBlock = startInt
	return thumpFilter
}

func main() {
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	thumpFilter := makeFilter("Thump(address,uint256)", "0x7F3A19eDC1fFBd3fC2b4fc978b0d8f8C19316186", 1875695)
	logs, err := client.FilterLogs(ctx, thumpFilter)
	if err != nil {
		log.Fatal(err)
	}
	for _, log := range logs {
		address := common.BytesToAddress(log.Data[12:32])
		count, _ := strconv.ParseInt(common.Bytes2Hex(log.Data[32:]), 16, 64)
		fmt.Println(log.BlockNumber, " : Thump(", address.Hex(), ",", count, ")")
	}
}
