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

package txpool

import (
	"time"

	tmock "github.com/stretchr/testify/mock"

	"cosmossdk.io/log"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mock"
	"pkg.berachain.dev/polaris/eth/core"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("", func() {
	var h *handler
	t := GinkgoT()

	var subscription *mock.Subscription
	var serializer *mock.TxSerializer
	var broadcaster *mock.Broadcaster
	var subprovider *mock.TxSubProvider

	BeforeEach(func() {
		subscription = mock.NewSubscription(t)
		subscription.On("Err").Return(nil)
		subscription.On("Unsubscribe").Return()
		broadcaster = mock.NewBroadcaster(t)
		subprovider = mock.NewTxSubProvider(t)
		subprovider.On("SubscribeNewTxsEvent", tmock.Anything).Return(subscription)
		serializer = mock.NewTxSerializer(t)
		h = newHandler(broadcaster, subprovider, serializer, log.NewTestLogger(t))
		h.Start()
		// Wait for handler to start.
		time.Sleep(300 * time.Millisecond)
	})

	AfterEach(func() {
		h.Stop()
		// Wait for handler to stop
		time.Sleep(300 * time.Millisecond)
	})

	When("", func() {
		It("should start", func() {
			Expect(h.Running()).To(BeTrue())
		})

		It("should handle 1 tx", func() {
			serializer.On("SerializeToBytes", tmock.Anything).Return([]byte{123}, nil).Once()
			broadcaster.On("BroadcastTxSync", []byte{123}).Return(nil, nil).Once()

			h.txsCh <- core.NewTxsEvent{Txs: []*coretypes.Transaction{coretypes.NewTx(&coretypes.LegacyTx{Nonce: 5})}}
		})

		It("should handle multiple tx", func() {
			serializer.On("SerializeToBytes", tmock.Anything).Return([]byte{123}, nil).Twice()
			broadcaster.On("BroadcastTxSync", []byte{123}).Return(nil, nil).Twice()

			h.txsCh <- core.NewTxsEvent{Txs: []*coretypes.Transaction{
				coretypes.NewTx(&coretypes.LegacyTx{Nonce: 5}),
				coretypes.NewTx(&coretypes.LegacyTx{Nonce: 6}),
			}}
		})
	})

})
