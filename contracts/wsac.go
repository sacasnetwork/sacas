// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0( Sacas Network )

package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/sacasnetwork/sacas/v1/x/evm/types"
)

var (
	//go:embed compiled_contracts/WSAC.json
	WSACJSON []byte

	// WSACContract is the compiled contract of WSAC
	WSACContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(WSACJSON, &WSACContract)
	if err != nil {
		panic(err)
	}

	if len(WSACContract.Bin) == 0 {
		panic("failed to load WSAC smart contract")
	}
}
