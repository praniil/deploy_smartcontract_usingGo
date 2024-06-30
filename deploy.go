package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		panic(err)
	}
	chainId, err := client.NetworkID(context.Background())
	if err!= nil {
		panic(err)
	}
	fmt.Println(chainId);
}
