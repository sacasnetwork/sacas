package keeper_test

import (
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/sacasnetwork/sacas/v15/app"
	"github.com/sacasnetwork/sacas/v15/testutil"
	utils "github.com/sacasnetwork/sacas/v15/utils"
	epochstypes "github.com/sacasnetwork/sacas/v15/x/epochs/types"
	evm "github.com/sacasnetwork/sacas/v15/x/evm/types"
	"github.com/sacasnetwork/sacas/v15/x/inflation/types"
)

// Test helpers
func (suite *KeeperTestSuite) DoSetupTest() {
	checkTx := false

	// init app
	chainID := utils.MainnetChainID + "-1"
	suite.app = app.Setup(checkTx, nil, chainID)

	// setup context
	// use same
	header := testutil.NewHeader(
		1, time.Now().UTC(), chainID, suite.consAddress, nil, nil,
	)
	suite.ctx = suite.app.BaseApp.NewContext(checkTx, header)

	// setup query helpers
	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.app.InflationKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)

	// Set epoch start time and height for all epoch identifiers
	identifiers := []string{epochstypes.WeekEpochID, epochstypes.DayEpochID}
	for _, identifier := range identifiers {
		epoch, found := suite.app.EpochsKeeper.GetEpochInfo(suite.ctx, identifier)
		suite.Require().True(found)
		epoch.StartTime = suite.ctx.BlockTime()
		epoch.CurrentEpochStartHeight = suite.ctx.BlockHeight()
		suite.app.EpochsKeeper.SetEpochInfo(suite.ctx, epoch)
	}
}

func (suite *KeeperTestSuite) Commit() {
	suite.CommitAfter(time.Nanosecond)
}

func (suite *KeeperTestSuite) CommitAfter(t time.Duration) {
	var err error
	suite.ctx, err = testutil.CommitAndCreateNewCtx(suite.ctx, suite.app, t, nil)
	suite.Require().NoError(err)
	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	evm.RegisterQueryServer(queryHelper, suite.app.EvmKeeper)
	suite.queryClientEvm = evm.NewQueryClient(queryHelper)
}
