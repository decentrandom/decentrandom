package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc -
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec - Amino를 위한 concrete codec 등록
func RegisterCodec(cdc *codec.Codec) {

	cdc.RegisterConcrete(MsgNewRound{}, "rand/NewRound", nil)
	cdc.RegisterConcrete(MsgDeployNonce{}, "rand/DeployNonce", nil)
	cdc.RegisterConcrete(MsgAddTargets{}, "rand/AddTargets", nil)
	cdc.RegisterConcrete(MsgRemoveTargets{}, "rand/RemoveTargets", nil)
}
