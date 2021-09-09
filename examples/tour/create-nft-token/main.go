package main

import (
	"context"
	"fmt"
	"github.com/36625090/solana-go/client"
	"github.com/36625090/solana-go/client/rpc"
	"github.com/36625090/solana-go/common"
	"github.com/36625090/solana-go/program/assotokenprog"
	"github.com/36625090/solana-go/program/metaplex/tokenmeta"
	"github.com/36625090/solana-go/program/sysprog"
	"github.com/36625090/solana-go/program/tokenprog"
	"github.com/36625090/solana-go/types"
	"log"
	"os"
	"time"
)

var alice = types.AccountFromPrivateKeyBytes([]byte{
	61, 103, 131, 192, 166, 221, 206, 161, 9, 35, 0, 68, 42, 71, 136, 199, 24, 39, 146, 179, 140, 139, 58, 149, 172, 52, 81, 3, 205, 236, 212, 77, 108,
	177, 196, 22, 17, 53, 254, 10, 102, 110, 46, 250, 91, 28, 21, 184, 202, 194, 206, 0, 15, 147, 229, 224, 198, 197, 133, 147, 200, 177, 40, 246,
})

var alicePubkey = common.PublicKeyFromString("8KJFQsdnPyzVYtazr6YFjXHDiWpHh151yudd7BJL1e7P")

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	c := client.NewClient(rpc.TestnetRPCEndpoint)

	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	log.Println("create non-fungible token, using block hash: ", res)
	balance, err := c.GetBalance(context.Background(), alicePubkey.ToBase58())
	if 0 == balance {
		res, err := c.RequestAirdrop(context.Background(), alicePubkey.ToBase58(), 1/0.000000001)
		if err != nil {
			log.Fatalf("request airdrop, err: %v\n", err)
		}
		log.Println("request airdrop, txn: ", res)
		balance = 1 / .000000001
	}
	if balance == 0 {
		fmt.Println("balance not enough: ", res.Blockhash)
		os.Exit(1)
	}
	rentExemptionBalance, err := c.GetMinimumBalanceForRentExemption(context.Background(), tokenprog.TokenAccountSize)
	if err != nil {
		log.Fatalf("get min balacne for rent exemption, err: %v", err)
	}
	mint := types.NewAccount()
	token := mint.PublicKey

	inst := tokenprog.InitializeMint(0, token, alice.PublicKey, common.PublicKey{})
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions: []types.Instruction{
			sysprog.CreateAccount(
				alice.PublicKey,
				mint.PublicKey,
				common.TokenProgramID,
				rentExemptionBalance,
				tokenprog.MintAccountSize,
			),
			inst,
		},
		Signers:         []types.Account{alice, mint},
		FeePayer:        alice.PublicKey,
		RecentBlockHash: res.Blockhash,
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	_, err = c.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		log.Fatalf("send raw tx error, err: %v\n", err)
	}

	log.Println("create token:", mint.PublicKey)
	ch := make(chan struct{})

	go createAccount(mint.PublicKey.String(), ch)

	select {
	case <-ch:
	}

	//// TODO
	//token = common.PublicKeyFromString(mint.PublicKey.String())
	//
	//inst, account := assotokenprog.CreateAssociatedTokenAccount(
	//	alice.PublicKey,
	//	alice.PublicKey,
	//	token,
	//)
	//log.Println("create account: ", account.String())
	//
	//rawTx, err = types.CreateRawTransaction(types.CreateRawTransactionParam{
	//	Instructions:    []types.Instruction{inst},
	//	Signers:         []types.Account{alice},
	//	FeePayer:        alice.PublicKey,
	//	RecentBlockHash: res.Blockhash,
	//})
	//if err != nil {
	//	log.Fatalf("generate tx error, err: %v\n", err)
	//}
	//
	//_, err = c.SendRawTransaction(context.Background(), rawTx)
	//if err != nil {
	//	log.Fatalf("send raw tx error, err: %v\n", err)
	//}
}

func createAccount(token string, ch chan struct{}) {
	time.Sleep(60 * time.Second)
	var mintPubkey = common.PublicKeyFromString(token)

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
	metadataAccount, err := tokenmeta.GetTokenMetaPubkey(mintPubkey)
	info, err := c.GetAccountInfo(context.Background(), metadataAccount.ToBase58())
	metadata, err := tokenmeta.MetadataDeserialize(info.Data)

	log.Println(info.Owner == common.TokenProgramID.String(),
		metadata, err)
	close(ch)
}
