// Copyright SacDev0S . (SacaS)

package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	evmante "github.com/sacasnetwork/sacas/v20/app/ante/evm"
)

func newMonoEVMAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		evmante.NewMonoDecorator(
			options.AccountKeeper,
			options.BankKeeper,
			options.FeeMarketKeeper,
			options.EvmKeeper,
			options.DistributionKeeper,
			options.StakingKeeper,
			options.MaxTxGasWanted,
		),
	)
}
