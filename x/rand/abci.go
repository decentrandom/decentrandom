package rand

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker -
func EndBlocker(ctx sdk.Context, k Keeper) {

	fmt.Println()
	fmt.Printf("----------------")
	fmt.Println()
	fmt.Printf("RAND Deposit")
	fmt.Println()

	// TODO : process RAND deposit
	// ex : sdk.DecCoins{{urand, sdk.NewDec()}}
	// ex : k.AllocateTokensToAccount(ctx, 주소, tokens)
	//
	// supply keeper의 SendCoinsFromModuleToAccount, SendCoinsFromAccountToModule 사용

}
