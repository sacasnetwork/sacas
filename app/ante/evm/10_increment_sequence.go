// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)

package evm

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"

	evmtypes "github.com/sacasnetwork/sacas/v1/x/evm/types"
)

// IncrementNonce increments the sequence of the account.
func IncrementNonce(
	ctx sdk.Context,
	accountKeeper evmtypes.AccountKeeper,
	account sdk.AccountI,
	txNonce uint64,
) error {
	nonce := account.GetSequence()
	// we merged the nonce verification to nonce increment, so when tx includes multiple messages
	// with same sender, they'll be accepted.
	if txNonce != nonce {
		return errorsmod.Wrapf(
			errortypes.ErrInvalidSequence,
			"invalid nonce; got %d, expected %d", txNonce, nonce,
		)
	}

	nonce++

	if err := account.SetSequence(nonce); err != nil {
		return errorsmod.Wrapf(err, "failed to set sequence to %d", nonce)
	}

	accountKeeper.SetAccount(ctx, account)
	return nil
}
