package stableswap

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"

	types "github.com/tessornetwork/mokita/x/gamm/types"
)

// RegisterLegacyAminoCodec registers the necessary x/gamm interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&Pool{}, "mokita/gamm/StableswapPool", nil)
	cdc.RegisterConcrete(&MsgCreateStableswapPool{}, "mokita/gamm/create-stableswap-pool", nil)
	cdc.RegisterConcrete(&MsgStableSwapAdjustScalingFactors{}, "mokita/gamm/stableswap-adjust-scaling-factors", nil)
	cdc.RegisterConcrete(&PoolParams{}, "mokita/gamm/StableswapPoolParams", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"mokita.gamm.v1beta1.PoolI",
		(*types.PoolI)(nil),
		&Pool{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateStableswapPool{},
		&MsgStableSwapAdjustScalingFactors{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/bank module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	sdk.RegisterLegacyAminoCodec(amino)
	RegisterLegacyAminoCodec(authzcodec.Amino)
	amino.Seal()
}

const PoolTypeName string = "Stableswap"
