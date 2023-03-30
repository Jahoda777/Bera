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

package bank_test

import (
	"math/big"
	"testing"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"

	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/cosmos/precompile/bank"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile/log"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBankPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/precompile/bank")
}

var _ = Describe("Bank Precompile Test", func() {
	var (
		contract *bank.Contract
		addr     sdk.AccAddress
		factory  *log.Factory
		bk       bankkeeper.BaseKeeper
		ctx      sdk.Context
	)

	BeforeEach(func() {
		ctx, _, bk, _ = testutil.SetupMinimalKeepers()
		contract = utils.MustGetAs[*bank.Contract](bank.NewPrecompileContract(bk))
		addr = sdk.AccAddress([]byte("bank"))

		// Register the events.
		factory = log.NewFactory([]vm.RegistrablePrecompile{contract})
	})

	It("should register the send event", func() {
		event := sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdk.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).ToNot(HaveOccurred())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	It("should register the transfer event", func() {
		event := sdk.NewEvent(
			banktypes.EventTypeTransfer,
			sdk.NewAttribute(banktypes.AttributeKeyRecipient, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdk.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).ToNot(HaveOccurred())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	It("should register the coin spent event", func() {
		event := sdk.NewEvent(
			banktypes.EventTypeCoinSpent,
			sdk.NewAttribute(banktypes.AttributeKeySpender, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdk.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).ToNot(HaveOccurred())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	It("should register the coin received event", func() {
		event := sdk.NewEvent(
			banktypes.EventTypeCoinReceived,
			sdk.NewAttribute(banktypes.AttributeKeyReceiver, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdk.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).ToNot(HaveOccurred())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	It("should register the burn event", func() {
		event := sdk.NewEvent(
			banktypes.EventTypeCoinBurn,
			sdk.NewAttribute(banktypes.AttributeKeyBurner, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdk.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).ToNot(HaveOccurred())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	When("Calling Precompile Methods", func() {
		var (
			acc    sdk.AccAddress
			caller common.Address
		)

		denom := "bera"

		BeforeEach(func() {})

		When("GetBalance", func() {
			It("should fail if input is not a common.Address", func() {
				res, err := contract.GetBalance(
					ctx,
					nil,
					caller,
					big.NewInt(0),
					true,
					"0x",
				)
				Expect(err).To(MatchError(precompile.ErrInvalidHexAddress))
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				balanceAmount, ok := new(big.Int).SetString("22000000000000000000", 10)
				Expect(ok).To(BeTrue())

				acc = simtestutil.CreateRandomAccounts(1)[0]

				err := FundAccount(
					ctx,
					bk,
					acc,
					sdk.NewCoins(
						sdk.NewCoin(
							denom,
							sdk.NewIntFromBigInt(balanceAmount),
						),
					),
				)
				Expect(err).ToNot(HaveOccurred())

				res, err := contract.GetBalance(
					ctx,
					nil,
					caller,
					big.NewInt(0),
					true,
					cosmlib.AccAddressToEthAddress(acc),
					denom,
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res[0]).To(Equal(balanceAmount))
			})
		})
	})
})

func FundAccount(ctx sdk.Context, bk bankkeeper.BaseKeeper, account sdk.AccAddress, coins sdk.Coins) error {
	if err := bk.MintCoins(ctx, evmtypes.ModuleName, coins); err != nil {
		return err
	}
	return bk.SendCoinsFromModuleToAccount(ctx, evmtypes.ModuleName, account, coins)
}
