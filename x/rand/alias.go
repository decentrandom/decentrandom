package rand

import (
	"github.com/decentrandom/decentrandom/x/rand/internal/keeper"
	"github.com/decentrandom/decentrandom/x/rand/internal/types"
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

	// NewKeeper -
	NewKeeper = keeper.NewKeeper

	// NewQuerier -
	NewQuerier = keeper.NewQuerier

	// NewMsgNewRound -
	NewMsgNewRound = types.NewMsgNewRound

	// NewMsgDeployNonce -
	NewMsgDeployNonce = types.NewMsgDeployNonce

	// NewMsgAddTargets -
	NewMsgAddTargets = types.NewMsgAddTargets

	// NewMsgUpdateTargets -
	NewMsgUpdateTargets = types.NewMsgUpdateTargets

	// ModuleCdc -
	ModuleCdc = types.ModuleCdc

	// RegisterCodec -
	RegisterCodec = types.RegisterCodec
)

type (

	// Keeper -
	Keeper = keeper.Keeper

	// MsgNewRound -
	MsgNewRound = types.MsgNewRound

	// MsgDeployNonce -
	MsgDeployNonce = types.MsgDeployNonce

	// MsgAddTargets -
	MsgAddTargets = types.MsgAddTargets

	// MsgUpdateTargets -
	MsgUpdateTargets = types.MsgUpdateTargets

	// Seed -
	//Seed = types.Seed

	// Round -
	Round = types.Round

	// Rounds -
	Rounds = types.Rounds

	// QueryResRoundIDs -
	QueryResRoundIDs = types.QueryResRoundIDs

	// Nonce -
	Nonce = types.Nonce
)
