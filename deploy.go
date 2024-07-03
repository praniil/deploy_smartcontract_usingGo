package main

import (
	// "context"
	// "crypto/ecdsa"
	// "fmt"
	// "math/big"
	// "os"

	// "simple_storage/api"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	fmt.Println("we have a connection")

	blockNumber := big.NewInt(0)

	account := common.HexToAddress("0x57F013799D0762224852a016C2221F437482CDB5")
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	//pending balance:= for checking after submitting or waiting for a transaction to be confirmed
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)

	if err != nil {
		fmt.Println("failed to display the pending balance due to: ", err)
	}
	fmt.Println(pendingBalance)
	fmt.Println(balance)


}
