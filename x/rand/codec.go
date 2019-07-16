package rand

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec -
func RegisterCodec(cdc *codec.Codec) {

	cdc.RegisterConcrete(MsgNewRound{}, "rand/NewRound", nil)
	cdc.RegisterConcrete(MsgDeployNonce{}, "rand/DeployNonce", nil)
	cdc.RegisterConcrete(MsgAddTargets{}, "rand/AddTargets", nil)
	cdc.RegisterConcrete(MsgRemoveTargets{}, "rand/UpdateTargets", nil)
}
