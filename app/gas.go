// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0( Sacas Network )

package app

import "cosmossdk.io/math"

var (
	// MainnetMinGasPrices defines 1B sac (or atsacas) as the minimum gas price value on the fee market module.
	// See https://commonwealth.im/evmos/discussion/5073-global-min-gas-price-value-for-cosmos-sdk-and-evm-transaction-choosing-a-value for reference
	MainnetMinGasPrices = math.LegacyNewDec(1_000_000_000)
	// MainnetMinGasMultiplier defines the min gas multiplier value on the fee market module.
	// 50% of the leftover gas will be refunded
	MainnetMinGasMultiplier = math.LegacyNewDecWithPrec(8, 1)
)
