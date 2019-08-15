# DecentRandom

P2P기반 탈중앙화 오픈 소스 난수 발생기

- [https://decentrandom.com](https://decentrandom.com)
- Cosmos SDK v0.36.0-rc1
- Tendermint v0.32.1
- Go 1.12 이상

## Warnings

- semidecentrandom은 decentrandom의 전단계로 검증인의 시드 생성 기능이 생략되어 있습니다.
- 검증인의 시드 대신 블록헤더를 시드로 사용합니다.
- 검증인의 시드 생성 기능은 Cosmos SDK의 IBC 구현이 완료된 이후 진행합니다.

## Building

    go mod tidy
    make build
