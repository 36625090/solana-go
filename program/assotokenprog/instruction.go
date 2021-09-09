package assotokenprog

import (
	"github.com/36625090/solana-go/common"
	"github.com/36625090/solana-go/types"
)

// CreateAssociatedTokenAccount is the only instruction in associated token program
func CreateAssociatedTokenAccount(funder, wallet, tokenMint common.PublicKey) (types.Instruction, common.PublicKey) {
	assosiatedAccount, _, _ := common.FindAssociatedTokenAddress(wallet, tokenMint)
	return types.Instruction{
		ProgramID: common.SPLAssociatedTokenAccountProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: funder, IsSigner: true, IsWritable: true},
			{PubKey: assosiatedAccount, IsSigner: false, IsWritable: true},
			{PubKey: wallet, IsSigner: false, IsWritable: false},
			{PubKey: tokenMint, IsSigner: false, IsWritable: false},
			{PubKey: common.SystemProgramID, IsSigner: false, IsWritable: false},
			{PubKey: common.TokenProgramID, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: []byte{},
	}, assosiatedAccount
}
