# DecentRandom

P2P기반 탈중앙화 오픈 소스 난수 발생기

- [https://decentrandom.com](https://decentrandom.com)
- Cosmos SDK v0.33.0 기준
- Tendermint v0.31.0-dev0-fix0
- Go 1.12 이상

## Warnings

- 현재 pre-alpha 버전입니다.
- 이 경고문이 사라지기 전에는 안정적인 버전이 아닙니다.
- Go언어에 익숙하지 않은 상태에서 commit 중입니다.
- 최적화나 캡슐화는 기능 구현이 어느 정도 된 이후 진행할 예정입니다.
- 의존성 관리가 dep에서 go module로 변경되었습니다.
- 현재 테스트넷 체인아이디는 mssp_0001입니다.

## To-do

- 블록 헤더 수정

## Building

    make install

## 초기화 및 실행

    randd init --chain-id=mssp_0001 mssp_0001
    randcli keys add validator
    randd add-genesis-account $(randcli keys show validator -a) 1000000000000mrand
    randd gentx --name validator
    randd collect-gentxs
    ulimit -n 4096
    randd start