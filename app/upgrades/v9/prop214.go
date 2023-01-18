package v9

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	gammkeeper "github.com/tessornetwork/mokita/x/gamm/keeper"
	"github.com/tessornetwork/mokita/x/gamm/pool-models/balancer"
)

// Executes prop214, https://www.mintscan.io/mokita/proposals/214
// Run `mokitad q gov proposal 214` to see the text.
// It was voted in, and it has update instructions:
// Voting YES for this proposal would reduce the Pool 1 (MOKI/ATOM) swap fee from 0.3% to 0.2%
func ExecuteProp214(ctx sdk.Context, gamm *gammkeeper.Keeper) {
	poolId := 1
	pool, err := gamm.GetPoolAndPoke(ctx, uint64(poolId))
	if err != nil {
		panic(err)
	}

	balancerPool, ok := pool.(*balancer.Pool)
	if !ok {
		panic(ok)
	}

	balancerPool.PoolParams.SwapFee = sdk.MustNewDecFromStr("0.002")

	// Kept as comments for recordkeeping. SetPool is now private:
	// 		err = gamm.SetPool(ctx, balancerPool)
	// 		if err != nil {
	//	 		panic(err)
	//  	}
}
