// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package governance

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmostestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	governancekeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtestutil "github.com/cosmos/cosmos-sdk/x/gov/testutil"
	governancetypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	"pkg.berachain.dev/polaris/lib/utils"
	polarisprecompile "pkg.berachain.dev/polaris/precompile"
	"pkg.berachain.dev/polaris/precompile/contracts/solidity/generated"
	testutil "pkg.berachain.dev/polaris/testing/utils"
	evmutils "pkg.berachain.dev/polaris/x/evm/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Test Reporter to use governance module tests with Ginkgo.
type GinkgoTestReporter struct{}

func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

func fundAccount(ctx sdk.Context, bk bankkeeper.Keeper, acc sdk.AccAddress, coins sdk.Coins) {
	if err := bk.MintCoins(ctx, governancetypes.ModuleName, coins); err != nil {
		panic(err)
	}
	if err := bk.SendCoinsFromModuleToAccount(ctx, governancetypes.ModuleName, acc, coins); err != nil {
		panic(err)
	}
}

// Helper functions for setting up the tests.
func setup(ctrl *gomock.Controller, caller sdk.AccAddress) (
	sdk.Context,
	bankkeeper.Keeper,
	*governancekeeper.Keeper,
) {
	// Setup the keepers and context.
	ctx, ak, bk, sk := testutil.SetupMinimalKeepers()
	dk := govtestutil.NewMockDistributionKeeper(ctrl)

	// Register the governance module account.
	ak.SetModuleAccount(
		ctx,
		authtypes.NewEmptyModuleAccount(governancetypes.ModuleName, authtypes.Minter),
	)

	// Create the codec.
	encCfg := cosmostestutil.MakeTestEncodingConfig(
		gov.AppModuleBasic{},
		bank.AppModuleBasic{},
	)

	// Create the base app msgRouter.
	msr := baseapp.NewMsgServiceRouter()

	// Create the governance keeper.
	gk := governancekeeper.NewKeeper(
		encCfg.Codec,
		testutil.EvmKey,
		ak,
		bk,
		sk,
		dk,
		msr,
		governancetypes.DefaultConfig(),
		authtypes.NewModuleAddress(governancetypes.ModuleName).String(),
	)

	// Register the msg Service Handlers.
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	v1.RegisterMsgServer(msr, governancekeeper.NewMsgServerImpl(gk))
	banktypes.RegisterMsgServer(msr, bankkeeper.NewMsgServerImpl(bk))

	// Set the Params and first proposal ID.
	params := v1.DefaultParams()
	err := gk.SetParams(ctx, params)
	if err != nil {
		panic(err)
	}
	gk.SetProposalID(ctx, 1)

	// Fund the caller with some coins.
	fundAccount(ctx, bk, caller, sdk.NewCoins(sdk.NewInt64Coin("usdc", 100000000)))

	return ctx, bk, gk
}

func TestGovernancePrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "precompile/governance")
}

var _ = Describe("Governance Precompile", func() {
	var (
		ctx      sdk.Context
		bk       bankkeeper.Keeper
		gk       *governancekeeper.Keeper
		caller   sdk.AccAddress
		mockCtrl *gomock.Controller
		contract *Contract
	)

	BeforeEach(func() {
		t := GinkgoTestReporter{}
		mockCtrl = gomock.NewController(t)
		caller = evmutils.AddressToAccAddress(testutil.Alice)
		ctx, bk, gk = setup(mockCtrl, caller)
		contract = utils.MustGetAs[*Contract](NewContract(&gk))
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	When("Submitting a proposal", func() {
		It("should fail if the message is not of type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				"invalid",
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidAny))
			Expect(res).To(BeNil())
		})
		It("should fail if the initial deposit is wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				"invalid",
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidCoin))
			Expect(res).To(BeNil())
		})
		It("should fail if metadata is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				[]generated.IGovernanceModuleCoin{},
				123,
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})
		It("should fail if title is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				[]generated.IGovernanceModuleCoin{},
				"metadata",
				123,
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})
		It("should fail if summary is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				[]generated.IGovernanceModuleCoin{},
				"metadata",
				"title",
				123,
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})
		It("should fail if expadited is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				[]generated.IGovernanceModuleCoin{},
				"metadata",
				"title",
				"summary",
				123,
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidBool))
			Expect(res).To(BeNil())
		})
		It("should succeed", func() {
			initDeposit := sdk.NewCoins(sdk.NewInt64Coin("usdc", 100))
			govAcct := gk.GetGovernanceAccount(ctx).GetAddress()
			fundAccount(ctx, bk, govAcct, initDeposit)
			message := &banktypes.MsgSend{
				FromAddress: govAcct.String(),
				ToAddress:   caller.String(),
				Amount:      initDeposit,
			}

			metadata := "metadata"
			title := "title"
			summary := "summary"

			msg, err := codectypes.NewAnyWithValue(message)
			Expect(err).ToNot(HaveOccurred())

			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{msg},
				[]generated.IGovernanceModuleCoin{
					{
						Amount: big.NewInt(100),
						Denom:  "usdc",
					},
				},
				metadata,
				title,
				summary,
				false,
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})
	})

	When("Canceling a proposal", func() {
		It("should fail if the proposal ID is invalid", func() {
			res, err := contract.CancelProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				"invalid",
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidBigInt))
			Expect(res).To(BeNil())
		})
		It("should fail if the proposal does not exist", func() {
			res, err := contract.CancelProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				big.NewInt(1),
			)
			Expect(err).To(HaveOccurred())
			Expect(res).To(BeNil())
		})
		It("should succeed", func() {
			gk.SetProposal(ctx, v1.Proposal{
				Id:       1,
				Proposer: caller.String(),
				Messages: []*codectypes.Any{},
				Status:   v1.StatusVotingPeriod,
			})
			res, err := contract.CancelProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				big.NewInt(1),
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})

	})

	When("Voting on a proposal", func() {
		BeforeEach(func() {
			gk.SetProposal(ctx, v1.Proposal{
				Id:       1,
				Proposer: caller.String(),
				Messages: []*codectypes.Any{},
				Status:   v1.StatusVotingPeriod,
			})
		})
		It("should fail if the proposal ID is of invalid type", func() {
			res, err := contract.Vote(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				"invalid",
				int32(1),
				"metadata",
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidBigInt))
			Expect(res).To(BeNil())
		})
		It("should fail if the vote option is of invalid type", func() {
			res, err := contract.Vote(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				big.NewInt(1),
				"invalid",
				"metadata",
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidInt32))
			Expect(res).To(BeNil())
		})
		It("should fail if the metadata is of invalid type", func() {
			res, err := contract.Vote(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				big.NewInt(1),
				int32(1),
				123,
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})
		It("should fail if the proposal does not exist", func() {
			res, err := contract.Vote(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				big.NewInt(100),
				int32(1),
				"metadata",
			)
			Expect(err).To(HaveOccurred())
			Expect(res).To(BeNil())
		})
		It("should succeed", func() {
			res, err := contract.Vote(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				big.NewInt(1),
				int32(1),
				"metadata",
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})

		When("Voting Weight", func() {
			It("should fail if the proposal ID is of invalid type", func() {
				res, err := contract.VoteWeighted(
					ctx,
					evmutils.AccAddressToEthAddress(caller),
					big.NewInt(0),
					false,
					"invalid",
					[]generated.IGovernanceModuleWeightedVoteOption{},
					"metadata",
				)
				Expect(err).To(MatchError(polarisprecompile.ErrInvalidBigInt))
				Expect(res).To(BeNil())
			})
			It("should fail if the vote option is of invalid type", func() {
				res, err := contract.VoteWeighted(
					ctx,
					evmutils.AccAddressToEthAddress(caller),
					big.NewInt(0),
					false,
					big.NewInt(1),
					12,
					"metadata",
				)
				Expect(err).To(MatchError(polarisprecompile.ErrInvalidOptions))
				Expect(res).To(BeNil())
			})
			It("should fail if the metadata is of invalid type", func() {
				res, err := contract.VoteWeighted(
					ctx,
					evmutils.AccAddressToEthAddress(caller),
					big.NewInt(0),
					false,
					big.NewInt(1),
					[]generated.IGovernanceModuleWeightedVoteOption{},
					123,
				)
				Expect(err).To(MatchError(polarisprecompile.ErrInvalidString))
				Expect(res).To(BeNil())
			})
			It("should fail if the proposal does not exist", func() {
				res, err := contract.VoteWeighted(
					ctx,
					evmutils.AccAddressToEthAddress(caller),
					big.NewInt(0),
					false,
					big.NewInt(100),
					[]generated.IGovernanceModuleWeightedVoteOption{},
					"metadata",
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})
			It("should succeed", func() {
				weight, err := math.LegacyNewDecFromStr("0.4")
				Expect(err).ToNot(HaveOccurred())
				options := []generated.IGovernanceModuleWeightedVoteOption{
					{
						VoteOption: int32(1),
						Weight:     weight.String(),
					},
				}
				res, err := contract.VoteWeighted(
					ctx,
					evmutils.AccAddressToEthAddress(caller),
					big.NewInt(0),
					false,
					big.NewInt(1),
					options,
					"metadata",
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).ToNot(BeNil())
			})
		})
		When("Reading Methods", func() {
			BeforeEach(func() {
				gk.SetProposal(ctx, v1.Proposal{
					Id:               2,
					Proposer:         caller.String(),
					Messages:         []*codectypes.Any{},
					Status:           v1.StatusVotingPeriod,
					FinalTallyResult: &v1.TallyResult{},
					SubmitTime:       &time.Time{},
					DepositEndTime:   &time.Time{},
					TotalDeposit:     sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))),
					VotingStartTime:  &time.Time{},
					VotingEndTime:    &time.Time{},
					Metadata:         "metadata",
					Title:            "title",
					Summary:          "summary",
					Expedited:        false,
				})
				gk.SetProposal(ctx, v1.Proposal{
					Id:               3,
					Proposer:         caller.String(),
					Messages:         []*codectypes.Any{},
					Status:           v1.StatusVotingPeriod,
					FinalTallyResult: &v1.TallyResult{},
					SubmitTime:       &time.Time{},
					DepositEndTime:   &time.Time{},
					TotalDeposit:     sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))),
					VotingStartTime:  &time.Time{},
					VotingEndTime:    &time.Time{},
					Metadata:         "metadata",
					Title:            "title",
					Summary:          "summary",
					Expedited:        false,
				})
			})
			When("GetProposal", func() {
				It("should fail if the proposal ID is of invalid type", func() {
					res, err := contract.GetProposal(
						ctx,
						evmutils.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						"invalid",
					)
					Expect(err).To(MatchError(polarisprecompile.ErrInvalidBigInt))
					Expect(res).To(BeNil())
				})
				It("should get the proposal", func() {
					res, err := contract.GetProposal(
						ctx,
						evmutils.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						big.NewInt(2),
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
					Expect(res).To(HaveLen(1))
				})
			})
			When("GetProposalsStringAddr", func() {
				It("should fail if the proposal status is of invalid type", func() {
					res, err := contract.GetProposalStringAddr(
						ctx,
						evmutils.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						"invalid",
					)
					Expect(err).To(MatchError(polarisprecompile.ErrInvalidInt32))
					Expect(res).To(BeNil())
				})
				It("should fail if the voter string is not string", func() {
					res, err := contract.GetProposalStringAddr(
						ctx,
						evmutils.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						int32(0),
						123,
					)
					Expect(err).To(MatchError(polarisprecompile.ErrInvalidString))
					Expect(res).To(BeNil())
				})
				It("should fail if the voter address is not string", func() {
					res, err := contract.GetProposalStringAddr(
						ctx,
						evmutils.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						int32(0),
						"addr",
						123,
					)
					Expect(err).To(MatchError(polarisprecompile.ErrInvalidString))
					Expect(res).To(BeNil())
				})
				It("should fail if the voter address is not valid bech32", func() {
					res, err := contract.GetProposalStringAddr(
						ctx,
						evmutils.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						int32(0),
						"addr",
						"addr",
					)
					Expect(err).To(HaveOccurred())
					Expect(res).To(BeNil())
				})
				It("should fail if the depositor address is not bech32", func() {
					res, err := contract.GetProposalStringAddr(
						ctx,
						evmutils.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						int32(0),
						caller.String(),
						"addr",
					)
					Expect(err).To(HaveOccurred())
					Expect(res).To(BeNil())
				})
				It("should get the proposals", func() {
					res, err := contract.GetProposalStringAddr(
						ctx,
						evmutils.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						int32(0),
						caller.String(),
						caller.String(),
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
				})
			})
			When("GetProposalsAddr", func() {
				It("should fail if the proposal status is of invalid type", func() {
					res, err := contract.GetProposalsAddr(
						ctx,
						evmutils.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						"invalid",
					)
					Expect(err).To(MatchError(polarisprecompile.ErrInvalidInt32))
					Expect(res).To(BeNil())
				})
				It("should fail if the voter address is not common.Address", func() {
					res, err := contract.GetProposalsAddr(
						ctx,
						evmutils.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						int32(0),
						"addr",
					)
					Expect(err).To(MatchError(polarisprecompile.ErrInvalidHexAddress))
					Expect(res).To(BeNil())
				})
				It("should fail if the depositor address is not common.Address", func() {
					res, err := contract.GetProposalsAddr(
						ctx,
						evmutils.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						int32(0),
						evmutils.AccAddressToEthAddress(caller),
						"addr",
					)
					Expect(err).To(MatchError(polarisprecompile.ErrInvalidHexAddress))
					Expect(res).To(BeNil())
				})
				It("should get the proposals", func() {
					res, err := contract.GetProposalsAddr(
						ctx,
						evmutils.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						int32(0),
						evmutils.AccAddressToEthAddress(caller),
						evmutils.AccAddressToEthAddress(caller),
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
				})
			})
		})
	})
})
