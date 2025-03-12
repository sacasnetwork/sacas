// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)

package testdata

import (
	contractutils "github.com/sacasnetwork/sacas/v1/contracts/utils"
	evmtypes "github.com/sacasnetwork/sacas/v1/x/evm/types"
)

// LoadMaliciousDelayedContract loads the ERC20MaliciousDelayed contract.
//
// This is an evil token. Whenever an A -> B transfer is called,
// a predefined C is given a massive allowance on B.
func LoadMaliciousDelayedContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ERC20MaliciousDelayed.json")
}
