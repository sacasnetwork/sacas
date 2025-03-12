// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)
package evm_test

import (
	storetypes "cosmossdk.io/store/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/sacasnetwork/sacas/v1/app/ante/evm"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/factory"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/grpc"
	testkeyring "github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/keyring"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/network"
	integrationutils "github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/utils"
)

func (suite *EvmAnteTestSuite) TestCheckGasWanted() {
	keyring := testkeyring.New(1)
	unitNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	grpcHandler := grpc.NewIntegrationHandler(unitNetwork)
	txFactory := factory.New(unitNetwork, grpcHandler)
	commonGasLimit := uint64(100000)

	testCases := []struct {
		name                       string
		expectedError              error
		getCtx                     func() sdktypes.Context
		isLondon                   bool
		expectedTransientGasWanted uint64
	}{
		{
			name:          "success: if isLondon false it should not error",
			expectedError: nil,
			getCtx: func() sdktypes.Context {
				// Even if the gasWanted is more than the blockGasLimit, it should not error
				blockMeter := storetypes.NewGasMeter(commonGasLimit - 10000)
				return unitNetwork.GetContext().WithBlockGasMeter(blockMeter)
			},
			isLondon:                   false,
			expectedTransientGasWanted: 0,
		},
		{
			name:          "success: gasWanted is less than blockGasLimit",
			expectedError: nil,
			getCtx: func() sdktypes.Context {
				blockMeter := storetypes.NewGasMeter(commonGasLimit + 10000)
				return unitNetwork.GetContext().WithBlockGasMeter(blockMeter)
			},
			isLondon:                   true,
			expectedTransientGasWanted: commonGasLimit,
		},
		{
			name:          "fail: gasWanted is more than blockGasLimit",
			expectedError: errortypes.ErrOutOfGas,
			getCtx: func() sdktypes.Context {
				blockMeter := storetypes.NewGasMeter(commonGasLimit - 10000)
				return unitNetwork.GetContext().WithBlockGasMeter(blockMeter)
			},
			isLondon:                   true,
			expectedTransientGasWanted: 0,
		},
		{
			name:          "success: gasWanted is less than blockGasLimit and basefee param is disabled",
			expectedError: nil,
			getCtx: func() sdktypes.Context {
				// Set basefee param to false
				feeMarketParams, err := grpcHandler.GetFeeMarketParams()
				suite.Require().NoError(err)
				feeMarketParams.Params.NoBaseFee = true
				err = integrationutils.UpdateFeeMarketParams(integrationutils.UpdateParamsInput{
					Tf:      txFactory,
					Network: unitNetwork,
					Pk:      keyring.GetPrivKey(0),
					Params:  feeMarketParams.Params,
				})
				suite.Require().NoError(err)

				blockMeter := storetypes.NewGasMeter(commonGasLimit + 10000)
				return unitNetwork.GetContext().WithBlockGasMeter(blockMeter)
			},
			isLondon:                   true,
			expectedTransientGasWanted: 0,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			sender := keyring.GetKey(0)
			txArgs, err := txFactory.GenerateDefaultTxTypeArgs(
				sender.Addr,
				suite.ethTxType,
			)
			suite.Require().NoError(err)
			txArgs.GasLimit = commonGasLimit
			tx, err := txFactory.GenerateSignedEthTx(sender.Priv, txArgs)
			suite.Require().NoError(err)

			ctx := tc.getCtx()

			// Function under test
			err = evm.CheckGasWanted(
				ctx,
				unitNetwork.App.FeeMarketKeeper,
				tx,
				tc.isLondon,
			)

			if tc.expectedError != nil {
				suite.Require().Error(err)
				suite.Contains(err.Error(), tc.expectedError.Error())
			} else {
				suite.Require().NoError(err)
				transientGasWanted := unitNetwork.App.FeeMarketKeeper.GetTransientGasWanted(
					unitNetwork.GetContext(),
				)
				suite.Require().Equal(tc.expectedTransientGasWanted, transientGasWanted)
			}

			// Start from a fresh block and ctx
			err = unitNetwork.NextBlock()
			suite.Require().NoError(err)
		})
	}
}
