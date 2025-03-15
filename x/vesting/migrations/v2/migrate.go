// Copyright SacDev0S . (SacaS)
package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	accounttypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	v1vestingtypes "github.com/sacasnetwork/sacas/v15/x/vesting/migrations/types"
	vestingtypes "github.com/sacasnetwork/sacas/v15/x/vesting/types"
)

// MigrateStore migrates the x/vesting module state from the consensus version 1 to
// version 2. Specifically, it converts all vesting accounts from their v1 proto definitions to v2.
func MigrateStore(
	ctx sdk.Context,
	ak vestingtypes.AccountKeeper,
) error {
	ak.IterateAccounts(ctx, func(account accounttypes.AccountI) bool {
		if oldAccount, ok := account.(*v1vestingtypes.ClawbackVestingAccount); ok {
			newAccount := &vestingtypes.ClawbackVestingAccount{
				BaseVestingAccount: oldAccount.BaseVestingAccount,
				FunderAddress:      oldAccount.FunderAddress,
				StartTime:          oldAccount.StartTime,
				LockupPeriods:      oldAccount.LockupPeriods,
				VestingPeriods:     oldAccount.VestingPeriods,
			}
			ak.RemoveAccount(ctx, oldAccount)
			ak.SetAccount(ctx, newAccount)
		}

		return false
	})

	return nil
}
