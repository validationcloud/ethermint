package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"

	cryptocodec "github.com/validationcloud/validator-service/cmd/cosmos_staker/encoding/ethermint/crypto/codec"

	"github.com/validationcloud/validator-service/cmd/cosmos_staker/encoding/ethermint/types"
)

// RegisterLegacyAminoCodec registers Interfaces from types, crypto, and SDK std.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
}

// RegisterInterfaces register the Ethermint key concrete types.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	//registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})
	//registry.RegisterImplementations((*cryptotypes.PrivKey)(nil), &ethsecp256k1.PrivKey{})
	std.RegisterInterfaces(registry)
	cryptocodec.RegisterInterfaces(registry)
	types.RegisterInterfaces(registry)
}
