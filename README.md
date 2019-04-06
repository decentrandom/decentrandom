# DecentRandom
P2P기반 탈중앙화 오픈 소스 난수 발생기

- https://decentrandom.com

# Warnings!!

- 현재 pre-alpha 버전입니다.
- 이 경고문이 사라지기 전에는 안정적인 버전이 아닙니다.
- Go언어에 익숙하지 않은 상태에서 commit 중입니다.
- 최적화나 캡슐화는 기능 구현이 어느 정도 된 이후 진행할 예정입니다.

# To-do

- 블록 헤더 수정
- 라운드 seed 정보 처리 방식 확정

# Building

- golang 1.12 이상을 필요로 합니다.

<pre><code>make get_tools
dep ensure -v
dep ensure -update -v
make install
</code></pre>