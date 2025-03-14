// Copyright SacDev0S . (SacaS)
package types

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"

	"github.com/sacasnetwork/sacas/v15/types"
	"github.com/sacasnetwork/sacas/v15/utils"
)

var (
	// DefaultEVMDenom defines the default EVM denomination on Sacas
	DefaultEVMDenom = utils.BaseDenom
	// DefaultAllowUnprotectedTxs rejects all unprotected txs (i.e false)
	DefaultAllowUnprotectedTxs = false
	// DefaultEnableCreate enables contract creation (i.e true)
	DefaultEnableCreate = true
	// DefaultEnableCall enables contract calls (i.e true)
	DefaultEnableCall = true
	// AvailableEVMExtensions defines the default active precompiles
	AvailableEVMExtensions = []string{
		"0x0000000000000000000000000000000000000800", // Staking precompile
		"0x0000000000000000000000000000000000000801", // Distribution precompile
		"0x0000000000000000000000000000000000000802", // ICS20 transfer precompile
		"0x0000000000000000000000000000000000000803", // Vesting precompile
	}
	// DefaultExtraEIPs defines the default extra EIPs to be included
	// On v15, EIP 3855 was enabled
	DefaultExtraEIPs = []int64{3855}
)

// NewParams creates a new Params instance
func NewParams(
	evmDenom string,
	allowUnprotectedTxs,
	enableCreate,
	enableCall bool,
	config ChainConfig,
	extraEIPs []int64,
	activePrecompiles ...string,
) Params {
	return Params{
		EvmDenom:            evmDenom,
		AllowUnprotectedTxs: allowUnprotectedTxs,
		EnableCreate:        enableCreate,
		EnableCall:          enableCall,
		ExtraEIPs:           extraEIPs,
		ChainConfig:         config,
		ActivePrecompiles:   activePrecompiles,
	}
}

// DefaultParams returns default evm parameters
// ExtraEIPs is empty to prevent overriding the latest hard fork instruction set
// ActivePrecompiles is empty to prevent overriding the default precompiles
// from the EVM configuration.
func DefaultParams() Params {
	return Params{
		EvmDenom:            DefaultEVMDenom,
		EnableCreate:        DefaultEnableCreate,
		EnableCall:          DefaultEnableCall,
		ChainConfig:         DefaultChainConfig(),
		ExtraEIPs:           DefaultExtraEIPs,
		AllowUnprotectedTxs: DefaultAllowUnprotectedTxs,
		ActivePrecompiles:   AvailableEVMExtensions,
	}
}

// Validate performs basic validation on evm parameters.
func (p Params) Validate() error {
	if err := validateEVMDenom(p.EvmDenom); err != nil {
		return err
	}

	if err := validateEIPs(p.ExtraEIPs); err != nil {
		return err
	}

	if err := validateBool(p.EnableCall); err != nil {
		return err
	}

	if err := validateBool(p.EnableCreate); err != nil {
		return err
	}

	if err := validateBool(p.AllowUnprotectedTxs); err != nil {
		return err
	}

	if err := validateChainConfig(p.ChainConfig); err != nil {
		return err
	}

	return validatePrecompiles(p.ActivePrecompiles)
}

// EIPs returns the ExtraEIPS as a int slice
func (p Params) EIPs() []int {
	eips := make([]int, len(p.ExtraEIPs))
	for i, eip := range p.ExtraEIPs {
		eips[i] = int(eip)
	}
	return eips
}

// HasCustomPrecompiles returns true if the ActivePrecompiles slice is not empty.
func (p Params) HasCustomPrecompiles() bool {
	return len(p.ActivePrecompiles) > 0
}

// GetActivePrecompilesAddrs is a util function that the Active Precompiles
// as a slice of addresses.
func (p Params) GetActivePrecompilesAddrs() []common.Address {
	precompiles := make([]common.Address, len(p.ActivePrecompiles))
	for i, precompile := range p.ActivePrecompiles {
		precompiles[i] = common.HexToAddress(precompile)
	}
	return precompiles
}

func validateEVMDenom(i interface{}) error {
	denom, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter EVM denom type: %T", i)
	}

	return sdk.ValidateDenom(denom)
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateEIPs(i interface{}) error {
	eips, ok := i.([]int64)
	if !ok {
		return fmt.Errorf("invalid EIP slice type: %T", i)
	}

	uniqueEIPs := make(map[int64]struct{})

	for _, eip := range eips {
		if !vm.ValidEip(int(eip)) {
			return fmt.Errorf("EIP %d is not activateable, valid EIPS are: %s", eip, vm.ActivateableEips())
		}

		if _, ok := uniqueEIPs[eip]; ok {
			return fmt.Errorf("found duplicate EIP: %d", eip)
		}
		uniqueEIPs[eip] = struct{}{}
	}

	return nil
}

func validateChainConfig(i interface{}) error {
	cfg, ok := i.(ChainConfig)
	if !ok {
		return fmt.Errorf("invalid chain config type: %T", i)
	}

	return cfg.Validate()
}

func validatePrecompiles(i interface{}) error {
	precompiles, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid precompile slice type: %T", i)
	}

	seenPrecompiles := make(map[string]bool)
	for _, precompile := range precompiles {
		if seenPrecompiles[precompile] {
			return fmt.Errorf("duplicate precompile %s", precompile)
		}

		if err := types.ValidateAddress(precompile); err != nil {
			return fmt.Errorf("invalid precompile %s", precompile)
		}

		seenPrecompiles[precompile] = true
	}

	return nil
}

// IsLondon returns if london hardfork is enabled.
func IsLondon(ethConfig *params.ChainConfig, height int64) bool {
	return ethConfig.IsLondon(big.NewInt(height))
}
