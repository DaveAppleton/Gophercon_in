package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/DaveAppleton/ether_go/ethKeys"

	"keys"
	"wrapper"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// startClient OMIT
var baseClient *backends.SimulatedBackend
var banker *ethKeys.AccountKey

func getClient() (client *backends.SimulatedBackend, err error) {
	if baseClient != nil {
		return baseClient, nil
	}
	funds, _ := new(big.Int).SetString("1000000000000000000000000000", 10)
	banker = keys.RoleKey("banker")
	baseClient = backends.NewSimulatedBackend(core.GenesisAlloc{
		banker.PublicKey(): {Balance: funds},
	})
	return baseClient, nil
}

// endClient OMIT

// startMine OMIT
func mine(what string, tx *types.Transaction, err error, client *backends.SimulatedBackend) {
	if err != nil {
		log.Fatal(err)
	}
	client.Commit()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	receipt, err := client.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(what, " Status : ", receipt.Status, "GasUsed ", receipt.GasUsed)
}

// endMine OMIT

func main() {
	client, err := getClient()
	if err != nil {
		log.Fatal(err)
	}
	// startCode OMIT
	gConAddress, tx, gCon, err := wrapper.DeployGopherCon(keys.KeyTx(banker), client)
	if err != nil {
		log.Fatal(err)
	}
	mine("deploy", tx, err, client)
	fmt.Println("Contract deployed at : ", gConAddress.Hex(), " in tx ", tx.Hash().Hex())
	//
	count, _ := gCon.NumTargets(nil)
	fmt.Println(count, " Target")
	tx, err = gCon.HitMe(keys.KeyTx(banker))
	mine("HitMe", tx, err, client)
	//
	count, _ = gCon.NumTargets(nil)
	fmt.Println(count, " Target")
	// endCode OMIT
}
