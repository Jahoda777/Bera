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

package prepare

import (
	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/miner"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool"
	"pkg.berachain.dev/polaris/eth/polar"
)

type (
	// ProposalTxVerifier defines the interface that is implemented by BaseApp,
	// that any custom ABCI PrepareProposal and ProcessProposal handler can use
	// to verify a transaction.
	TxVerifier interface {
		PrepareProposalVerifyTx(tx sdk.Tx) ([]byte, error)
		ProcessProposalVerifyTx(txBz []byte) (sdk.Tx, error)
	}
)

type Handler struct {
	polaris    *polar.Polaris
	txVerifier TxVerifier
}

func NewHandler(txVerifier TxVerifier) Handler {
	return Handler{
		txVerifier: txVerifier,
	}
}

func (h *Handler) SetPolaris(polaris *polar.Polaris) {
	h.polaris = polaris
}

func (h *Handler) PrepareProposal(
	ctx sdk.Context, req *abci.RequestPrepareProposal,
) (*abci.ResponsePrepareProposal, error) {
	var maxBlockGas int64
	if b := ctx.ConsensusParams().Block; b != nil {
		maxBlockGas = b.MaxGas
	}

	var (
		selectedTxs  [][]byte
		totalTxBytes int64
		totalTxGas   uint64
	)

	// TODO: this is definitely big bad and should be fixed.
	// We should prime the blockchain object after app.Load().
	if h.polaris.TxPool() == nil {
		ctx.Logger().Error("waiting for txpool to be initialized, proposing empty block")
		return &abci.ResponsePrepareProposal{}, nil
	}

	pending := h.polaris.TxPool().Pending(false)
	txp, _ := h.polaris.Host().GetTxPoolPlugin().(txpool.Plugin)

	byPriceAndNonce := miner.NewTransactionsByPriceAndNonce(types.LatestSigner(
		h.polaris.Host().GetConfigurationPlugin().ChainConfig(),
	), pending, h.polaris.Miner().NextBaseFee())

	for _tx := byPriceAndNonce.Peek(); _tx != nil; _tx = byPriceAndNonce.Peek() {
		bz, err := txp.SerializeToBytes(_tx.Resolve())
		if err != nil {
			ctx.Logger().Error("Failed sdk.Tx Serialization", _tx.Resolve().Hash(), err)
			continue
		}
		txGasLimit := _tx.Tx.Gas()
		txSize := int64(len(bz))
		// only add the transaction to the proposal if we have enough capacity
		if (txSize + totalTxBytes) < req.MaxTxBytes {
			// If there is a max block gas limit, add the tx only if the limit has
			// not been met.
			if maxBlockGas > 0 {
				if (txGasLimit + totalTxGas) <= uint64(maxBlockGas) {
					totalTxGas += txGasLimit
					totalTxBytes += txSize
					selectedTxs = append(selectedTxs, bz)
				}
			} else {
				totalTxBytes += txSize
				selectedTxs = append(selectedTxs, bz)
			}
		}
		// Check if we've reached capacity. If so, we cannot select any more
		// transactions.
		if totalTxBytes >= req.MaxTxBytes ||
			(maxBlockGas > 0 && (totalTxGas >= uint64(maxBlockGas))) {
			break
		}

		byPriceAndNonce.Shift()
	}

	return &abci.ResponsePrepareProposal{Txs: selectedTxs}, nil
}
