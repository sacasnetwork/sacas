// Copyright SacDev0S . (SacaS)

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v3 "github.com/sacasnetwork/sacas/v15/x/claims/migrations/v3"
	"github.com/sacasnetwork/sacas/v15/x/claims/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper         Keeper
	legacySubspace types.Subspace
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper, legacySubspace types.Subspace) Migrator {
	return Migrator{
		keeper:         keeper,
		legacySubspace: legacySubspace,
	}
}

// Migrate2to3 migrates the store from consensus version 2 to 3
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.MigrateStore(ctx, m.keeper.storeKey, m.legacySubspace, m.keeper.cdc)
}
