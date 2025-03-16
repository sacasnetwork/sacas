// Copyright SacDev0S . (SacaS)

package testdata

import (
	contractutils "github.com/sacasnetwork/sacas/v20/contracts/utils"
	evmtypes "github.com/sacasnetwork/sacas/v20/x/evm/types"
)

func LoadVestingCallerContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("VestingCaller.json")
}
