package main

import (
	"context"
	"fmt"
	"log"

	"github.com/36625090/solana-go/client"
	"github.com/36625090/solana-go/client/rpc"
)

func main() {
	c := client.NewClient(rpc.MainnetRPCEndpoint)

	resp, err := c.GetVersion(context.TODO())
	if err != nil {
		log.Fatalf("failed to version info, err: %v", err)
	}

	fmt.Println("version", resp.SolanaCore)
}
