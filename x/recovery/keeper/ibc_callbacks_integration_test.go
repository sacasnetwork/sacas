package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	"github.com/sacasnetwork/sacas/v15/app"
	ibctesting "github.com/sacasnetwork/sacas/v15/ibc/testing"
	"github.com/sacasnetwork/sacas/v15/testutil"
	teststypes "github.com/sacasnetwork/sacas/v15/types/tests"
	claimstypes "github.com/sacasnetwork/sacas/v15/x/claims/types"
	"github.com/sacasnetwork/sacas/v15/x/recovery/types"

	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/ginkgo/v2"
	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/gomega"
)

var _ = Describe("Recovery: Performing an IBC Transfer", Ordered, func() {
	coinSacas := sdk.NewCoin("asacas", sdk.NewInt(10000))
	coinOsmo := sdk.NewCoin("uosmo", sdk.NewInt(10))
	coinAtom := sdk.NewCoin("uatom", sdk.NewInt(10))

	var (
		sender, receiver       string
		senderAcc, receiverAcc sdk.AccAddress
		timeout                uint64
		claim                  claimstypes.ClaimsRecord
	)

	BeforeEach(func() {
		s.SetupTest()
	})

	Describe("from a non-authorized chain", func() {
		BeforeEach(func() {
			params := claimstypes.DefaultParams()
			params.AuthorizedChannels = []string{}
			err := s.SacasChain.App.(*app.Sacas).ClaimsKeeper.SetParams(s.SacasChain.GetContext(), params)
			s.Require().NoError(err)

			sender = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
			receiver = s.SacasChain.SenderAccount.GetAddress().String()
			senderAcc = sdk.MustAccAddressFromBech32(sender)
			receiverAcc = sdk.MustAccAddressFromBech32(receiver)
		})
		It("should transfer and not recover tokens", func() {
			s.SendAndReceiveMessage(s.pathOsmosisSacas, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 1)

			nativeSacas := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), senderAcc, "asacas")
			Expect(nativeSacas).To(Equal(coinSacas))
			ibcOsmo := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
			Expect(ibcOsmo).To(Equal(sdk.NewCoin(teststypes.UosmoIbcdenom, coinOsmo.Amount)))
		})
	})

	Describe("from an authorized, non-EVM chain (e.g. Osmosis)", func() {
		Describe("to a different account on Sacas (sender != recipient)", func() {
			BeforeEach(func() {
				sender = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
				receiver = s.SacasChain.SenderAccount.GetAddress().String()
				senderAcc = sdk.MustAccAddressFromBech32(sender)
				receiverAcc = sdk.MustAccAddressFromBech32(receiver)
			})

			It("should transfer and not recover tokens", func() {
				s.SendAndReceiveMessage(s.pathOsmosisSacas, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 1)

				nativeSacas := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), senderAcc, "asacas")
				Expect(nativeSacas).To(Equal(coinSacas))
				ibcOsmo := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
				Expect(ibcOsmo).To(Equal(sdk.NewCoin(teststypes.UosmoIbcdenom, coinOsmo.Amount)))
			})
		})

		Describe("to the sender's own eth_secp256k1 account on Sacas (sender == recipient)", func() {
			BeforeEach(func() {
				sender = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
				receiver = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
				senderAcc = sdk.MustAccAddressFromBech32(sender)
				receiverAcc = sdk.MustAccAddressFromBech32(receiver)
			})

			Context("with disabled recovery parameter", func() {
				BeforeEach(func() {
					params := types.DefaultParams()
					params.EnableRecovery = false
					s.SacasChain.App.(*app.Sacas).RecoveryKeeper.SetParams(s.SacasChain.GetContext(), params) //nolint:errcheck
				})

				It("should not transfer or recover tokens", func() {
					s.SendAndReceiveMessage(s.pathOsmosisSacas, s.IBCOsmosisChain, coinOsmo.Denom, coinOsmo.Amount.Int64(), sender, receiver, 1)
					nativeSacas := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), senderAcc, "asacas")
					Expect(nativeSacas).To(Equal(coinSacas))
					ibcOsmo := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
					Expect(ibcOsmo).To(Equal(sdk.NewCoin(teststypes.UosmoIbcdenom, coinOsmo.Amount)))
				})
			})

			Context("with a sender's claims record", func() {
				Context("without completed actions", func() {
					BeforeEach(func() {
						amt := sdk.NewInt(int64(100))
						coins := sdk.NewCoins(sdk.NewCoin("asacas", amt))
						claim = claimstypes.NewClaimsRecord(amt)
						s.SacasChain.App.(*app.Sacas).ClaimsKeeper.SetClaimsRecord(s.SacasChain.GetContext(), senderAcc, claim)

						// update the escrowed account balance to maintain the invariant
						err := testutil.FundModuleAccount(s.SacasChain.GetContext(), s.SacasChain.App.(*app.Sacas).BankKeeper, claimstypes.ModuleName, coins)
						s.Require().NoError(err)
					})

					It("should not transfer or recover tokens", func() {
						// Prevent further funds from getting stuck
						s.SendAndReceiveMessage(s.pathOsmosisSacas, s.IBCOsmosisChain, coinOsmo.Denom, coinOsmo.Amount.Int64(), sender, receiver, 1)
						nativeSacas := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), senderAcc, "asacas")
						Expect(nativeSacas).To(Equal(coinSacas))
						ibcOsmo := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
					})
				})

				Context("with completed actions", func() {
					// Already has stuck funds
					BeforeEach(func() {
						amt := sdk.NewInt(int64(100))
						coins := sdk.NewCoins(sdk.NewCoin("asacas", sdk.NewInt(int64(75))))
						claim = claimstypes.NewClaimsRecord(amt)
						claim.MarkClaimed(claimstypes.ActionIBCTransfer)
						s.SacasChain.App.(*app.Sacas).ClaimsKeeper.SetClaimsRecord(s.SacasChain.GetContext(), senderAcc, claim)

						// update the escrowed account balance to maintain the invariant
						err := testutil.FundModuleAccount(s.SacasChain.GetContext(), s.SacasChain.App.(*app.Sacas).BankKeeper, claimstypes.ModuleName, coins)
						s.Require().NoError(err)

						// asacas & ibc tokens that originated from the sender's chain
						s.SendAndReceiveMessage(s.pathOsmosisSacas, s.IBCOsmosisChain, coinOsmo.Denom, coinOsmo.Amount.Int64(), sender, receiver, 1)
						timeout = uint64(s.SacasChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())
					})

					It("should transfer tokens to the recipient and perform recovery", func() {
						// Escrow before relaying packets
						balanceEscrow := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), transfertypes.GetEscrowAddress("transfer", "channel-0"), "asacas")
						Expect(balanceEscrow).To(Equal(coinSacas))
						ibcOsmo := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())

						// Relay both packets that were sent in the ibc_callback
						err := s.pathOsmosisSacas.RelayPacket(CreatePacket("10000", "asacas", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisSacas.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)

						// Check that the asacas were recovered
						nativeSacas := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), senderAcc, "asacas")
						Expect(nativeSacas.IsZero()).To(BeTrue())
						ibcSacas := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, teststypes.AsacasIbcdenom)
						Expect(ibcSacas).To(Equal(sdk.NewCoin(teststypes.AsacasIbcdenom, coinSacas.Amount)))

						// Check that the uosmo were recovered
						ibcOsmo = s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))
					})

					It("should not claim/migrate/merge claims records", func() {
						// Relay both packets that were sent in the ibc_callback
						err := s.pathOsmosisSacas.RelayPacket(CreatePacket("10000", "asacas", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisSacas.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)

						claimAfter, _ := s.SacasChain.App.(*app.Sacas).ClaimsKeeper.GetClaimsRecord(s.SacasChain.GetContext(), senderAcc)
						Expect(claim).To(Equal(claimAfter))
					})
				})
			})

			Context("without a sender's claims record", func() {
				When("recipient has no ibc vouchers that originated from other chains", func() {
					It("should transfer and recover tokens", func() {
						// asacas & ibc tokens that originated from the sender's chain
						s.SendAndReceiveMessage(s.pathOsmosisSacas, s.IBCOsmosisChain, coinOsmo.Denom, coinOsmo.Amount.Int64(), sender, receiver, 1)
						timeout = uint64(s.SacasChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())

						// Escrow before relaying packets
						balanceEscrow := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), transfertypes.GetEscrowAddress("transfer", "channel-0"), "asacas")
						Expect(balanceEscrow).To(Equal(coinSacas))
						ibcOsmo := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())

						// Relay both packets that were sent in the ibc_callback
						err := s.pathOsmosisSacas.RelayPacket(CreatePacket("10000", "asacas", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisSacas.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)

						// Check that the asacas were recovered
						nativeSacas := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), senderAcc, "asacas")
						Expect(nativeSacas.IsZero()).To(BeTrue())
						ibcSacas := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, teststypes.AsacasIbcdenom)
						Expect(ibcSacas).To(Equal(sdk.NewCoin(teststypes.AsacasIbcdenom, coinSacas.Amount)))

						// Check that the uosmo were recovered
						ibcOsmo = s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))
					})
				})

				// Do not recover uatom sent from Cosmos when performing recovery through IBC transfer from Osmosis
				When("recipient has additional ibc vouchers that originated from other chains", func() {
					BeforeEach(func() {
						params := types.DefaultParams()
						params.EnableRecovery = false
						s.SacasChain.App.(*app.Sacas).RecoveryKeeper.SetParams(s.SacasChain.GetContext(), params) //nolint:errcheck

						// Send uatom from Cosmos to Sacas
						s.SendAndReceiveMessage(s.pathCosmosSacas, s.IBCCosmosChain, coinAtom.Denom, coinAtom.Amount.Int64(), s.IBCCosmosChain.SenderAccount.GetAddress().String(), receiver, 1)

						params.EnableRecovery = true
						s.SacasChain.App.(*app.Sacas).RecoveryKeeper.SetParams(s.SacasChain.GetContext(), params) //nolint:errcheck
					})
					It("should not recover tokens that originated from other chains", func() {
						// Send uosmo from Osmosis to Sacas
						s.SendAndReceiveMessage(s.pathOsmosisSacas, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 1)

						// Relay both packets that were sent in the ibc_callback
						timeout := uint64(s.SacasChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())
						err := s.pathOsmosisSacas.RelayPacket(CreatePacket("10000", "asacas", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisSacas.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)

						// Asacas was recovered from user address
						nativeSacas := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), senderAcc, "asacas")
						Expect(nativeSacas.IsZero()).To(BeTrue())
						ibcSacas := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, teststypes.AsacasIbcdenom)
						Expect(ibcSacas).To(Equal(sdk.NewCoin(teststypes.AsacasIbcdenom, coinSacas.Amount)))

						// Check that the uosmo were retrieved
						ibcOsmo := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))

						// Check that the atoms were not retrieved
						ibcAtom := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), senderAcc, teststypes.UatomIbcdenom)
						Expect(ibcAtom).To(Equal(sdk.NewCoin(teststypes.UatomIbcdenom, coinAtom.Amount)))

						// Repeat transaction from Osmosis to Sacas
						s.SendAndReceiveMessage(s.pathOsmosisSacas, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 2)

						timeout = uint64(s.SacasChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())
						err = s.pathOsmosisSacas.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 3, timeout))
						s.Require().NoError(err)

						// No further tokens recovered
						nativeSacas = s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), senderAcc, "asacas")
						Expect(nativeSacas.IsZero()).To(BeTrue())
						ibcSacas = s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, teststypes.AsacasIbcdenom)
						Expect(ibcSacas).To(Equal(sdk.NewCoin(teststypes.AsacasIbcdenom, coinSacas.Amount)))

						ibcOsmo = s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo = s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))

						ibcAtom = s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), senderAcc, teststypes.UatomIbcdenom)
						Expect(ibcAtom).To(Equal(sdk.NewCoin(teststypes.UatomIbcdenom, coinAtom.Amount)))
					})
				})

				// Recover ibc/uatom that was sent from Osmosis back to Osmosis
				When("recipient has additional non-native ibc vouchers that originated from senders chains", func() {
					BeforeEach(func() {
						params := types.DefaultParams()
						params.EnableRecovery = false
						err := s.SacasChain.App.(*app.Sacas).RecoveryKeeper.SetParams(s.SacasChain.GetContext(), params)
						s.Require().NoError(err)
						s.SendAndReceiveMessage(s.pathOsmosisCosmos, s.IBCCosmosChain, coinAtom.Denom, coinAtom.Amount.Int64(), s.IBCCosmosChain.SenderAccount.GetAddress().String(), receiver, 1)

						// Send IBC transaction of 10 ibc/uatom
						transferMsg := transfertypes.NewMsgTransfer(s.pathOsmosisSacas.EndpointA.ChannelConfig.PortID, s.pathOsmosisSacas.EndpointA.ChannelID, sdk.NewCoin(teststypes.UatomIbcdenom, sdk.NewInt(10)), sender, receiver, timeoutHeight, 0, "")
						_, err = ibctesting.SendMsgs(s.IBCOsmosisChain, ibctesting.DefaultFeeAmt, transferMsg)
						s.Require().NoError(err) // message committed
						transfer := transfertypes.NewFungibleTokenPacketData("transfer/channel-1/uatom", "10", sender, receiver, "")
						packet := channeltypes.NewPacket(transfer.GetBytes(), 1, s.pathOsmosisSacas.EndpointA.ChannelConfig.PortID, s.pathOsmosisSacas.EndpointA.ChannelID, s.pathOsmosisSacas.EndpointB.ChannelConfig.PortID, s.pathOsmosisSacas.EndpointB.ChannelID, timeoutHeight, 0)
						// Receive message on the sacas side, and send ack
						err = s.pathOsmosisSacas.RelayPacket(packet)
						s.Require().NoError(err)

						// Check that the ibc/uatom are available
						osmoIBCAtom := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UatomOsmoIbcdenom)
						s.Require().Equal(osmoIBCAtom.Amount, coinAtom.Amount)

						params.EnableRecovery = true
						s.SacasChain.App.(*app.Sacas).RecoveryKeeper.SetParams(s.SacasChain.GetContext(), params) //nolint:errcheck
					})
					It("should not recover tokens that originated from other chains", func() {
						s.SendAndReceiveMessage(s.pathOsmosisSacas, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 2)

						// Relay packets that were sent in the ibc_callback
						timeout := uint64(s.SacasChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())
						err := s.pathOsmosisSacas.RelayPacket(CreatePacket("10000", "asacas", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisSacas.RelayPacket(CreatePacket("10", "transfer/channel-0/transfer/channel-1/uatom", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisSacas.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 3, timeout))
						s.Require().NoError(err)

						// Asacas was recovered from user address
						nativeSacas := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), senderAcc, "asacas")
						Expect(nativeSacas.IsZero()).To(BeTrue())
						ibcSacas := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, teststypes.AsacasIbcdenom)
						Expect(ibcSacas).To(Equal(sdk.NewCoin(teststypes.AsacasIbcdenom, coinSacas.Amount)))

						// Check that the uosmo were recovered
						ibcOsmo := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))

						// Check that the ibc/uatom were retrieved
						osmoIBCAtom := s.SacasChain.App.(*app.Sacas).BankKeeper.GetBalance(s.SacasChain.GetContext(), receiverAcc, teststypes.UatomOsmoIbcdenom)
						Expect(osmoIBCAtom.IsZero()).To(BeTrue())
						ibcAtom := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), senderAcc, teststypes.UatomIbcdenom)
						Expect(ibcAtom).To(Equal(sdk.NewCoin(teststypes.UatomIbcdenom, sdk.NewInt(10))))
					})
				})
			})
		})
	})
})
