package tokenprog

import (
	"context"
	"github.com/36625090/solana-go/client"
	"github.com/36625090/solana-go/client/rpc"
	"log"
	"reflect"
	"testing"

	"github.com/36625090/solana-go/common"
	"github.com/36625090/solana-go/types"
)

func TestMintTo(t *testing.T) {
	type args struct {
		mintPubkey    common.PublicKey
		destPubkey    common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
		amount        uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				mintPubkey:    common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				destPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys: []common.PublicKey{},
				amount:        99999,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{7, 159, 134, 1, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MintTo(tt.args.mintPubkey, tt.args.destPubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MintTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMintToChecked(t *testing.T) {
	type args struct {
		mintPubkey    common.PublicKey
		destPubkey    common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
		decimals      uint8
		amount        uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				mintPubkey:    common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				destPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys: []common.PublicKey{},
				decimals:      5,
				amount:        99999,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{14, 159, 134, 1, 0, 0, 0, 0, 0, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MintToChecked(tt.args.mintPubkey, tt.args.destPubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.amount, tt.args.decimals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MintToChecked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	type args struct {
		srcPubkey     common.PublicKey
		destPubkey    common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
		amount        uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				srcPubkey:     common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				destPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys: []common.PublicKey{},
				amount:        99999,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{3, 159, 134, 1, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Transfer(tt.args.srcPubkey, tt.args.destPubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransferChecked(t *testing.T) {
	type args struct {
		srcPubkey     common.PublicKey
		destPubkey    common.PublicKey
		mintPubkey    common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
		amount        uint64
		decimals      uint8
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				srcPubkey:     common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				destPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				mintPubkey:    common.PublicKeyFromString("HFCNHUwPxRqqW6gaLd3uUjJcEUfjnRptJzh4xvnNmavv"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys: []common.PublicKey{},
				amount:        99999,
				decimals:      4,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("HFCNHUwPxRqqW6gaLd3uUjJcEUfjnRptJzh4xvnNmavv"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{12, 159, 134, 1, 0, 0, 0, 0, 0, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransferChecked(tt.args.srcPubkey, tt.args.destPubkey, tt.args.mintPubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.amount, tt.args.decimals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransferChecked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBurn(t *testing.T) {
	type args struct {
		accountPubkey common.PublicKey
		mintPubkey    common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
		amount        uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				accountPubkey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				mintPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys: []common.PublicKey{},
				amount:        99999,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{8, 159, 134, 1, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Burn(tt.args.accountPubkey, tt.args.mintPubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Burn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBurnChecked(t *testing.T) {
	type args struct {
		accountPubkey common.PublicKey
		mintPubkey    common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
		amount        uint64
		decimals      uint8
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				accountPubkey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				mintPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys: []common.PublicKey{},
				amount:        99999,
				decimals:      9,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{15, 159, 134, 1, 0, 0, 0, 0, 0, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BurnChecked(tt.args.accountPubkey, tt.args.mintPubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.amount, tt.args.decimals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BurnChecked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloseAccount(t *testing.T) {
	type args struct {
		accountPubkey common.PublicKey
		destPubkey    common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				accountPubkey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				destPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys: []common.PublicKey{},
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CloseAccount(tt.args.accountPubkey, tt.args.destPubkey, tt.args.authPubkey, tt.args.signerPubkeys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CloseAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeAccount2(t *testing.T) {
	type args struct {
		accountPubkey common.PublicKey
		mintPubkey    common.PublicKey
		ownerPubkey   common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				accountPubkey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				mintPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				ownerPubkey:   common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
				},
				Data: []byte{16, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeAccount2(tt.args.accountPubkey, tt.args.mintPubkey, tt.args.ownerPubkey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeAccount2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreezeAccount(t *testing.T) {
	type args struct {
		accountPubkey common.PublicKey
		mintPubkey    common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				accountPubkey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				mintPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FreezeAccount(tt.args.accountPubkey, tt.args.mintPubkey, tt.args.authPubkey, tt.args.signerPubkeys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FreezeAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThawAccount(t *testing.T) {
	type args struct {
		accountPubkey common.PublicKey
		mintPubkey    common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				accountPubkey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				mintPubkey:    common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:    common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThawAccount(tt.args.accountPubkey, tt.args.mintPubkey, tt.args.authPubkey, tt.args.signerPubkeys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThawAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApprove(t *testing.T) {
	type args struct {
		sourcePubkey   common.PublicKey
		delegatePubkey common.PublicKey
		authPubkey     common.PublicKey
		signerPubkeys  []common.PublicKey
		amount         uint64
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				sourcePubkey:   common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				delegatePubkey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				authPubkey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys:  []common.PublicKey{},
				amount:         99999,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{4, 159, 134, 1, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Approve(tt.args.sourcePubkey, tt.args.delegatePubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Approve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevoke(t *testing.T) {
	type args struct {
		srcPubkey     common.PublicKey
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				srcPubkey:  common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				authPubkey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Revoke(tt.args.srcPubkey, tt.args.authPubkey, tt.args.signerPubkeys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Revoke() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApproveChecked(t *testing.T) {
	type args struct {
		sourcePubkey   common.PublicKey
		mintPubkey     common.PublicKey
		delegatePubkey common.PublicKey
		authPubkey     common.PublicKey
		signerPubkeys  []common.PublicKey
		amount         uint64
		decimals       uint8
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				sourcePubkey:   common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				mintPubkey:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
				delegatePubkey: common.PublicKeyFromString("DuNVVSmxNkXZvzT7fEDAWhfDvEgBYohuCGYB9AQzrctY"),
				authPubkey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				signerPubkeys:  []common.PublicKey{},
				amount:         99999,
				decimals:       9,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("DuNVVSmxNkXZvzT7fEDAWhfDvEgBYohuCGYB9AQzrctY"), IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{13, 159, 134, 1, 0, 0, 0, 0, 0, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApproveChecked(tt.args.sourcePubkey, tt.args.mintPubkey, tt.args.delegatePubkey, tt.args.authPubkey, tt.args.signerPubkeys, tt.args.amount, tt.args.decimals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApproveChecked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeMultisig(t *testing.T) {
	type args struct {
		authPubkey    common.PublicKey
		signerPubkeys []common.PublicKey
		miniRequired  uint8
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				authPubkey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
				signerPubkeys: []common.PublicKey{
					common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					common.PublicKeyFromString("DuNVVSmxNkXZvzT7fEDAWhfDvEgBYohuCGYB9AQzrctY"),
					common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				},
				miniRequired: 2,
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{

					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
					{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("DuNVVSmxNkXZvzT7fEDAWhfDvEgBYohuCGYB9AQzrctY"), IsSigner: true, IsWritable: false},
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: false},
				},
				Data: []byte{2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeMultiSig(tt.args.authPubkey, tt.args.signerPubkeys, tt.args.miniRequired); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeMultiSig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSyncNative(t *testing.T) {
	type args struct {
		accountPubkey common.PublicKey
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				accountPubkey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"),
			},
			want: types.Instruction{
				ProgramID: common.TokenProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("FtvD2ymcAFh59DGGmJkANyJzEpLDR1GLgqDrUxfe2dPm"), IsSigner: false, IsWritable: true},
				},
				Data: []byte{17},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SyncNative(tt.args.accountPubkey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SyncNative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisableMint(t *testing.T) {

	var alice = types.AccountFromPrivateKeyBytes([]byte{
		61, 103, 131, 192, 166, 221, 206, 161, 9, 35, 0, 68, 42, 71, 136, 199, 24, 39, 146, 179, 140, 139, 58, 149, 172, 52, 81, 3, 205, 236, 212, 77, 108,
		177, 196, 22, 17, 53, 254, 10, 102, 110, 46, 250, 91, 28, 21, 184, 202, 194, 206, 0, 15, 147, 229, 224, 198, 197, 133, 147, 200, 177, 40, 246,
	})

	var token = common.PublicKeyFromString("5HwM7QxqjGKyNMFcNNv7tVFWu67itbVykjpZNnJoADjC")

	inst := DisableMint(token, alice.PublicKey, []common.PublicKey{alice.PublicKey})
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	c := client.NewClient(rpc.TestnetRPCEndpoint)

	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions: []types.Instruction{
			inst,
		},
		Signers:         []types.Account{alice},
		FeePayer:        alice.PublicKey,
		RecentBlockHash: res.Blockhash,
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	txn, err := c.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		log.Fatalf("send raw tx error, err: %v\n", err)
	}

	log.Println("disable mint:", txn)
}
