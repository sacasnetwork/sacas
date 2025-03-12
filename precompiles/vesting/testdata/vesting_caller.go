// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)

package testdata

import (
	contractutils "github.com/sacasnetwork/sacas/v1/contracts/utils"
	evmtypes "github.com/sacasnetwork/sacas/v1/x/evm/types"
)

func LoadVestingCallerContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("VestingCaller.json")
}
