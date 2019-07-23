package rand

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler -
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {

		case MsgNewRound:
			return handleMsgNewRound(ctx, keeper, msg)

		case MsgDeployNonce:
			return handleMsgDeployNonce(ctx, keeper, msg)

		case MsgAddTargets:
			return handleMsgAddTargets(ctx, keeper, msg)

		case MsgUpdateTargets:
			return handleMsgUpdateTargets(ctx, keeper, msg)

		default:
			errMsg := fmt.Sprintf("Unknown rand Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// handleMsgNewRound -
func handleMsgNewRound(ctx sdk.Context, keeper Keeper, msg MsgNewRound) sdk.Result {
	//if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.ID)) {
	//return sdk.ErrUnauthorized("Owner mismatch").Result()
	//}

	if msg.Owner.Empty() {
		return sdk.ErrUnauthorized("Owner is empty").Result()
	}

	keeper.SetRound(ctx, msg.ID, Round{Difficulty: msg.Difficulty, Owner: msg.Owner, Nonce: msg.Nonce, NonceHash: msg.NonceHash, Targets: msg.Targets, ScheduledTime: msg.ScheduledTime})
	return sdk.Result{}
}

// handleMsgDepoloyNonce -
func handleMsgDeployNonce(ctx sdk.Context, keeper Keeper, msg MsgDeployNonce) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.ID)) {
		return sdk.ErrUnauthorized("Owner mismatch").Result()
	}

	keeper.SetNonce(ctx, msg.ID, msg.Nonce)
	return sdk.Result{}
}

// handleMsgAddTargets -
func handleMsgAddTargets(ctx sdk.Context, keeper Keeper, msg MsgAddTargets) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.ID)) {
		return sdk.ErrUnauthorized("Owner mismatch").Result()
	}

	// 기존 Targets에 추가
	// 현재 Target에 중복된 값을 넣을 수는 없음
	// important ******
	// 만약 중복 응모가 가능하게 하려면 이를 어떻게 처리할 것인가? 이것이 필요한가는 의문
	newTargets := keeper.GetTargets(ctx, msg.ID)

	for _, b := range msg.Targets {
		exist := false
		for _, n := range newTargets {
			if n == b {
				exist = true
				break
			}
		}

		if !exist {
			newTargets = append(newTargets, b)
		}
	}

	keeper.SetTargets(ctx, msg.ID, newTargets)
	return sdk.Result{}
}

// handleMsgRemoveTargets -
func handleMsgUpdateTargets(ctx sdk.Context, keeper Keeper, msg MsgUpdateTargets) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.ID)) {
		return sdk.ErrUnauthorized("Owner mismatch").Result()
	}

	// 기존 Targets에서 일치하는 것 삭제
	//updateTargets := keeper.GetTargets(ctx, msg.ID)

	/*
		for i := 0; i < len(updateTargets); {
			exist := false
			for _, b := range msg.Targets {
				if b == updateTargets[i] {
					exist = true
					break
				}

				if !exist {
					i++
				} else {
					updateTargets = append(updateTargets[:i], updateTargets[i+1:]...)
				}
			}
		}
	*/

	keeper.SetTargets(ctx, msg.ID, msg.Targets)
	return sdk.Result{}
}
