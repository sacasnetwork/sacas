// Copyright SacDev0S . (SacaS)

package testdata

import (
	contractutils "github.com/sacasnetwork/sacas/v20/contracts/utils"
	evmtypes "github.com/sacasnetwork/sacas/v20/x/evm/types"
)

func LoadERC20MinterV5Contract() (evmtypes.CompiledContract, error) {
	return contractutils.LegacyLoadContractFromJSONFile("ERC20Minter_OpenZeppelinV5.json")
}
