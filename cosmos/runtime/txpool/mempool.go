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
	"context"
	"errors"
	"sync/atomic"
	"time"

	"cosmossdk.io/log"

	"github.com/berachain/polaris/cosmos/x/evm/types"
	"github.com/berachain/polaris/eth"
	"github.com/berachain/polaris/eth/core"
	"github.com/berachain/polaris/lib/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/mempool"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Mempool implements the mempool.Mempool & Lifecycle interfaces.
var (
	_ mempool.Mempool = (*Mempool)(nil)
	_ Lifecycle       = (*Mempool)(nil)
)

// Lifecycle represents a lifecycle object.
type Lifecycle interface {
	Start() error
	Stop() error
}

// GethTxPool is used for generating mocks.
type GethTxPool interface {
	eth.TxPool
}

// Mempool is a mempool that adheres to the cosmos mempool interface.
// It purposefully does not implement `Select` or `Remove` as the purpose of this mempool
// is to allow for transactions coming in from CometBFT's gossip to be added to the underlying
// geth txpool during `CheckTx`, that is the only purpose of `Mempool“.
type Mempool struct {
	txpool   eth.TxPool
	lifetime time.Duration
	chain    core.ChainReader
	handler  Lifecycle

	// when pause inserts is enabled, we use a channel to queue transactions
	// to be inserted into the txpool after a block is committed.
	pauseInserts atomic.Bool
	insertQueue  chan *ethtypes.Transaction
}

// New creates a new Mempool.
func New(chain core.ChainReader, txpool eth.TxPool, lifetime time.Duration) *Mempool {
	return &Mempool{
		txpool:       txpool,
		chain:        chain,
		lifetime:     lifetime,
		insertQueue:  make(chan *ethtypes.Transaction, 30000), // TODO: needs to be equal to comet mempoool size.
		pauseInserts: atomic.Bool{},
	}
}

// Init initializes the Mempool (notably the TxHandler).
func (m *Mempool) Init(
	logger log.Logger,
	txBroadcaster TxBroadcaster,
	txSerializer TxSerializer,
) {
	m.handler = newHandler(txBroadcaster, m.txpool, txSerializer, logger)
}

// Start starts the Mempool TxHandler.
func (m *Mempool) Start() error {
	go m.processInserts(context.Background())
	return m.handler.Start()
}

// Stop stops the Mempool TxHandler.
func (m *Mempool) Stop() error {
	return m.handler.Stop()
}

// Insert attempts to insert a Tx into the app-side mempool returning
// an error upon failure.
func (m *Mempool) Insert(ctx context.Context, sdkTx sdk.Tx) error {
	sCtx := sdk.UnwrapSDKContext(ctx)
	msgs := sdkTx.GetMsgs()
	if len(msgs) != 1 {
		return errors.New("only one message is supported")
	}

	wet, ok := utils.GetAs[*types.WrappedEthereumTransaction](msgs[0])
	if !ok {
		// We have to return nil for non-ethereum transactions as to not fail check-tx.
		return nil
	}

	// If we are currently protecting against block inserts, we queue the transaction
	// to be inserted until after we are ready.
	select {
	case <-sCtx.Done():
		return sCtx.Err()
	case m.insertQueue <- wet.Unwrap():
		return nil
	}
}

// processInserts processes inserts into the txpool.
func (m *Mempool) processInserts(ctx context.Context) error {
	txs := make([]*ethtypes.Transaction, 0)
	timer := time.NewTimer(500 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			if !(m.pauseInserts.Load()) {
				_ = m.txpool.Add(txs, false, false)
				txs = make([]*ethtypes.Transaction, 0)
			}
			continue
		case tx := <-m.insertQueue:
			txs = append(txs, tx)
		default:
			return nil
		}
	}
}

// CountTx returns the number of transactions currently in the mempool.
func (m *Mempool) CountTx() int {
	runnable, blocked := m.txpool.Stats()
	return runnable + blocked
}

// Select is an intentional no-op as we use a custom prepare proposal.
func (m *Mempool) Select(context.Context, [][]byte) mempool.Iterator {
	return nil
}

// Remove is an intentional no-op as the eth txpool handles removals.
func (m *Mempool) Remove(sdk.Tx) error {
	return nil
}
