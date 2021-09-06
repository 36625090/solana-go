package main

import (
	"context"
	"fmt"
	"log"

	"github.com/36625090/solana-go/client"
	"github.com/36625090/solana-go/client/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	balance, err := c.GetBalance(
		context.Background(),
		"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
	)
	if err != nil {
		log.Fatalln("get balance error", err)
	}
	fmt.Println(balance)
}
