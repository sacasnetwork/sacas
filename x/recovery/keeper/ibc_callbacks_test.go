package keeper_test

import (
	"fmt"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/sacasnetwork/sacas/v15/crypto/ethsecp256k1"
	"github.com/sacasnetwork/sacas/v15/testutil"
	utiltx "github.com/sacasnetwork/sacas/v15/testutil/tx"
	"github.com/stretchr/testify/mock"

	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibcgotesting "github.com/cosmos/ibc-go/v7/testing"
	ibcmock "github.com/cosmos/ibc-go/v7/testing/mock"

	claimstypes "github.com/sacasnetwork/sacas/v15/x/claims/types"
	incentivestypes "github.com/sacasnetwork/sacas/v15/x/incentives/types"
	"github.com/sacasnetwork/sacas/v15/x/recovery/keeper"
	"github.com/sacasnetwork/sacas/v15/x/recovery/types"
	vestingtypes "github.com/sacasnetwork/sacas/v15/x/vesting/types"
)

func (suite *KeeperTestSuite) TestOnRecvPacket() {
	// secp256k1 account
	secpPk := secp256k1.GenPrivKey()
	secpAddr := sdk.AccAddress(secpPk.PubKey().Address())
	secpAddrSacas := secpAddr.String()
	secpAddrCosmos := sdk.MustBech32ifyAddressBytes(sdk.Bech32MainPrefix, secpAddr)

	// ethsecp256k1 account
	ethPk, err := ethsecp256k1.GenerateKey()
	suite.Require().Nil(err)
	ethsecpAddr := sdk.AccAddress(ethPk.PubKey().Address())
	ethsecpAddrSacas := sdk.AccAddress(ethPk.PubKey().Address()).String()
	ethsecpAddrCosmos := sdk.MustBech32ifyAddressBytes(sdk.Bech32MainPrefix, ethsecpAddr)

	// Setup Cosmos <=> Sacas IBC relayer
	denom := "uatom"
	sourceChannel := "channel-292"
	sacasChannel := claimstypes.DefaultAuthorizedChannels[1]
	path := fmt.Sprintf("%s/%s", transfertypes.PortID, sacasChannel)

	timeoutHeight := clienttypes.NewHeight(0, 100)
	disabledTimeoutTimestamp := uint64(0)
	mockPacket := channeltypes.NewPacket(ibcgotesting.MockPacketData, 1, transfertypes.PortID, "channel-0", transfertypes.PortID, "channel-0", timeoutHeight, disabledTimeoutTimestamp)
	packet := mockPacket
	expAck := ibcmock.MockAcknowledgement

	coins := sdk.NewCoins(
		sdk.NewCoin("asacas", sdk.NewInt(1000)),
		sdk.NewCoin(ibcAtomDenom, sdk.NewInt(1000)),
		sdk.NewCoin(ibcOsmoDenom, sdk.NewInt(1000)),
		sdk.NewCoin(erc20Denom, sdk.NewInt(1000)),
	)

	testCases := []struct {
		name        string
		malleate    func()
		ackSuccess  bool
		expRecovery bool
		expCoins    sdk.Coins
	}{
		{
			"continue - params disabled",
			func() {
				params := suite.app.RecoveryKeeper.GetParams(suite.ctx)
				params.EnableRecovery = false
				suite.app.RecoveryKeeper.SetParams(suite.ctx, params) //nolint:errcheck
			},
			true,
			false,
			coins,
		},
		{
			"continue - destination channel not authorized",
			func() {
				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", ethsecpAddrSacas, ethsecpAddrCosmos, "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 1, transfertypes.PortID, sourceChannel, transfertypes.PortID, "channel-100", timeoutHeight, 0)
			},
			true,
			false,
			coins,
		},
		{
			"continue - destination channel is EVM",
			func() {
				EVMChannels := suite.app.ClaimsKeeper.GetParams(suite.ctx).EVMChannels
				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", ethsecpAddrSacas, ethsecpAddrCosmos, "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 1, transfertypes.PortID, sourceChannel, transfertypes.PortID, EVMChannels[0], timeoutHeight, 0)
			},
			true,
			false,
			coins,
		},
		{
			"fail - non ics20 packet",
			func() {
				packet = mockPacket
			},
			false,
			false,
			coins,
		},
		{
			"fail - invalid sender - missing '1' ",
			func() {
				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", "sacas", ethsecpAddrCosmos, "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)
			},
			false,
			false,
			coins,
		},
		{
			"fail - invalid sender - invalid bech32",
			func() {
				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", "badba1sv9m0g7ycejwr3s369km58h5qe7xj77hvcxrms", ethsecpAddrCosmos, "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)
			},
			false,
			false,
			coins,
		},
		{
			"fail - invalid recipient",
			func() {
				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", ethsecpAddrSacas, "badbadhf0468jjpe6m6vx38s97z2qqe8ldu0njdyf625", "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)
			},
			false,
			false,
			coins,
		},
		{
			"fail - case: receiver address is in deny list",
			func() {
				blockedAddr := authtypes.NewModuleAddress(transfertypes.ModuleName)

				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", secpAddrCosmos, blockedAddr.String(), "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)
			},
			false,
			false,
			coins,
		},
		{
			"continue - sender != receiver",
			func() {
				pk1 := secp256k1.GenPrivKey()
				otherSecpAddrSacas := sdk.AccAddress(pk1.PubKey().Address()).String()

				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", secpAddrCosmos, otherSecpAddrSacas, "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)
			},
			true,
			false,
			coins,
		},
		{
			"continue - receiver is a vesting account",
			func() {
				// Set vesting account
				bacc := authtypes.NewBaseAccount(ethsecpAddr, nil, 0, 0)
				acc := vestingtypes.NewClawbackVestingAccount(bacc, ethsecpAddr, nil, suite.ctx.BlockTime(), nil, nil)

				suite.app.AccountKeeper.SetAccount(suite.ctx, acc)

				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", ethsecpAddrCosmos, ethsecpAddrSacas, "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)
			},
			true,
			false,
			coins,
		},
		{
			"continue - receiver is a module account",
			func() {
				incentivesAcc := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, incentivestypes.ModuleName)
				suite.Require().NotNil(incentivesAcc)
				addr := incentivesAcc.GetAddress().String()
				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", addr, addr, "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)
			},
			true,
			false,
			coins,
		},
		{
			"continue - receiver pubkey is a supported key",
			func() {
				// Set account to generate a pubkey
				suite.app.AccountKeeper.SetAccount(suite.ctx, authtypes.NewBaseAccount(ethsecpAddr, ethPk.PubKey(), 0, 0))

				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", ethsecpAddrCosmos, ethsecpAddrSacas, "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)
			},
			true,
			false,
			coins,
		},
		{
			"partial recovery - account has invalid ibc vouchers balance",
			func() {
				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", secpAddrCosmos, secpAddrSacas, "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)

				invalidDenom := "ibc/1"
				coins := sdk.NewCoins(sdk.NewCoin(invalidDenom, sdk.NewInt(1000)))
				err := testutil.FundAccount(suite.ctx, suite.app.BankKeeper, secpAddr, coins)
				suite.Require().NoError(err)
			},
			false,
			false,
			sdk.NewCoins(
				sdk.NewCoin("ibc/1", sdk.NewInt(1000)),
				sdk.NewCoin(ibcAtomDenom, sdk.NewInt(1000)),
				sdk.NewCoin(ibcOsmoDenom, sdk.NewInt(1000)),
			),
		},
		{
			"recovery - send uatom from cosmos to sacas",
			func() {
				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", secpAddrCosmos, secpAddrSacas, "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)
			},
			true,
			true,
			nil,
		},
		{
			"recovery - send ibc/uosmo from cosmos to sacas",
			func() {
				denom = ibcOsmoDenom

				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", secpAddrCosmos, secpAddrSacas, "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)
			},
			true,
			true,
			nil,
		},
		{
			"recovery - send uosmo from osmosis to sacas",
			func() {
				// Setup Osmosis <=> Sacas IBC relayer
				denom = "uosmo"
				sourceChannel = "channel-204"
				sacasChannel = claimstypes.DefaultAuthorizedChannels[0]
				path = fmt.Sprintf("%s/%s", transfertypes.PortID, sacasChannel)

				transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", secpAddrCosmos, secpAddrSacas, "")
				bz := transfertypes.ModuleCdc.MustMarshalJSON(&transfer)
				packet = channeltypes.NewPacket(bz, 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)
				// TODO TEST
			},
			true,
			true,
			nil,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			// Enable Recovery
			params := suite.app.RecoveryKeeper.GetParams(suite.ctx)
			params.EnableRecovery = true
			err := suite.app.RecoveryKeeper.SetParams(suite.ctx, params)
			suite.Require().NoError(err)

			tc.malleate()

			// Set Denom Trace
			denomTrace := transfertypes.DenomTrace{
				Path:      path,
				BaseDenom: denom,
			}
			suite.app.TransferKeeper.SetDenomTrace(suite.ctx, denomTrace)

			// Set Cosmos Channel
			channel := channeltypes.Channel{
				State:          channeltypes.INIT,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   channeltypes.NewCounterparty(transfertypes.PortID, sourceChannel),
				ConnectionHops: []string{sourceChannel},
			}
			suite.app.IBCKeeper.ChannelKeeper.SetChannel(suite.ctx, transfertypes.PortID, sacasChannel, channel)

			// Set Next Sequence Send
			suite.app.IBCKeeper.ChannelKeeper.SetNextSequenceSend(suite.ctx, transfertypes.PortID, sacasChannel, 1)

			// Mock the Transferkeeper to always return nil on SendTransfer(), as this
			// method requires a successful handshake with the counterparty chain.
			// This, however, exceeds the requirements of the unit tests.
			mockTransferKeeper := &MockTransferKeeper{
				Keeper: suite.app.BankKeeper,
			}

			mockTransferKeeper.On("GetDenomTrace", mock.Anything, mock.Anything).Return(denomTrace, true)
			mockTransferKeeper.On("Transfer", mock.Anything, mock.Anything).Return(nil, nil)

			suite.app.RecoveryKeeper = keeper.NewKeeper(
				suite.app.GetKey(types.StoreKey),
				suite.app.AppCodec(),
				authtypes.NewModuleAddress(govtypes.ModuleName),
				suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.IBCKeeper.ChannelKeeper, mockTransferKeeper, suite.app.ClaimsKeeper)

			// Fund receiver account with SACAS, ERC20 coins and IBC vouchers
			err = testutil.FundAccount(suite.ctx, suite.app.BankKeeper, secpAddr, coins)
			suite.Require().NoError(err)

			// Perform IBC callback
			ack := suite.app.RecoveryKeeper.OnRecvPacket(suite.ctx, packet, expAck)

			// Check acknowledgement
			if tc.ackSuccess {
				suite.Require().True(ack.Success(), string(ack.Acknowledgement()))
				suite.Require().Equal(expAck, ack)
			} else {
				suite.Require().False(ack.Success(), string(ack.Acknowledgement()))
			}

			// Check recovery
			balances := suite.app.BankKeeper.GetAllBalances(suite.ctx, secpAddr)
			if tc.expRecovery {
				suite.Require().True(balances.IsZero())
			} else {
				suite.Require().Equal(tc.expCoins, balances)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGetIBCDenomDestinationIdentifiers() {
	address := sdk.AccAddress(utiltx.GenerateAddress().Bytes()).String()

	testCases := []struct {
		name                                      string
		denom                                     string
		malleate                                  func()
		expError                                  bool
		expDestinationPort, expDestinationChannel string
	}{
		{
			"invalid native denom",
			"asacas",
			func() {},
			true,
			"", "",
		},
		{
			"invalid IBC denom hash",
			"ibc/asacas",
			func() {},
			true,
			"", "",
		},
		{
			"denom trace not found",
			ibcAtomDenom,
			func() {},
			true,
			"", "",
		},
		{
			"channel not found",
			ibcAtomDenom,
			func() {
				denomTrace := transfertypes.DenomTrace{
					Path:      "transfer/channel-3",
					BaseDenom: "uatom",
				}
				suite.app.TransferKeeper.SetDenomTrace(suite.ctx, denomTrace)
			},
			true,
			"", "",
		},
		{
			"invalid destination port - insufficient length",
			"ibc/B9A49AA0AB0EB977D4EC627D7D9F747AF11BB1D74F430DE759CA37B22ECACF30", // denomTrace.Hash()
			func() {
				denomTrace := transfertypes.DenomTrace{
					Path:      "t/channel-3",
					BaseDenom: "uatom",
				}
				suite.app.TransferKeeper.SetDenomTrace(suite.ctx, denomTrace)

				channel := channeltypes.Channel{
					Counterparty: channeltypes.NewCounterparty("t", "channel-292"),
				}
				suite.app.IBCKeeper.ChannelKeeper.SetChannel(suite.ctx, "t", "channel-3", channel)
			},
			true,
			"", "",
		},
		{
			"invalid channel identifier - insufficient length",
			"ibc/5E3E083402F07599C795A7B75058EC3F13A8E666A8FEA2E51B6F3D93C755DFBC", // denomTrace.Hash()
			func() {
				denomTrace := transfertypes.DenomTrace{
					Path:      "transfer/c-3",
					BaseDenom: "uatom",
				}
				suite.app.TransferKeeper.SetDenomTrace(suite.ctx, denomTrace)

				channel := channeltypes.Channel{
					Counterparty: channeltypes.NewCounterparty("transfer", "channel-292"),
				}
				suite.app.IBCKeeper.ChannelKeeper.SetChannel(suite.ctx, "transfer", "c-3", channel)
			},
			true,
			"", "",
		},
		{
			"success - ATOM",
			ibcAtomDenom,
			func() {
				denomTrace := transfertypes.DenomTrace{
					Path:      "transfer/channel-3",
					BaseDenom: "uatom",
				}
				suite.app.TransferKeeper.SetDenomTrace(suite.ctx, denomTrace)

				channel := channeltypes.Channel{
					Counterparty: channeltypes.NewCounterparty("transfer", "channel-292"),
				}
				suite.app.IBCKeeper.ChannelKeeper.SetChannel(suite.ctx, "transfer", "channel-3", channel)
			},
			false,
			"transfer", "channel-3",
		},
		{
			"success - OSMO",
			ibcOsmoDenom,
			func() {
				denomTrace := transfertypes.DenomTrace{
					Path:      "transfer/channel-0",
					BaseDenom: "uosmo",
				}
				suite.app.TransferKeeper.SetDenomTrace(suite.ctx, denomTrace)

				channel := channeltypes.Channel{
					Counterparty: channeltypes.NewCounterparty("transfer", "channel-204"),
				}
				suite.app.IBCKeeper.ChannelKeeper.SetChannel(suite.ctx, "transfer", "channel-0", channel)
			},
			false,
			"transfer", "channel-0",
		},
		{
			"success - ibcATOM (via Osmosis)",
			"ibc/6CDD4663F2F09CD62285E2D45891FC149A3568E316CE3EBBE201A71A78A69388",
			func() {
				denomTrace := transfertypes.DenomTrace{
					Path:      "transfer/channel-0/transfer/channel-0",
					BaseDenom: "uatom",
				}

				suite.app.TransferKeeper.SetDenomTrace(suite.ctx, denomTrace)

				channel := channeltypes.Channel{
					Counterparty: channeltypes.NewCounterparty("transfer", "channel-204"),
				}
				suite.app.IBCKeeper.ChannelKeeper.SetChannel(suite.ctx, "transfer", "channel-0", channel)
			},
			false,
			"transfer", "channel-0",
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			tc.malleate()

			destinationPort, destinationChannel, err := suite.app.RecoveryKeeper.GetIBCDenomDestinationIdentifiers(suite.ctx, tc.denom, address)
			if tc.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expDestinationPort, destinationPort)
				suite.Require().Equal(tc.expDestinationChannel, destinationChannel)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnRecvPacketFailTransfer() {
	// secp256k1 account
	secpPk := secp256k1.GenPrivKey()
	secpAddr := sdk.AccAddress(secpPk.PubKey().Address())
	secpAddrSacas := secpAddr.String()
	secpAddrCosmos := sdk.MustBech32ifyAddressBytes(sdk.Bech32MainPrefix, secpAddr)

	// Setup Cosmos <=> Sacas IBC relayer
	denom := "uatom"
	sourceChannel := "channel-292"
	sacasChannel := claimstypes.DefaultAuthorizedChannels[1]
	path := fmt.Sprintf("%s/%s", transfertypes.PortID, sacasChannel)

	var mockTransferKeeper *MockTransferKeeper
	expAck := ibcmock.MockAcknowledgement
	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"Fail to retrieve ibc denom trace",
			func() {
				mockTransferKeeper.On("GetDenomTrace", mock.Anything, mock.Anything).Return(transfertypes.DenomTrace{}, false)
				mockTransferKeeper.On("Transfer", mock.Anything, mock.Anything).Return(nil, nil)
			},
		},
		{
			"invalid ibc denom trace",
			func() {
				// Set Denom Trace
				denomTrace := transfertypes.DenomTrace{
					Path:      "badpath",
					BaseDenom: denom,
				}
				suite.app.TransferKeeper.SetDenomTrace(suite.ctx, denomTrace)
				mockTransferKeeper.On("GetDenomTrace", mock.Anything, mock.Anything).Return(denomTrace, true)
				mockTransferKeeper.On("Transfer", mock.Anything, mock.Anything).Return(nil, nil)
			},
		},

		{
			"Fail to send transfer",
			func() {
				// Set Denom Trace
				denomTrace := transfertypes.DenomTrace{
					Path:      path,
					BaseDenom: denom,
				}
				suite.app.TransferKeeper.SetDenomTrace(suite.ctx, denomTrace)
				mockTransferKeeper.On("GetDenomTrace", mock.Anything, mock.Anything).Return(denomTrace, true)
				mockTransferKeeper.On("Transfer", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("Fail to transfer"))
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			// Enable Recovery
			params := suite.app.RecoveryKeeper.GetParams(suite.ctx)
			params.EnableRecovery = true
			err := suite.app.RecoveryKeeper.SetParams(suite.ctx, params)
			suite.Require().NoError(err)

			transfer := transfertypes.NewFungibleTokenPacketData(denom, "100", secpAddrCosmos, secpAddrSacas, "")
			packet := channeltypes.NewPacket(transfer.GetBytes(), 100, transfertypes.PortID, sourceChannel, transfertypes.PortID, sacasChannel, timeoutHeight, 0)

			mockTransferKeeper = &MockTransferKeeper{
				Keeper: suite.app.BankKeeper,
			}

			tc.malleate()

			suite.app.RecoveryKeeper = keeper.NewKeeper(
				suite.app.GetKey(types.StoreKey),
				suite.app.AppCodec(),
				authtypes.NewModuleAddress(govtypes.ModuleName),
				suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.IBCKeeper.ChannelKeeper, mockTransferKeeper, suite.app.ClaimsKeeper)

			// Fund receiver account with SACAS
			coins := sdk.NewCoins(
				sdk.NewCoin("asacas", sdk.NewInt(1000)),
				sdk.NewCoin(ibcAtomDenom, sdk.NewInt(1000)),
			)
			err = testutil.FundAccount(suite.ctx, suite.app.BankKeeper, secpAddr, coins)
			suite.Require().NoError(err)

			// Perform IBC callback
			ack := suite.app.RecoveryKeeper.OnRecvPacket(suite.ctx, packet, expAck)
			// Recovery should Fail
			suite.Require().False(ack.Success())
		})
	}
}
