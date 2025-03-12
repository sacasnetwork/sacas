// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)

package contracts

import (
	contractutils "github.com/sacasnetwork/sacas/v1/contracts/utils"
	evmtypes "github.com/sacasnetwork/sacas/v1/x/evm/types"
)

func LoadInterchainSenderContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("InterchainSender.json")
}
