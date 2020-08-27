# x/rand

현재, rand 모듈은 LCD를 지원하지 않습니다. (추후 지원 예정)

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
