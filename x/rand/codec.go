package rand

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {

	cdc.RegisterConcrete(MsgNewRound{}, "rand/NewRound", nil)
	cdc.RegisterConcrete(MsgAddTargets{}, "rand/AddTargets", nil)
}
