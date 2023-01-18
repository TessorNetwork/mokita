package wasmbinding

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tessornetwork/mokita/app"
	"github.com/tessornetwork/mokita/wasmbinding"
	"github.com/tessornetwork/mokita/wasmbinding/bindings"
	"github.com/tessornetwork/mokita/x/gamm/pool-models/balancer"
)

// we must pay this many umoki for every pool we create
var poolFee int64 = 1000000000

var defaultFunds = sdk.NewCoins(
	sdk.NewInt64Coin("uatom", 333000000),
	sdk.NewInt64Coin("umoki", 555000000+2*poolFee),
	sdk.NewInt64Coin("ustar", 999000000),
)

func SetupCustomApp(t *testing.T, addr sdk.AccAddress) (*app.MokitaApp, sdk.Context) {
	mokita, ctx := CreateTestInput()
	wasmKeeper := mokita.WasmKeeper

	storeReflectCode(t, ctx, mokita, addr)

	cInfo := wasmKeeper.GetCodeInfo(ctx, 1)
	require.NotNil(t, cInfo)

	return mokita, ctx
}

func TestQueryFullDenom(t *testing.T) {
	actor := RandomAccountAddress()
	mokita, ctx := SetupCustomApp(t, actor)

	reflect := instantiateReflectContract(t, ctx, mokita, actor)
	require.NotEmpty(t, reflect)

	// query full denom
	query := bindings.MokitaQuery{
		FullDenom: &bindings.FullDenom{
			CreatorAddr: reflect.String(),
			Subdenom:    "ustart",
		},
	}
	resp := bindings.FullDenomResponse{}
	queryCustom(t, ctx, mokita, reflect, query, &resp)

	expected := fmt.Sprintf("factory/%s/ustart", reflect.String())
	require.EqualValues(t, expected, resp.Denom)
}

func TestQueryPool(t *testing.T) {
	actor := RandomAccountAddress()
	mokita, ctx := SetupCustomApp(t, actor)

	fundAccount(t, ctx, mokita, actor, defaultFunds)

	poolFunds := []sdk.Coin{
		sdk.NewInt64Coin("umoki", 12000000),
		sdk.NewInt64Coin("ustar", 240000000),
	}
	// 2 star to 1 moki
	starPool := preparePool(t, ctx, mokita, actor, poolFunds)

	pool2Funds := []sdk.Coin{
		sdk.NewInt64Coin("uatom", 6000000),
		sdk.NewInt64Coin("umoki", 12000000),
	}
	// 2 star to 1 moki
	atomPool := preparePool(t, ctx, mokita, actor, pool2Funds)

	reflect := instantiateReflectContract(t, ctx, mokita, actor)
	require.NotEmpty(t, reflect)

	// query pool state
	query := bindings.MokitaQuery{
		PoolState: &bindings.PoolState{PoolId: starPool},
	}
	resp := bindings.PoolStateResponse{}
	queryCustom(t, ctx, mokita, reflect, query, &resp)
	expected := wasmbinding.ConvertSdkCoinsToWasmCoins(poolFunds)
	require.EqualValues(t, expected, resp.Assets)
	assertValidShares(t, resp.Shares, starPool)

	// query second pool state
	query = bindings.MokitaQuery{
		PoolState: &bindings.PoolState{PoolId: atomPool},
	}
	resp = bindings.PoolStateResponse{}
	queryCustom(t, ctx, mokita, reflect, query, &resp)
	expected = wasmbinding.ConvertSdkCoinsToWasmCoins(pool2Funds)
	require.EqualValues(t, expected, resp.Assets)
	assertValidShares(t, resp.Shares, atomPool)
}

func TestQuerySpotPrice(t *testing.T) {
	actor := RandomAccountAddress()
	mokita, ctx := SetupCustomApp(t, actor)
	swapFee := 0. // FIXME: Set / support an actual fee
	epsilon := 1e-6

	fundAccount(t, ctx, mokita, actor, defaultFunds)

	poolFunds := []sdk.Coin{
		sdk.NewInt64Coin("umoki", 12000000),
		sdk.NewInt64Coin("ustar", 240000000),
	}
	// 20 star to 1 moki
	starPool := preparePool(t, ctx, mokita, actor, poolFunds)

	reflect := instantiateReflectContract(t, ctx, mokita, actor)
	require.NotEmpty(t, reflect)

	// query spot price
	query := bindings.MokitaQuery{
		SpotPrice: &bindings.SpotPrice{
			Swap: bindings.Swap{
				PoolId:   starPool,
				DenomIn:  "ustar",
				DenomOut: "umoki",
			},
			WithSwapFee: false,
		},
	}
	resp := bindings.SpotPriceResponse{}
	queryCustom(t, ctx, mokita, reflect, query, &resp)

	price, err := strconv.ParseFloat(resp.Price, 64)
	require.NoError(t, err)

	umoki, err := poolFunds[0].Amount.ToDec().Float64()
	require.NoError(t, err)
	ustar, err := poolFunds[1].Amount.ToDec().Float64()
	require.NoError(t, err)

	expected := ustar / umoki
	require.InEpsilonf(t, expected, price, epsilon, fmt.Sprintf("Outside of tolerance (%f)", epsilon))

	// and the reverse conversion (with swap fee)
	// query spot price
	query = bindings.MokitaQuery{
		SpotPrice: &bindings.SpotPrice{
			Swap: bindings.Swap{
				PoolId:   starPool,
				DenomIn:  "umoki",
				DenomOut: "ustar",
			},
			WithSwapFee: true,
		},
	}
	resp = bindings.SpotPriceResponse{}
	queryCustom(t, ctx, mokita, reflect, query, &resp)

	price, err = strconv.ParseFloat(resp.Price, 32)
	require.NoError(t, err)

	expected = 1. / expected
	require.InEpsilonf(t, expected+swapFee, price, epsilon, fmt.Sprintf("Outside of tolerance (%f)", epsilon))
}

func TestQueryEstimateSwap(t *testing.T) {
	actor := RandomAccountAddress()
	mokita, ctx := SetupCustomApp(t, actor)
	epsilon := 2e-3

	fundAccount(t, ctx, mokita, actor, defaultFunds)

	poolFunds := []sdk.Coin{
		sdk.NewInt64Coin("umoki", 12000000),
		sdk.NewInt64Coin("ustar", 240000000),
	}
	// 2 star to 1 moki
	starPool := preparePool(t, ctx, mokita, actor, poolFunds)

	reflect := instantiateReflectContract(t, ctx, mokita, actor)
	require.NotEmpty(t, reflect)

	// The contract/sender needs to have funds for estimating the price
	fundAccount(t, ctx, mokita, reflect, defaultFunds)

	// Estimate swap rate
	umoki, err := poolFunds[0].Amount.ToDec().Float64()
	require.NoError(t, err)
	ustar, err := poolFunds[1].Amount.ToDec().Float64()
	require.NoError(t, err)
	swapRate := ustar / umoki

	// Query estimate cost (Exact in. No route)
	amountIn := sdk.NewInt(10000)
	query := bindings.MokitaQuery{
		EstimateSwap: &bindings.EstimateSwap{
			Sender: reflect.String(),
			First: bindings.Swap{
				PoolId:   starPool,
				DenomIn:  "umoki",
				DenomOut: "ustar",
			},
			Route: []bindings.Step{},
			Amount: bindings.SwapAmount{
				In: &amountIn,
			},
		},
	}
	resp := bindings.EstimatePriceResponse{}
	queryCustom(t, ctx, mokita, reflect, query, &resp)
	require.NotNil(t, resp.Amount.Out)
	require.Nil(t, resp.Amount.In)
	cost, err := (*resp.Amount.Out).ToDec().Float64()
	require.NoError(t, err)

	amount, err := amountIn.ToDec().Float64()
	require.NoError(t, err)
	expected := amount * swapRate // out
	require.InEpsilonf(t, expected, cost, epsilon, fmt.Sprintf("Outside of tolerance (%f)", epsilon))

	// And the other way around
	// Query estimate cost (Exact out. No route)
	amountOut := sdk.NewInt(10000)
	query = bindings.MokitaQuery{
		EstimateSwap: &bindings.EstimateSwap{
			Sender: reflect.String(),
			First: bindings.Swap{
				PoolId:   starPool,
				DenomIn:  "umoki",
				DenomOut: "ustar",
			},
			Route: []bindings.Step{},
			Amount: bindings.SwapAmount{
				Out: &amountOut,
			},
		},
	}
	resp = bindings.EstimatePriceResponse{}
	queryCustom(t, ctx, mokita, reflect, query, &resp)
	require.NotNil(t, resp.Amount.In)
	require.Nil(t, resp.Amount.Out)
	cost, err = (*resp.Amount.In).ToDec().Float64()
	require.NoError(t, err)

	amount, err = amountOut.ToDec().Float64()
	require.NoError(t, err)
	expected = amount * 1. / swapRate
	require.InEpsilonf(t, expected, cost, epsilon, fmt.Sprintf("Outside of tolerance (%f)", epsilon))
}

type ReflectQuery struct {
	Chain *ChainRequest `json:"chain,omitempty"`
}

type ChainRequest struct {
	Request wasmvmtypes.QueryRequest `json:"request"`
}

type ChainResponse struct {
	Data []byte `json:"data"`
}

func queryCustom(t *testing.T, ctx sdk.Context, mokita *app.MokitaApp, contract sdk.AccAddress, request bindings.MokitaQuery, response interface{}) {
	msgBz, err := json.Marshal(request)
	require.NoError(t, err)

	query := ReflectQuery{
		Chain: &ChainRequest{
			Request: wasmvmtypes.QueryRequest{Custom: msgBz},
		},
	}
	queryBz, err := json.Marshal(query)
	require.NoError(t, err)

	resBz, err := mokita.WasmKeeper.QuerySmart(ctx, contract, queryBz)
	require.NoError(t, err)
	var resp ChainResponse
	err = json.Unmarshal(resBz, &resp)
	require.NoError(t, err)
	err = json.Unmarshal(resp.Data, response)
	require.NoError(t, err)
}

func assertValidShares(t *testing.T, shares wasmvmtypes.Coin, poolID uint64) {
	// sanity check: check the denom and ensure at least 18 decimal places
	denom := fmt.Sprintf("gamm/pool/%d", poolID)
	require.Equal(t, denom, shares.Denom)
	require.Greater(t, len(shares.Amount), 18)
}

func storeReflectCode(t *testing.T, ctx sdk.Context, mokita *app.MokitaApp, addr sdk.AccAddress) {
	govKeeper := mokita.GovKeeper
	wasmCode, err := os.ReadFile("../testdata/moki_reflect.wasm")
	require.NoError(t, err)

	src := wasmtypes.StoreCodeProposalFixture(func(p *wasmtypes.StoreCodeProposal) {
		p.RunAs = addr.String()
		p.WASMByteCode = wasmCode
	})

	// when stored
	storedProposal, err := govKeeper.SubmitProposal(ctx, src, false)
	require.NoError(t, err)

	// and proposal execute
	handler := govKeeper.Router().GetRoute(storedProposal.ProposalRoute())
	err = handler(ctx, storedProposal.GetContent())
	require.NoError(t, err)
}

func instantiateReflectContract(t *testing.T, ctx sdk.Context, mokita *app.MokitaApp, funder sdk.AccAddress) sdk.AccAddress {
	initMsgBz := []byte("{}")
	contractKeeper := keeper.NewDefaultPermissionKeeper(mokita.WasmKeeper)
	codeID := uint64(1)
	addr, _, err := contractKeeper.Instantiate(ctx, codeID, funder, funder, initMsgBz, "demo contract", nil)
	require.NoError(t, err)

	return addr
}

func fundAccount(t *testing.T, ctx sdk.Context, mokita *app.MokitaApp, addr sdk.AccAddress, coins sdk.Coins) {
	err := simapp.FundAccount(
		mokita.BankKeeper,
		ctx,
		addr,
		coins,
	)
	require.NoError(t, err)
}

func preparePool(t *testing.T, ctx sdk.Context, mokita *app.MokitaApp, addr sdk.AccAddress, funds []sdk.Coin) uint64 {
	var assets []balancer.PoolAsset
	for _, coin := range funds {
		assets = append(assets, balancer.PoolAsset{
			Weight: sdk.NewInt(100),
			Token:  coin,
		})
	}

	poolParams := balancer.PoolParams{
		SwapFee: sdk.NewDec(0),
		ExitFee: sdk.NewDec(0),
	}

	msg := balancer.NewMsgCreateBalancerPool(addr, poolParams, assets, "")
	poolId, err := mokita.GAMMKeeper.CreatePool(ctx, &msg)
	require.NoError(t, err)
	return poolId
}