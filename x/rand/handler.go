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

		default:
			errMsg := fmt.Sprintf("Unrecognized rand Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgNewRound(ctx sdk.Context, keeper Keeper, msg MsgNewRound) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(Ctx, msg.Id)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	keeper.SetRound(ctx, msg.Id, msg)
	return sdk.Result{}
}
