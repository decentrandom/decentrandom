package rand

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker -
func EndBlocker(ctx sdk.Context, k Keeper) {

	fmt.Println()
	fmt.Printf("RAND Deposit")
	fmt.Println()

	// TODO : process RAND depost
}
