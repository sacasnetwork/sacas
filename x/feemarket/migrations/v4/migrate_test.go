// Copyright 2022 Sacas Foundation
// This file is part of the Sacas Network packages.
//
// Sacas is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Sacas packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Sacas packages. If not, see https://github.com/sacasnetwork/sacas/blob/main/LICENSE
package v4_test

import (
	"testing"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sacasnetwork/sacas/v20/encoding"
	v4 "github.com/sacasnetwork/sacas/v20/x/feemarket/migrations/v4"
	"github.com/sacasnetwork/sacas/v20/x/feemarket/types"
	"github.com/stretchr/testify/require"
)

type mockSubspace struct {
	ps types.Params
}

func newMockSubspaceEmpty() mockSubspace {
	return mockSubspace{}
}

func newMockSubspace(ps types.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSetIfExists(_ sdk.Context, ps types.LegacyParams) {
	*ps.(*types.Params) = ms.ps
}

func TestMigrate(t *testing.T) {
	encCfg := encoding.MakeConfig()
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	kvStore := ctx.KVStore(storeKey)

	legacySubspaceEmpty := newMockSubspaceEmpty()
	require.Error(t, v4.MigrateStore(ctx, storeKey, legacySubspaceEmpty, cdc))

	legacySubspace := newMockSubspace(types.DefaultParams())
	require.NoError(t, v4.MigrateStore(ctx, storeKey, legacySubspace, cdc))

	paramsBz := kvStore.Get(types.ParamsKey)
	var params types.Params
	cdc.MustUnmarshal(paramsBz, &params)

	require.Equal(t, params, legacySubspace.ps)
}
