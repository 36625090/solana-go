package tokenmeta

import (
	"github.com/36625090/solana-go/common"
)

func GetTokenMetaPubkey(mint common.PublicKey) (common.PublicKey, error) {
	metadataAccount, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
		},
		common.MetaplexTokenMetaProgramID,
	)
	if err != nil {
		return common.PublicKey{}, err
	}
	return metadataAccount, nil
}
