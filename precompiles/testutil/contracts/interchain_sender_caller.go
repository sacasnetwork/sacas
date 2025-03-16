// Copyright SacDev0S . (SacaS)

package contracts

import (
	contractutils "github.com/sacasnetwork/sacas/v20/contracts/utils"
	evmtypes "github.com/sacasnetwork/sacas/v20/x/evm/types"
)

func LoadInterchainSenderCallerContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("InterchainSenderCaller.json")
}
