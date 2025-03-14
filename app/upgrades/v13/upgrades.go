// Copyright SacDev0S . (SacaS)

package v13

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	evmkeeper "github.com/sacasnetwork/sacas/v15/x/evm/keeper"
	evmtypes "github.com/sacasnetwork/sacas/v15/x/evm/types"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v13
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	ek evmkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		if err := setPrecompilesParams(ctx, ek); err != nil {
			logger.Error("error while setting precompiles parameters", "error", err)
		}
		// Leave modules are as-is to avoid running InitGenesis.
		logger.Debug("running module migrations ...")
		return mm.RunMigrations(ctx, configurator, vm)
	}
}

func setPrecompilesParams(ctx sdk.Context, ek evmkeeper.Keeper) error {
	params := ek.GetParams(ctx)
	params.ActivePrecompiles = evmtypes.AvailableEVMExtensions
	return ek.SetParams(ctx, params)
}
