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

package types

import (
	fmt "fmt"

	"pkg.berachain.dev/polaris/eth/common"
)

// polarisDenomPrefix represents the bank module prefix all polaris erc20 module
// created denoms will have.
const polarisDenomPrefix = "polaris/"

// DenomForAddress returns the ERC20 denom for a given address.
func DenomForAddress(addr common.Address) string {
	return fmt.Sprintf("%s%s", polarisDenomPrefix, addr.Hex())
}

// IsPolarisDenom returns true if the address is
// a Polaris native token.
func IsPolarisDenom(denom string) bool {
	return len(denom) > 8 && denom[:8] == polarisDenomPrefix
}

// ShimHandlerType returns the type of the handler.
func ShimHandlerType(denoom string) int8 {
	// We mint burn for polaris controlled ERC20 tokens.
	if IsPolarisDenom(denoom) {
		return MintBurn
	}
	// We escrow for all other ERC20 tokens.
	return Escrow
}
