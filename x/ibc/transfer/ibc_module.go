// Copyright SacDev0S . (SacaS)

package transfer

import (
	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	"github.com/sacasnetwork/sacas/v15/x/ibc/transfer/keeper"
)

var _ porttypes.IBCModule = IBCModule{}

// IBCModule implements the ICS26 interface for transfer given the transfer keeper.
type IBCModule struct {
	*ibctransfer.IBCModule
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k keeper.Keeper) IBCModule {
	transferModule := ibctransfer.NewIBCModule(*k.Keeper)
	return IBCModule{
		IBCModule: &transferModule,
	}
}
