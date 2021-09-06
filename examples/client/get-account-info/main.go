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

	// get account info
	accountInfo, err := c.GetAccountInfo(
		context.TODO(),
		"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
	)
	if err != nil {
		log.Fatalf("failed to get balance, err: %v", err)
	}
	fmt.Printf("accountInfo: %v\n", accountInfo)
}
