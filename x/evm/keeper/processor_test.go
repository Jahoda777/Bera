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

package keeper_test

import (
	"math/big"
	"os"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"

	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"pkg.berachain.dev/stargazer/eth/accounts/abi"
	"pkg.berachain.dev/stargazer/eth/common"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/crypto"
	"pkg.berachain.dev/stargazer/eth/params"
	"pkg.berachain.dev/stargazer/eth/testutil/contracts/solidity/generated"
	"pkg.berachain.dev/stargazer/precompile"
	"pkg.berachain.dev/stargazer/testutil"
	"pkg.berachain.dev/stargazer/x/evm/keeper"
	"pkg.berachain.dev/stargazer/x/evm/plugins/state"
	evmmempool "pkg.berachain.dev/stargazer/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/stargazer/x/evm/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Processor", func() {
	var (
		k            *keeper.Keeper
		ak           state.AccountKeeper
		bk           state.BankKeeper
		sk           stakingkeeper.Keeper
		ctx          sdk.Context
		key, _       = crypto.GenerateEthKey()
		signer       = coretypes.LatestSignerForChainID(params.DefaultChainConfig.ChainID)
		legacyTxData *coretypes.LegacyTx
	)

	BeforeEach(func() {
		err := os.RemoveAll("tmp/berachain")
		Expect(err).ToNot(HaveOccurred())

		legacyTxData = &coretypes.LegacyTx{
			Nonce:    0,
			Gas:      10000000,
			Data:     []byte("abcdef"),
			GasPrice: big.NewInt(1),
		}

		// before chain, init genesis state
		ctx, ak, bk, sk = testutil.SetupMinimalKeepers()
		k = keeper.NewKeeper(
			storetypes.NewKVStoreKey("evm"),
			ak, bk,
			precompile.NewProvider(&sk),
			"authority",
			sims.NewAppOptionsWithFlagHome("tmp/berachain"),
			evmmempool.NewEthTxPoolFrom(sdkmempool.NewPriorityMempool()),
		)
		for _, plugin := range k.GetAllPlugins() {
			plugin.InitGenesis(ctx, types.DefaultGenesis())
		}

		// before every block
		ctx = ctx.WithBlockGasMeter(storetypes.NewGasMeter(100000000)).
			WithKVGasConfig(storetypes.GasConfig{}).
			WithBlockHeight(1)
		k.BeginBlocker(ctx)
	})

	Context("New Block", func() {
		BeforeEach(func() {
			// before every tx
			ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())
		})

		AfterEach(func() {
			k.EndBlocker(ctx)
			err := os.RemoveAll("tmp/berachain")
			Expect(err).ToNot(HaveOccurred())
		})

		It("should panic on nil, empty transaction", func() {
			Expect(func() {
				_, err := k.ProcessTransaction(ctx, nil)
				Expect(err).To(HaveOccurred())
			}).To(Panic())
			Expect(func() {
				_, err := k.ProcessTransaction(ctx, &coretypes.Transaction{})
				Expect(err).To(HaveOccurred())
			}).To(Panic())
		})

		It("should successfully deploy a valid contract and call it", func() {
			legacyTxData.Data = common.FromHex(generated.SolmateERC20Bin)
			tx := coretypes.MustSignNewTx(key, signer, legacyTxData)
			addr, err := signer.Sender(tx)
			Expect(err).ToNot(HaveOccurred())
			k.GetStatePlugin().CreateAccount(addr)
			k.GetStatePlugin().AddBalance(addr, big.NewInt(1000000000))
			k.GetStatePlugin().Finalize()

			// create the contract
			result, err := k.ProcessTransaction(ctx, tx)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Err).ToNot(HaveOccurred())
			// call the contract non-view function
			deployAddress := crypto.CreateAddress(crypto.PubkeyToAddress(key.PublicKey), 0)
			legacyTxData.To = &deployAddress
			var solmateABI abi.ABI
			err = solmateABI.UnmarshalJSON([]byte(generated.SolmateERC20ABI))
			Expect(err).ToNot(HaveOccurred())
			input, err := solmateABI.Pack("mint", common.BytesToAddress([]byte{0x88}), big.NewInt(8888888))
			Expect(err).ToNot(HaveOccurred())
			legacyTxData.Data = input
			legacyTxData.Nonce++
			tx = coretypes.MustSignNewTx(key, signer, legacyTxData)
			result, err = k.ProcessTransaction(ctx, tx)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Err).ToNot(HaveOccurred())

			// call the contract view function
			legacyTxData.Data = crypto.Keccak256Hash([]byte("totalSupply()")).Bytes()[:4]
			legacyTxData.Nonce++
			tx = coretypes.MustSignNewTx(key, signer, legacyTxData)
			result, err = k.ProcessTransaction(ctx, tx)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Err).ToNot(HaveOccurred())
		})
	})
})
