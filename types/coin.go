// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)
package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AttoSacas defines the default coin denomination used in Sacas in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Sacas.
	AttoSacas string = "asacas"

	// BaseDenomUnit defines the base denomination unit for Sacas.
	// 1 sacas = 1x10^{BaseDenomUnit} asacas
	BaseDenomUnit = 18

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewSacasCoin is a utility function that returns an "asacas" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewSacasCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(AttoSacas, amount)
}

// NewSacasDecCoin is a utility function that returns an "asacas" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewSacasDecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoSacas, amount)
}

// NewSacasCoinInt64 is a utility function that returns an "asacas" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewSacasCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoSacas, amount)
}
