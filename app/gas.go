// Copyright SacDev0S . (SacaS)

package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// MainnetMinGasPrices defines 20B asacas (or atsacas) as the minimum gas price value on the fee market module.
	// See https://commonwealth.im/sacas/discussion/5073-global-min-gas-price-value-for-cosmos-sdk-and-evm-transaction-choosing-a-value for reference
	MainnetMinGasPrices = sdk.NewDec(20_000_000_000)
	// MainnetMinGasMultiplier defines the min gas multiplier value on the fee market module.
	// 50% of the leftover gas will be refunded
	MainnetMinGasMultiplier = sdk.NewDecWithPrec(5, 1)
)
