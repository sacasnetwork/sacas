package keeper_test

import (
	"testing"

	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/ginkgo/v2"
	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/gomega"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcgotesting "github.com/cosmos/ibc-go/v7/testing"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sacasnetwork/sacas/v15/app"
	ibctesting "github.com/sacasnetwork/sacas/v15/ibc/testing"
	"github.com/sacasnetwork/sacas/v15/x/erc20/types"
	evm "github.com/sacasnetwork/sacas/v15/x/evm/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx              sdk.Context
	app              *app.Sacas
	queryClientEvm   evm.QueryClient
	queryClient      types.QueryClient
	address          common.Address
	consAddress      sdk.ConsAddress
	clientCtx        client.Context //nolint:unused
	ethSigner        ethtypes.Signer
	priv             cryptotypes.PrivKey
	validator        stakingtypes.Validator
	signer           keyring.Signer
	mintFeeCollector bool

	coordinator *ibcgotesting.Coordinator

	// testing chains used for convenience and readability
	SacasChain      *ibcgotesting.TestChain
	IBCOsmosisChain *ibcgotesting.TestChain
	IBCCosmosChain  *ibcgotesting.TestChain

	pathOsmosisSacas  *ibctesting.Path
	pathCosmosSacas   *ibctesting.Path
	pathOsmosisCosmos *ibctesting.Path

	suiteIBCTesting bool
}

var (
	s *KeeperTestSuite
	// sendAndReceiveMsgFee corresponds to the fees paid on Sacas chain when calling the SendAndReceive function
	// This function makes 3 cosmos txs under the hood
	sendAndReceiveMsgFee = sdk.NewInt(ibctesting.DefaultFeeAmt * 3)
	// sendBackCoinsFee corresponds to the fees paid on Sacas chain when calling the SendBackCoins function
	// or calling the SendAndReceive from another chain to Sacas
	// This function makes 2 cosmos txs under the hood
	sendBackCoinsFee = sdk.NewInt(ibctesting.DefaultFeeAmt * 2)
)

func TestKeeperTestSuite(t *testing.T) {
	s = new(KeeperTestSuite)
	suite.Run(t, s)

	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.DoSetupTest(suite.T())
}
