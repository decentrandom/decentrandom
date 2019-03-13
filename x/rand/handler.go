package rand

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "rand" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {

		case MsgNewRound:
			return handleMsgNewRound(ctx, keeper, msg)

		case MsgAddTargets:
			return handleMsgAddTargets(ctx, keeper, msg)

		default:
			errMsg := fmt.Sprintf("Unrecognized rand Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgNewRound(ctx sdk.Context, keeper Keeper, msg MsgNewRound) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Id)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	keeper.SetRound(ctx, msg.Id, Round{Difficulty: msg.Difficulty, Owner: msg.Owner, Nonce: msg.Nonce, NonceHash: msg.NonceHash, Targets: msg.Targets, ScheduledTime: msg.ScheduledTime, SeedHeights: msg.SeedHeights})
	return sdk.Result{}
}

func handleMsgAddTargets(ctx sdk.Context, keeper Keeper, msg MsgAddTargets) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Id)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	// ****** important : It only sets, not adds
	keeper.SetTargets(ctx, msg.Id, msg.Targets)
	return sdk.Result{}
}
