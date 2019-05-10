package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/decentrandom/decentrandom/types/util"

	"github.com/tendermint/tendermint/libs/bech32"
)

var bech32Prefixes = []string{util.Bech32PrefixAccAddr, util.Bech32PrefixAccPub, util.Bech32PrefixValAddr, util.Bech32PrefixValPub, util.Bech32PrefixConsAddr, util.Bech32PrefixConsPub}

// main -
func main() {
	if len(os.Args) < 2 {
		fmt.Println("입력 값을 넣어주세요")
	}
	arg := os.Args[1]
	runFromBech32(arg)
	runFromHex(arg)
}

// runFromBech32 -
func runFromBech32(bech32str string) {
	hrp, bz, err := bech32.DecodeAndConvert(bech32str)
	if err != nil {
		fmt.Println("유효한 bech32 문자열이 아닙니다")
		return
	}
	fmt.Println("Bech32 분석:")
	fmt.Printf("Human readible part: %v\nBytes (hex): %X\n",
		hrp,
		bz,
	)
}

// runFromHex -
func runFromHex(hexaddr string) {
	bz, err := hex.DecodeString(hexaddr)
	if err != nil {
		fmt.Println("유효한 hex 문자열이 아닙니다")
		return
	}
	fmt.Println("Hex 분석:")
	fmt.Println("Bech32 formats:")
	for _, prefix := range bech32Prefixes {
		bech32Addr, err := bech32.ConvertAndEncode(prefix, bz)
		if err != nil {
			panic(err)
		}
		fmt.Println("  - " + bech32Addr)
	}
}
