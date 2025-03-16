// Copyright SacDev0S . (SacaS)

package factory

import (
	"github.com/sacasnetwork/sacas/v20/testutil/integration/sacas/grpc"
	"github.com/sacasnetwork/sacas/v20/testutil/integration/sacas/network"
)

const (
	GasAdjustment = float64(1.7)
)

// CoreTxFactory is the interface that wraps the methods
// to build and broadcast cosmos transactions, and also
// includes module-specific transactions
type CoreTxFactory interface {
	BaseTxFactory
	DistributionTxFactory
	StakingTxFactory
	FundTxFactory
}

var _ CoreTxFactory = (*IntegrationTxFactory)(nil)

// IntegrationTxFactory is a helper struct to build and broadcast transactions
// to the network on integration tests. This is to simulate the behavior of a real user.
type IntegrationTxFactory struct {
	BaseTxFactory
	DistributionTxFactory
	StakingTxFactory
	FundTxFactory
}

// New creates a new IntegrationTxFactory instance
func New(
	network network.Network,
	grpcHandler grpc.Handler,
) CoreTxFactory {
	bf := newBaseTxFactory(network, grpcHandler)
	return &IntegrationTxFactory{
		bf,
		newDistrTxFactory(bf),
		newStakingTxFactory(bf),
		newFundTxFactory(bf),
	}
}
