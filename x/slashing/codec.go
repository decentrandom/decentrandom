package slashing

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/slashing"
)

// RegisterCodec -
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(slashing.MsgUnjail{}, "slashing/MsgUnjail", nil)
}

var cdcEmpty = codec.New()
