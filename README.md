# DecentRandom

P2P기반 탈중앙화 오픈 소스 난수 발생기

- [https://decentrandom.com](https://decentrandom.com)
- Cosmos SDK v0.37.0
- Tendermint v0.32.2
- Go 1.12 이상
- Ledger 지원 시 gcc 필요

## Warnings

- master branch는 업데이트를 위해 개발이 진행 중이므로 구동에 사용할 수 없습니다. 최신 릴리즈를 사용하시기 바랍니다.
- semidecentrandom은 decentrandom의 전단계로 검증인의 시드 생성 기능이 생략되어 있습니다.
- 검증인의 시드 대신 블록헤더를 시드로 사용합니다.
- 검증인의 시드 생성 기능은 Cosmos SDK의 IBC 구현이 완료된 이후 진행합니다.
- 현재 난수요청자의 RAND Deposit 관련 코드 구현 중입니다.

## 가이드 문서

- [https://decentrandom.com/docs/](https://decentrandom.com/docs/) 를 참고하시기 바랍니다.
