package main

import (
	"fmt"
	"log"

	"github.com/36625090/solana-go/types"
)

func main() {
	// create new account
	newAccount, _ := types.AccountFromBase58("EwxPevTC3Y2di5GRmmBk8kxYMBj4ecHUisML3iyNZC8h")
	fmt.Println(newAccount.PublicKey.ToBase58())
	fmt.Println(newAccount.PrivateKey)

	// recover account by its private key
	recoverAccount, err := types.AccountFromBytes(
		newAccount.PrivateKey,
	)
	if err != nil {
		log.Fatalf("failed to retrieve account from bytes, err: %v", err)
	}
	fmt.Println(recoverAccount.PublicKey.ToBase58())
}
