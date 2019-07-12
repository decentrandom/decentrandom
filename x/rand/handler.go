package rand

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler - 각각의 Msg에 필요한 handler를 리턴
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {

		case MsgNewRound:
			return handleMsgNewRound(ctx, keeper, msg)

		case MsgDeployNonce:
			return handleMsgDeployNonce(ctx, keeper, msg)

		case MsgAddTargets:
			return handleMsgAddTargets(ctx, keeper, msg)

		case MsgRemoveTargets:
			return handleMsgRemoveTargets(ctx, keeper, msg)

		case MsgDeploySeeds:
			return handleMsgDeploySeeds(ctx, keeper, msg)

		default:
			errMsg := fmt.Sprintf("unknown rand msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// handleMsgNewRound - 신규 라운드 생성을 위한 handler
func handleMsgNewRound(ctx sdk.Context, keeper Keeper, msg MsgNewRound) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.ID)) {
		return sdk.ErrUnauthorized("소유주 불일치").Result()
	}

	keeper.SetRound(ctx, msg.ID, Round{Difficulty: msg.Difficulty, Owner: msg.Owner, Nonce: msg.Nonce, NonceHash: msg.NonceHash, Targets: msg.Targets, ScheduledTime: msg.ScheduledTime})
	return sdk.Result{}
}

// handleMsgDepoloyNonce - Nonce 배포를 위한 handler
func handleMsgDeployNonce(ctx sdk.Context, keeper Keeper, msg MsgDeployNonce) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.ID)) {
		return sdk.ErrUnauthorized("소유주 불일치").Result()
	}

	keeper.SetNonce(ctx, msg.ID, msg.Nonce)
	return sdk.Result{}
}

// handleMsgAddTargets - 모집단 추가를 위한 handler
func handleMsgAddTargets(ctx sdk.Context, keeper Keeper, msg MsgAddTargets) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.ID)) {
		return sdk.ErrUnauthorized("소유주 불일치").Result()
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

// handleMsgRemoveTargets - 모집단 삭제를 위한 handler
func handleMsgRemoveTargets(ctx sdk.Context, keeper Keeper, msg MsgRemoveTargets) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.ID)) {
		return sdk.ErrUnauthorized("소유주 불일치").Result()
	}

	// 기존 Targets에서 일치하는 것 삭제
	updateTargets := keeper.GetTargets(ctx, msg.ID)

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

	keeper.SetTargets(ctx, msg.ID, updateTargets)
	return sdk.Result{}
}

// handleMsgDeploySeeds 0
func handleMsgDeploySeeds(ctx sdk.Context, keeper Keeper, msg MsgDeploySeeds) sdk.Result {

	// To-do

	return sdk.Result{}
}
