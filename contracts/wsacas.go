// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)

package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/sacasnetwork/sacas/v1/x/evm/types"
)

var (
	//go:embed compiled_contracts/WSACAS.json
	WSACASJSON []byte

	// WSACASContract is the compiled contract of WSACAS
	WSACASContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(WSACASJSON, &WSACASContract)
	if err != nil {
		panic(err)
	}

	if len(WSACASContract.Bin) == 0 {
		panic("failed to load WSACAS smart contract")
	}
}
