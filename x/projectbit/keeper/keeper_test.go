package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/math"

	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/store/types"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"projectBit/x/projectbit/keeper"
	module "projectBit/x/projectbit/module"
	"projectBit/x/projectbit/types"
)

// Mock bank
type mockBankKeeper struct{}

// Everyone have infinite money for test
func (mockBankKeeper) SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin("token", math.NewInt(1000000)))
}

// Send coins without effort (sucess/nil)
func (mockBankKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	return nil
}

type fixture struct {
	ctx          context.Context
	keeper       keeper.Keeper
	addressCodec address.Codec
}

func initFixture(t *testing.T) *fixture {
	t.Helper()

	encCfg := moduletestutil.MakeTestEncodingConfig(module.AppModule{})
	addressCodec := addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	storeService := runtime.NewKVStoreService(storeKey)
	ctx := testutil.DefaultContextWithDB(t, storeKey, storetypes.NewTransientStoreKey("transient_test")).Ctx

	authority := authtypes.NewModuleAddress(types.GovModuleName)

	bankKeeper := mockBankKeeper{}

	k := keeper.NewKeeper(
		storeService,
		encCfg.Codec,
		addressCodec,
		authority,
		bankKeeper,
	)

	// Initialize params
	if err := k.Params.Set(ctx, types.DefaultParams()); err != nil {
		t.Fatalf("failed to set params: %v", err)
	}

	return &fixture{
		ctx:          ctx,
		keeper:       k,
		addressCodec: addressCodec,
	}
}
