// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)

package testdata

import (
	contractutils "github.com/sacasnetwork/sacas/v1/contracts/utils"
	evmtypes "github.com/sacasnetwork/sacas/v1/x/evm/types"
)

func LoadERC20TestCaller() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ERC20TestCaller.json")
}
