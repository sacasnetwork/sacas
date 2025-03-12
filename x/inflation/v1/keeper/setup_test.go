package keeper_test

import (
	"github.com/stretchr/testify/suite"

	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/factory"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/grpc"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/keyring"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/network"
	"github.com/sacasnetwork/sacas/v1/x/inflation/v1/types"
)

var denomMint = types.DefaultInflationDenom

type KeeperTestSuite struct {
	suite.Suite

	network *network.UnitTestNetwork
	handler grpc.Handler
	keyring keyring.Keyring
	factory factory.TxFactory
}

func (suite *KeeperTestSuite) SetupTest() {
	keys := keyring.New(2)
	nw := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keys.GetAllAccAddrs()...),
	)
	gh := grpc.NewIntegrationHandler(nw)
	tf := factory.New(nw, gh)

	suite.network = nw
	suite.factory = tf
	suite.handler = gh
	suite.keyring = keys
}
