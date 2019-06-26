# DecentRandom

**주의 : 테스트넷 베타로 이동하기 위해 현재 사용하지 않습니다.**

P2P기반 탈중앙화 오픈 소스 난수 발생기

- [https://decentrandom.com](https://decentrandom.com)
- Cosmos SDK v0.35.0
- Tendermint v0.37.1
- Go 1.12 이상

## Warnings

- 현재 pre-alpha 버전입니다.
- 이 경고문이 사라지기 전에는 안정적인 버전이 아닙니다.
- 최적화나 캡슐화는 기능 구현이 어느 정도 된 이후 진행할 예정입니다.
- 의존성 관리가 dep에서 go module로 변경되었습니다.
- 현재 테스트넷 체인아이디는 mssp_0002로 전환 중 입니다.

## To-do

- 블록 헤더 수정

## Building

    make distclean
    make go-mod-cache
    make install

## 초기화 및 실행

    randd init --chain-id=mssp_0002 <Moniker>
    ulimit -n 4096
    randd start