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

package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"pkg.berachain.dev/stargazer/x/evm/types"
)

// Compile-time check to ensure `Keeper` implements the `MsgServiceServer` interface.
var _ types.MsgServiceServer = &Keeper{}

// `EthTransaction` implements the  MsgServiceServer interface. It processes an incoming request
// and applies it to the Stargazer Chain.
func (k *Keeper) EthTransaction(
	ctx context.Context, msg *types.EthTransactionRequest,
) (*types.EthTransactionResponse, error) {
	sCtx := sdk.UnwrapSDKContext(ctx)

	sCtx = sdk.UnwrapSDKContext(ctx).
		WithKVGasConfig(storetypes.GasConfig{}).WithTransientKVGasConfig(storetypes.GasConfig{})

	// We zero-out the gas meter prior to evm execution in order to ensure that the receipt output from the EVM
	// is correct. In the future, we will revisit this to allow gas metering for more complex operations prior to entering the EVM.
	sCtx.GasMeter().RefundGas(sCtx.GasMeter().GasConsumed(), "reset gas meter prior to ethereum state transition")

	fmt.Println("GAS BEFORE", sCtx.GasMeter().GasConsumed())
	// Process the transaction and return the result.
	result, err := k.ProcessTransaction(sCtx, msg.AsTransaction())
	if err != nil {
		k.Logger(sCtx).Error("keeper.EthTransaction", "error", err)
		return nil, errorsmod.Wrapf(err, "failed to process transaction")
	}
	fmt.Println("GAS AFTER", sCtx.GasMeter().GasConsumed())
	k.Logger(sdk.UnwrapSDKContext(ctx)).Info("keeper.EthTransaction", "exec_result", result)
	// Build response and return.

	vmErr := ""
	if result.Err != nil {
		vmErr = result.Err.Error()
	}

	return &types.EthTransactionResponse{
		GasUsed:    result.UsedGas,
		VmError:    vmErr,
		ReturnData: result.ReturnData,
	}, nil
}

// `UpdateParams`  processes an incoming request and applies it to the Configuration plugin to
// update things about both the Stargazer Chain as well as the EVM Module.
func (k *Keeper) UpdateParams(
	ctx context.Context, req *types.UpdateParamsRequest,
) (*types.UpdateParamsResponse, error) {
	// Ensure the authority is valid.
	if k.authority != req.Authority {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority, expected %s, got %s", k.authority, req.Authority,
		)
	}

	// Update the params.
	k.cp.Prepare(ctx)
	k.cp.SetParams(&req.Params)
	return &types.UpdateParamsResponse{}, nil
}
