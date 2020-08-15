# x/rand

현재, 난수 관련 모듈은 LCD를 지원하지 않습니다.

## Keeper

- SetRound
- GetRound
- GetOwner
- GetTargets
- SetTargets
- SetNonce
- GetIDsIterator

## Msgs

- NewMsgNewRound
- NewMsgDeployNonce
- NewMsgAddTargets
- NewMsgRemoveTargets

## Handler

- handleMsgNewRound
- handleMsgDeployNonce
- handleMsgAddTargets
- handleMsgRemoveTargets
