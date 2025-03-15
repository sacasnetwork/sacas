// Copyright SacDev0S . (SacaS)
package network

import (
	evmtypes "github.com/sacasnetwork/sacas/v15/x/evm/types"
	infltypes "github.com/sacasnetwork/sacas/v15/x/inflation/types"
	revtypes "github.com/sacasnetwork/sacas/v15/x/revenue/v1/types"
)

func (n *IntegrationNetwork) UpdateEvmParams(params evmtypes.Params) error {
	return n.app.EvmKeeper.SetParams(n.ctx, params)
}

func (n *IntegrationNetwork) UpdateRevenueParams(params revtypes.Params) error {
	return n.app.RevenueKeeper.SetParams(n.ctx, params)
}

func (n *IntegrationNetwork) UpdateInflationParams(params infltypes.Params) error {
	return n.app.InflationKeeper.SetParams(n.ctx, params)
}
