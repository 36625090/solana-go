package main

import (
	"context"
	"log"

	"github.com/36625090/solana-go/client"
	"github.com/36625090/solana-go/client/rpc"
	"github.com/36625090/solana-go/common"
	"github.com/36625090/solana-go/program/tokenprog"
	"github.com/36625090/solana-go/types"
)

var alice = types.AccountFromPrivateKeyBytes([]byte{
	248, 218, 217, 179, 205, 246, 32, 71, 89, 196, 230, 186, 198, 3, 72, 129, 68, 123, 255, 168, 178, 159, 71, 77, 230, 224, 125, 128, 90, 71, 198, 151, 127, 110, 161, 46, 135, 199, 206, 180, 147, 196, 182, 171, 139, 194, 152, 37, 230, 55, 116, 178, 97, 9, 115, 255, 52, 86, 154, 215, 97, 168, 100, 213,
})
var mintPubkey = common.PublicKeyFromString("BNmuE7xMKtrfAYyb8tfXLcb5pAPYdaNEn6i1oMBddhX4")

var aliceTokenRandomTokenPubkey = common.PublicKeyFromString("8MZyXdPURRE5Tt5R8kcxemN87RuS6MUJ5BLmpn1V6CWZ")

//var aliceTokenATAPubkey = common.PublicKeyFromString("81Ck4pb8sZVYacLVHh4KbyiYHX8qnR4JvuZcyPiJN5cc")

func main() {
	c := client.NewClient(rpc.TestnetRPCEndpoint)
	//acct, _ := types.AccountFromBase58("9aScuM78feG8JXj3gCUJsXAkaaauUhkMpJyLftkj1XZW")
	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions: []types.Instruction{
			tokenprog.MintToChecked(
				mintPubkey,
				aliceTokenRandomTokenPubkey,
				alice.PublicKey,
				[]common.PublicKey{},
				1,
				0,
			),
		},

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
}
