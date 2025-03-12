// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)

package ante_test

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/network"
	evmante "github.com/sacasnetwork/sacas/v1/x/evm/ante"
)

func (suite *EvmAnteTestSuite) TestBuildEvmExecutionCtx() {
	network := network.New()

	ctx := evmante.BuildEvmExecutionCtx(network.GetContext())

	suite.Equal(storetypes.GasConfig{}, ctx.KVGasConfig())
	suite.Equal(storetypes.GasConfig{}, ctx.TransientKVGasConfig())
}
