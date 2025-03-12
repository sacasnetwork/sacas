package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/factory"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/grpc"
	testkeyring "github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/keyring"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/network"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	network     *network.UnitTestNetwork
	factory     factory.TxFactory
	grpcHandler grpc.Handler
	keyring     testkeyring.Keyring

	denom string
}

// SetupTest setup test environment
func (suite *KeeperTestSuite) SetupTest() {
	keyring := testkeyring.New(2)
	nw := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
		network.WithCustomBaseAppOpts(baseapp.SetMinGasPrices("10asacas")),
	)
	grpcHandler := grpc.NewIntegrationHandler(nw)
	txFactory := factory.New(nw, grpcHandler)

	ctx := nw.GetContext()
	sk := nw.App.StakingKeeper
	bondDenom, err := sk.BondDenom(ctx)
	if err != nil {
		panic(err)
	}

	suite.denom = bondDenom
	suite.factory = txFactory
	suite.grpcHandler = grpcHandler
	suite.keyring = keyring
	suite.network = nw
}
