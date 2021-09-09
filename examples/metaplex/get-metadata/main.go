package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/36625090/solana-go/client"
	"github.com/36625090/solana-go/client/rpc"
	"github.com/36625090/solana-go/common"
	"github.com/36625090/solana-go/program/metaplex/tokenmeta"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// NFT in solana is a normal mint but only mint 1.
	// If you want to get its metadata, you need to know where it stored.
	// and you can use `tokenmeta.GetTokenMetaPubkey` to get the metadata account key
	// here I take a random Degenerate Ape Academy as an example
	mint := common.PublicKeyFromString("Sim5R3eZkkkLtpFAdo9m8tZg7uszXKacZAYwkFn371P")
	log.Println(mint.String())
	metadataAccount, err := tokenmeta.GetTokenMetaPubkey(mint)
	if err != nil {
		log.Fatalf("faield to get metadata account, err: %v", err)
	}

	// new a client
	c := client.NewClient(rpc.TestnetRPCEndpoint)

	res, err := c.GetSignaturesForAddress(context.Background(),
		mint.ToBase58(),
		rpc.GetConfirmedSignaturesForAddressConfig{
			Limit:      10,
			Commitment: rpc.CommitmentConfirmed,
		})
	for _, re := range res {
		if re.Err != nil {
			continue
		}
		log.Printf("slot: %d, block time: %d, signature: %s",
			re.Slot, *re.BlockTime, re.Signature,
		)
	}

	// get data which stored in metadataAccount
	accountInfo, err := c.GetAccountInfo(context.Background(), metadataAccount.ToBase58())
	if err != nil {
		log.Fatalf("failed to get accountInfo, err: %v", err)
	}
	// parse it
	metadata, err := tokenmeta.MetadataDeserialize(accountInfo.Data)
	if err != nil {
		//log.Fatalf("failed to parse metaAccount, err: %v", err)
	}
	log.Printf("%+v\n", metadata)
	txn := "3AxL629uV2YgAjnswd6Zbhm4MA9SgCFKC8pCPzS2GXACEgNrBYpCDuGvH67ks53CrWpFprpT74zLo3srmE5rrk7c"
	trans, err := c.GetTransaction(context.Background(), txn, rpc.GetTransactionWithLimitConfig{})
	bs, _ := json.Marshal(trans)
	log.Println(string(bs))
}
