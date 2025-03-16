// Copyright SacDev0S . (SacaS)

package app

import (
	"github.com/sacasnetwork/sacas/v20/app/eips"
	evmconfig "github.com/sacasnetwork/sacas/v20/x/evm/config"
	"github.com/sacasnetwork/sacas/v20/x/evm/core/vm"
)

// The init function of the config file allows to setup the global
// configuration for the EVM, modifying the custom ones defined in evmOS.
func init() {
	err := evmconfig.NewEVMConfigurator().
		WithExtendedEips(sacasActivators).
		Configure()
	if err != nil {
		panic(err)
	}
}

// SacasActivators defines a map of opcode modifiers associated
// with a key defining the corresponding EIP.
var sacasActivators = map[string]func(*vm.JumpTable){
	"sacas_0": eips.Enable0000,
	"sacas_1": eips.Enable0001,
	"sacas_2": eips.Enable0002,
}
