package main

import (
	"fmt"
	"log"
	"math/big"
	"context"
	"crypto/ecdsa"
	
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
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

	account := common.HexToAddress("0x4e9001bb51F4cD89DdA8f154036CA1799471BcE9")
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

	//generating new wallet
	privateKey, err := crypto.HexToECDSA("48913ed55be67cb3283464797685ba3ac8624d10b737d3ddca7ecfd238154ab3")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	println("public key", publicKey)

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	fmt.Println("public key ecdsa", publicKeyECDSA)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)



	//how to produce public address manually, we take last 40 char of hash of public key and prefix it wit 0x.
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}
