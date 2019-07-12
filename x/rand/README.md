# x/rand

난수 관련 모듈

## Keeper

- SetRound
- GetRound
- GetOwner
- GetTargets
- SetTargets
- SetNonce
- GetIDsIterator
- SetSeeds

## Msgs

- NewMsgNewRound : 신규 라운드 생성
- NewMsgDeployNonce : 논스 공표
- NewMsgAddTargets : 모집단 추가
- NewMsgRemoveTargets : 모집단 삭제
- NewMsgDeploySeeds : 시드 공표

## Handler

- handleMsgNewRound
- handleMsgDeployNonce
- handleMsgAddTargets
- handleMsgRemoveTargets
- handleMsgDeploySeeds