package recovery_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cometbft/cometbft/crypto/tmhash"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmversion "github.com/cometbft/cometbft/proto/tendermint/version"
	"github.com/cometbft/cometbft/version"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	utiltx "github.com/sacasnetwork/sacas/v15/testutil/tx"
	"github.com/sacasnetwork/sacas/v15/utils"
	feemarkettypes "github.com/sacasnetwork/sacas/v15/x/feemarket/types"

	"github.com/sacasnetwork/sacas/v15/app"
	"github.com/sacasnetwork/sacas/v15/x/recovery"
	"github.com/sacasnetwork/sacas/v15/x/recovery/types"
)

type GenesisTestSuite struct {
	suite.Suite

	ctx sdk.Context

	app     *app.Sacas
	genesis types.GenesisState
}

func (suite *GenesisTestSuite) SetupTest() {
	// consensus key
	consAddress := sdk.ConsAddress(utiltx.GenerateAddress().Bytes())

	chainID := utils.TestnetChainID + "-1"
	suite.app = app.Setup(false, feemarkettypes.DefaultGenesisState(), chainID)
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{
		Height:          1,
		ChainID:         chainID,
		Time:            time.Now().UTC(),
		ProposerAddress: consAddress.Bytes(),

		Version: tmversion.Consensus{
			Block: version.BlockProtocol,
		},
		LastBlockId: tmproto.BlockID{
			Hash: tmhash.Sum([]byte("block_id")),
			PartSetHeader: tmproto.PartSetHeader{
				Total: 11,
				Hash:  tmhash.Sum([]byte("partset_header")),
			},
		},
		AppHash:            tmhash.Sum([]byte("app")),
		DataHash:           tmhash.Sum([]byte("data")),
		EvidenceHash:       tmhash.Sum([]byte("evidence")),
		ValidatorsHash:     tmhash.Sum([]byte("validators")),
		NextValidatorsHash: tmhash.Sum([]byte("next_validators")),
		ConsensusHash:      tmhash.Sum([]byte("consensus")),
		LastResultsHash:    tmhash.Sum([]byte("last_result")),
	})

	suite.genesis = *types.DefaultGenesisState()
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (suite *GenesisTestSuite) TestRecoveryInitGenesis() {
	testCases := []struct {
		name     string
		genesis  types.GenesisState
		expPanic bool
	}{
		{
			"default genesis",
			suite.genesis,
			false,
		},
		{
			"custom genesis - recovery disabled",
			types.GenesisState{
				Params: types.Params{
					EnableRecovery:        false,
					PacketTimeoutDuration: time.Hour * 10,
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			if tc.expPanic {
				suite.Require().Panics(func() {
					recovery.InitGenesis(suite.ctx, *suite.app.RecoveryKeeper, tc.genesis)
				})
			} else {
				suite.Require().NotPanics(func() {
					recovery.InitGenesis(suite.ctx, *suite.app.RecoveryKeeper, tc.genesis)
				})

				params := suite.app.RecoveryKeeper.GetParams(suite.ctx)
				suite.Require().Equal(tc.genesis.Params, params)
			}
		})
	}
}

func (suite *GenesisTestSuite) TestRecoveryExportGenesis() {
	recovery.InitGenesis(suite.ctx, *suite.app.RecoveryKeeper, suite.genesis)

	genesisExported := recovery.ExportGenesis(suite.ctx, *suite.app.RecoveryKeeper)
	suite.Require().Equal(genesisExported.Params, suite.genesis.Params)
}
