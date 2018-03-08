package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"wrapper"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var host = "https://radically-improved-albacore.quiknode.io/03f376cf-defe-49e9-a563-98859c7adf1e/6lXh94wSJ_N_wjKsCpkwiw==/"

func getOwnerTx(keyStore string, pass string) (oTx *bind.TransactOpts, err error) {
	file, err := os.Open(keyStore)
	if err != nil {
		return
	}
	return bind.NewTransactor(file, pass)
}

func main() {
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatal(err)
	}
	keyStore := "UTC--2018-03-03T18-22-00.825Z--794d7f1192e703b934cc92746c55a3d159f781fe"
	// start OMIT
	ownerTx, err := getOwnerTx(keyStore, "gophercon")
	contractAddress := common.HexToAddress("0x7F3A19eDC1fFBd3fC2b4fc978b0d8f8C19316186")
	gopherCon, err := wrapper.NewGopherCon(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := gopherCon.HitMe(ownerTx)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	rcpt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatal(err)
	}
	if rcpt.Status == 1 {
		// end OMIT
		address := common.BytesToAddress(rcpt.Logs[0].Data[12:32])
		count, _ := strconv.ParseInt(common.Bytes2Hex(rcpt.Logs[0].Data[32:]), 16, 64)
		fmt.Println("Transaction sent to ", contractAddress.Hex())
		fmt.Println("Status ", rcpt.Status, " (success)")
		fmt.Println(rcpt.Logs[0].BlockNumber, " : Thump(", address.Hex(), ",", count, ") detected")
	}

}
