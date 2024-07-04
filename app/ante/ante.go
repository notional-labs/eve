package ante

import (
	"errors"
	"fmt"

	ibcante "github.com/cosmos/ibc-go/v8/modules/core/ante"
	"github.com/cosmos/ibc-go/v8/modules/core/keeper"
	feeabskeeper "github.com/osmosis-labs/fee-abstraction/v8/x/feeabs/keeper"
	feeabstypes "github.com/osmosis-labs/fee-abstraction/v8/x/feeabs/types"
	feemarketante "github.com/skip-mev/feemarket/x/feemarket/ante"
	feemarkettypes "github.com/skip-mev/feemarket/x/feemarket/types"

	corestoretypes "cosmossdk.io/core/store"
	circuitante "cosmossdk.io/x/circuit/ante"
	circuitkeeper "cosmossdk.io/x/circuit/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC
// channel keeper.
type HandlerOptions struct {
	ante.HandlerOptions

	IBCKeeper             *keeper.Keeper
	WasmConfig            *wasmTypes.WasmConfig
	WasmKeeper            *wasmkeeper.Keeper
	TXCounterStoreService corestoretypes.KVStoreService
	CircuitKeeper         *circuitkeeper.Keeper
	FeeAbskeeper          feeabskeeper.Keeper
	FeeMarketKeeper       feemarketante.FeeMarketKeeper
	AccountKeeper         feemarketante.AccountKeeper
}

// NewAnteHandler constructor
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, errors.New("account keeper is required for ante builder")
	}
	if options.BankKeeper == nil {
		return nil, errors.New("bank keeper is required for ante builder")
	}
	if options.SignModeHandler == nil {
		return nil, errors.New("sign mode handler is required for ante builder")
	}
	if options.WasmConfig == nil {
		return nil, errors.New("wasm config is required for ante builder")
	}
	if options.TXCounterStoreService == nil {
		return nil, errors.New("wasm store service is required for ante builder")
	}
	if options.CircuitKeeper == nil {
		return nil, errors.New("circuit keeper is required for ante builder")
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		wasmkeeper.NewLimitSimulationGasDecorator(options.WasmConfig.SimulationGasLimit), // after setup context to enforce limits early
		wasmkeeper.NewCountTXDecorator(options.TXCounterStoreService),
		wasmkeeper.NewGasRegisterDecorator(options.WasmKeeper.GetGasRegister()),
		circuitante.NewCircuitBreakerDecorator(options.CircuitKeeper),
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		feemarketante.NewFeeMarketCheckDecorator( // fee market check replaces fee deduct decorator
			options.FeeMarketKeeper,
			ante.NewDeductFeeDecorator(
				options.AccountKeeper,
				options.BankKeeper,
				options.FeegrantKeeper,
				options.TxFeeChecker,
			),
		), // fees are deducted in the fee market deduct post handler
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}

// DenomResolverImpl is Eve's implementation of x/feemarket's DenomResolver
type DenomResolverImpl struct {
	FeeabsKeeper  feeabskeeper.Keeper
	StakingKeeper feeabstypes.StakingKeeper
}

var _ feemarkettypes.DenomResolver = &DenomResolverImpl{}

// ConvertToDenom returns the equivalent DecCoin in a given denom.
// Return error if the conversion is not possible, or the coin's denom is not in the ExtraDenoms list.
// TODO: make error more descriptive
func (r *DenomResolverImpl) ConvertToDenom(ctx sdk.Context, coin sdk.DecCoin, denom string) (sdk.DecCoin, error) {
	// TODO: Assume that the bond denom is the native token
	bondDenom, err := r.StakingKeeper.BondDenom(ctx)
	if err != nil {
		return sdk.DecCoin{}, err
	}
	var hostZoneConfig feeabstypes.HostChainFeeAbsConfig

	// If the denom is the bond denom, convert `coin` to the native denom
	if denom == bondDenom {
		hostZoneConfig, found := r.FeeabsKeeper.GetHostZoneConfig(ctx, coin.Denom)
		if !found {
			return sdk.DecCoin{}, fmt.Errorf("error resolving denom")
		}
		amount, err := r.getIBCCoinFromNative(ctx, sdk.NewCoins(sdk.NewCoin(coin.Denom, coin.Amount.TruncateInt())), hostZoneConfig)
		if err != nil {
			return sdk.DecCoin{}, err
		}
		return sdk.NewDecCoinFromDec(denom, amount[0].Amount.ToLegacyDec()), nil
	}

	// If the denom is not the bond denom, convert the `coin` to the given denom
	hostZoneConfig, found := r.FeeabsKeeper.GetHostZoneConfig(ctx, denom)
	if !found {
		return sdk.DecCoin{}, fmt.Errorf("error resolving denom")
	}
	amount, err := r.FeeabsKeeper.CalculateNativeFromIBCCoins(ctx, sdk.NewCoins(sdk.NewCoin(denom, coin.Amount.TruncateInt())), hostZoneConfig)
	if err != nil {
		return sdk.DecCoin{}, err
	}
	return sdk.NewDecCoinFromDec(denom, amount[0].Amount.ToLegacyDec()), nil
}

// TODO: implement this method
// extra denoms should be all denoms that have been registered via governance(host zone)
func (r *DenomResolverImpl) ExtraDenoms(ctx sdk.Context) ([]string, error) {
	allHostZoneConfigs, err := r.FeeabsKeeper.GetAllHostZoneConfig(ctx)
	if err != nil {
		return nil, err
	}
	denoms := make([]string, 0, len(allHostZoneConfigs))
	for _, hostZoneConfig := range allHostZoneConfigs {
		denoms = append(denoms, hostZoneConfig.IbcDenom)
	}
	return denoms, nil
}

// //////////////////////////////////////
// Helper functions for DenomResolver //
// //////////////////////////////////////

func (r *DenomResolverImpl) getIBCCoinFromNative(ctx sdk.Context, nativeCoins sdk.Coins, chainConfig feeabstypes.HostChainFeeAbsConfig) (coins sdk.Coins, err error) {
	if len(nativeCoins) != 1 {
		return sdk.Coins{}, fmt.Errorf("expected exactly one native coin, got %d", len(nativeCoins))
	}

	nativeCoin := nativeCoins[0]

	twapRate, err := r.FeeabsKeeper.GetTwapRate(ctx, chainConfig.IbcDenom)
	if err != nil {
		return sdk.Coins{}, err
	}

	// Divide native amount by twap rate to get IBC amount
	ibcAmount := nativeCoin.Amount.ToLegacyDec().Quo(twapRate).RoundInt()
	ibcCoin := sdk.NewCoin(chainConfig.IbcDenom, ibcAmount)

	// Verify the resulting IBC coin
	err = r.verifyIBCCoins(ctx, sdk.NewCoins(ibcCoin))
	if err != nil {
		return sdk.Coins{}, err
	}

	return sdk.NewCoins(ibcCoin), nil
}

// return err if IBC token isn't in allowed_list
func (r *DenomResolverImpl) verifyIBCCoins(ctx sdk.Context, ibcCoins sdk.Coins) error {
	if ibcCoins.Len() != 1 {
		return feeabstypes.ErrInvalidIBCFees
	}

	ibcDenom := ibcCoins[0].Denom
	if r.FeeabsKeeper.HasHostZoneConfig(ctx, ibcDenom) {
		return nil
	}
	return feeabstypes.ErrUnsupportedDenom.Wrapf("unsupported denom: %s", ibcDenom)
}
