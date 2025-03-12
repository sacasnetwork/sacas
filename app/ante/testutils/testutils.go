// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)
package testutils

import (
	"math"

	"github.com/stretchr/testify/suite"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"

	"github.com/sacasnetwork/sacas/v1/app"
	ante "github.com/sacasnetwork/sacas/v1/app/ante"
	evmante "github.com/sacasnetwork/sacas/v1/app/ante/evm"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/factory"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/grpc"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/keyring"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/network"
	"github.com/sacasnetwork/sacas/v1/types"
	evmtypes "github.com/sacasnetwork/sacas/v1/x/evm/types"
	feemarkettypes "github.com/sacasnetwork/sacas/v1/x/feemarket/types"
)

type AnteTestSuite struct {
	suite.Suite

	network   *network.UnitTestNetwork
	handler   grpc.Handler
	keyring   keyring.Keyring
	factory   factory.TxFactory
	clientCtx client.Context

	anteHandler     sdk.AnteHandler
	enableFeemarket bool
	baseFee         *sdkmath.Int
	enableLondonHF  bool
	evmParamsOption func(*evmtypes.Params)
}

const TestGasLimit uint64 = 100000

func (suite *AnteTestSuite) SetupTest() {
	keys := keyring.New(2)

	customGenesis := network.CustomGenesisState{}
	feemarketGenesis := feemarkettypes.DefaultGenesisState()
	if suite.enableFeemarket {
		feemarketGenesis.Params.EnableHeight = 1
		feemarketGenesis.Params.NoBaseFee = false
	} else {
		feemarketGenesis.Params.NoBaseFee = true
	}
	if suite.baseFee != nil {
		feemarketGenesis.Params.BaseFee = *suite.baseFee
	}
	customGenesis[feemarkettypes.ModuleName] = feemarketGenesis

	evmGenesis := evmtypes.DefaultGenesisState()
	if !suite.enableLondonHF {
		maxInt := sdkmath.NewInt(math.MaxInt64)
		evmGenesis.Params.ChainConfig.LondonBlock = &maxInt
		evmGenesis.Params.ChainConfig.ArrowGlacierBlock = &maxInt
		evmGenesis.Params.ChainConfig.GrayGlacierBlock = &maxInt
		evmGenesis.Params.ChainConfig.MergeNetsplitBlock = &maxInt
		evmGenesis.Params.ChainConfig.ShanghaiBlock = &maxInt
		evmGenesis.Params.ChainConfig.CancunBlock = &maxInt
	}
	if suite.evmParamsOption != nil {
		suite.evmParamsOption(&evmGenesis.Params)
	}
	customGenesis[evmtypes.ModuleName] = evmGenesis

	// set block max gas to be less than maxUint64
	cp := app.DefaultConsensusParams
	cp.Block.MaxGas = 1000000000000000000
	customGenesis[consensustypes.ModuleName] = cp

	nw := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keys.GetAllAccAddrs()...),
		network.WithCustomGenesis(customGenesis),
	)
	gh := grpc.NewIntegrationHandler(nw)
	tf := factory.New(nw, gh)

	suite.network = nw
	suite.factory = tf
	suite.handler = gh
	suite.keyring = keys

	encodingConfig := nw.GetEncodingConfig()

	suite.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)

	suite.Require().NotNil(suite.network.App.AppCodec())

	anteHandler := ante.NewAnteHandler(ante.HandlerOptions{
		Cdc:                    suite.network.App.AppCodec(),
		AccountKeeper:          suite.network.App.AccountKeeper,
		BankKeeper:             suite.network.App.BankKeeper,
		DistributionKeeper:     suite.network.App.DistrKeeper,
		EvmKeeper:              suite.network.App.EvmKeeper,
		FeegrantKeeper:         suite.network.App.FeeGrantKeeper,
		IBCKeeper:              suite.network.App.IBCKeeper,
		StakingKeeper:          suite.network.App.StakingKeeper,
		FeeMarketKeeper:        suite.network.App.FeeMarketKeeper,
		SignModeHandler:        encodingConfig.TxConfig.SignModeHandler(),
		SigGasConsumer:         ante.SigVerificationGasConsumer,
		ExtensionOptionChecker: types.HasDynamicFeeExtensionOption,
		TxFeeChecker:           evmante.NewDynamicFeeChecker(suite.network.App.EvmKeeper),
	})

	suite.anteHandler = anteHandler
}

func (suite *AnteTestSuite) WithFeemarketEnabled(enabled bool) {
	suite.enableFeemarket = enabled
}

func (suite *AnteTestSuite) WithLondonHardForkEnabled(enabled bool) {
	suite.enableLondonHF = enabled
}

func (suite *AnteTestSuite) WithBaseFee(baseFee *sdkmath.Int) {
	suite.baseFee = baseFee
}

func (suite *AnteTestSuite) WithEvmParamsOptions(evmParamsOpts func(*evmtypes.Params)) {
	suite.evmParamsOption = evmParamsOpts
}

func (suite *AnteTestSuite) ResetEvmParamsOptions() {
	suite.evmParamsOption = nil
}

func (suite *AnteTestSuite) GetKeyring() keyring.Keyring {
	return suite.keyring
}

func (suite *AnteTestSuite) GetTxFactory() factory.TxFactory {
	return suite.factory
}

func (suite *AnteTestSuite) GetNetwork() *network.UnitTestNetwork {
	return suite.network
}

func (suite *AnteTestSuite) GetClientCtx() client.Context {
	return suite.clientCtx
}

func (suite *AnteTestSuite) GetAnteHandler() sdk.AnteHandler {
	return suite.anteHandler
}
