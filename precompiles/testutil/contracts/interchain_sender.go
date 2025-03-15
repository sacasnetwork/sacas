// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/sacas/sacas/blob/main/LICENSE)

package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/sacasnetwork/sacas/v11/x/evm/types"
)

var (
	//go:embed InterchainSender.json
	InterchainSenderJSON []byte

	// InterchainSenderContract is the compiled contract calling the distribution precompile
	InterchainSenderContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(InterchainSenderJSON, &InterchainSenderContract)
	if err != nil {
		panic(err)
	}

	if len(InterchainSenderContract.Bin) == 0 {
		panic("failed to load smart contract that calls distribution precompile")
	}
}
