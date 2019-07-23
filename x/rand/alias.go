package rand

import (
	"github.com/decentrandom/decentrandom/x/rand/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgNewRound      = types.NewMsgNewRound
	NewMsgDeployNonce   = types.NewMsgDeployNonce
	NewMsgAddTargets    = types.NewMsgAddTargets
	NewMsgUpdateTargets = types.NewMsgUpdateTargets
	ModuleCdc           = types.ModuleCdc
	RegisterCodec       = types.RegisterCodec
)

type (
	MsgNewRound      = types.MsgNewRound
	MsgDeployNonce   = types.MsgDeployNonce
	MsgAddTargets    = types.MsgAddTargets
	MsgUpdateTargets = types.MsgUpdateTargets
	Round            = types.Round
	QueryResRoundIDs = types.QueryResRoundIDs
)
