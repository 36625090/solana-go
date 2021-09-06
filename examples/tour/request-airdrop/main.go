package main

import (
	"context"
	"fmt"
	"log"

	"github.com/36625090/solana-go/client"
	"github.com/36625090/solana-go/client/rpc"
	"github.com/36625090/solana-go/types"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	newAccount := types.NewAccount()
	fmt.Println(newAccount.PublicKey.ToBase58())

	txhash, err := c.RequestAirdrop(
		context.Background(),
		newAccount.PublicKey.ToBase58(),
		1e9, // 1 SOL = 10^9 lamports
	)
	if err != nil {
		log.Fatalf("failed to request airdrop, err: %v", err)
	}

	fmt.Println("txhash:", txhash)
}
