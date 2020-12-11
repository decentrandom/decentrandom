# x/rand

현재, rand 모듈은 LCD를 지원하지 않습니다. (추후 지원 예정)

## Keeper

- To-do : Rand Seed, Nonce 전체 관리를 위한 내용
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
