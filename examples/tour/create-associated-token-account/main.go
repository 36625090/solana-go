package main

import (
	"context"
	"log"

	"github.com/36625090/solana-go/client"
	"github.com/36625090/solana-go/client/rpc"
	"github.com/36625090/solana-go/common"
	"github.com/36625090/solana-go/program/assotokenprog"
	"github.com/36625090/solana-go/types"
)

//var alice = types.AccountFromPrivateKeyBytes([]byte{196, 114, 86, 165, 59, 177, 63, 87, 43, 10, 176, 101, 225, 42, 129, 158, 167, 43, 81, 214, 254, 28, 196, 158, 159, 64, 55, 123, 48, 211, 78, 166, 127, 96, 107, 250, 152, 133, 208, 224, 73, 251, 113, 151, 128, 139, 86, 80, 101, 70, 138, 50, 141, 153, 218, 110, 56, 39, 122, 181, 120, 55, 86, 185})
var alice = types.AccountFromPrivateKeyBytes([]byte{
	61, 103, 131, 192, 166, 221, 206, 161, 9, 35, 0, 68, 42, 71, 136, 199, 24, 39, 146, 179, 140, 139, 58, 149, 172, 52, 81, 3, 205, 236, 212, 77, 108,
	177, 196, 22, 17, 53, 254, 10, 102, 110, 46, 250, 91, 28, 21, 184, 202, 194, 206, 0, 15, 147, 229, 224, 198, 197, 133, 147, 200, 177, 40, 246,
})

var mintPubkey = common.PublicKeyFromString("CW28sov3Dseo2NQSvJJd1yQnMxxmMiPJ2jzBetnipigZ")

func main() {
	c := client.NewClient(rpc.TestnetRPCEndpoint)

	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	ins, account := assotokenprog.CreateAssociatedTokenAccount(
		alice.PublicKey,
		alice.PublicKey,
		mintPubkey,
	)

	log.Println("create token account:", account.String(),
		" token: ", mintPubkey.ToBase58(),
		" hash ", res.Blockhash)
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions:    []types.Instruction{ins},
		Signers:         []types.Account{alice},
		FeePayer:        alice.PublicKey,
		RecentBlockHash: res.Blockhash,
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	txhash, err := c.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		log.Fatalf("send raw tx error, err: %v\n", err)
	}

	log.Println("txhash:", txhash)
	resp, err := c.GetTransaction(context.Background(), txhash, rpc.GetTransactionWithLimitConfig{})
	log.Println(resp)
}
