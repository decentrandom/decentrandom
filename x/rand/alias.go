package rand

import (
	"github.com/decentrandom/decentrandom/x/rand/types"
)

const (
	// ModuleName -
	ModuleName = types.ModuleName

	// RouterKey -
	RouterKey = types.RouterKey

	// StoreKey -
	StoreKey = types.StoreKey
)

var (
	// NewMsgNewRound -
	NewMsgNewRound = types.NewMsgNewRound

	// NewMsgDeployNonce -
	NewMsgDeployNonce = types.NewMsgDeployNonce

	// NewMsgAddTargets -
	NewMsgAddTargets = types.NewMsgAddTargets

	// NewMsgRemoveTargets -
	NewMsgRemoveTargets = types.NewMsgRemoveTargets

	// NewMsgDeploySeeds -
	NewMsgDeploySeeds = types.NewMsgDeploySeeds

	// ModuleCdc -
	ModuleCdc = types.ModuleCdc

	// RegisterCodec -
	RegisterCodec = types.RegisterCodec
)

type (
	// MsgNewRound -
	MsgNewRound = types.MsgNewRound

	// MsgDeployNonce -
	MsgDeployNonce = types.MsgDeployNonce

	// MsgAddTargets -
	MsgAddTargets = types.MsgAddTargets

	// MsgRemoveTargets -
	MsgRemoveTargets = types.MsgRemoveTargets

	// MsgDeploySeeds -
	MsgDeploySeeds = types.MsgDeploySeeds

	// QueryResResolve -
	QueryResResolve = types.QueryResResolve

	// QueryResNames -
	QueryResNames = types.QueryResNames
	Whois         = types.Whois
)
