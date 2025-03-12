// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)

package types

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// HexString is a byte array that serializes to hex
type HexString []byte

// MarshalJSON serializes ByteArray to hex
func (s HexString) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%x", string(s)))
}

// UnmarshalJSON deserializes ByteArray to hex
func (s *HexString) UnmarshalJSON(data []byte) error {
	var x string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	str, err := hex.DecodeString(x)
	if err != nil {
		return err
	}
	*s = str
	return nil
}

// CompiledContract contains compiled bytecode and abi
type CompiledContract struct {
	ABI abi.ABI
	Bin HexString
}

type jsonCompiledContract struct {
	ABI string
	Bin HexString
}

// MarshalJSON serializes ByteArray to hex
func (s CompiledContract) MarshalJSON() ([]byte, error) {
	abi1, err := json.Marshal(s.ABI)
	if err != nil {
		return nil, err
	}
	return json.Marshal(jsonCompiledContract{ABI: string(abi1), Bin: s.Bin})
}

// UnmarshalJSON deserializes ByteArray to hex
func (s *CompiledContract) UnmarshalJSON(data []byte) error {
	var x jsonCompiledContract
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}

	s.Bin = x.Bin
	if err := json.Unmarshal([]byte(x.ABI), &s.ABI); err != nil {
		return fmt.Errorf("failed to unmarshal ABI: %w", err)
	}

	return nil
}

// HardhatCompiledContract is a type used to unpack the compiled JSON data
// which is generated by running `npx hardhat compile` for a given smart contract.
type HardhatCompiledContract struct {
	Format       string  `json:"_format"`
	ContractName string  `json:"contractName"`
	SourceName   string  `json:"sourceName"`
	ABI          abi.ABI `json:"abi"`
	Bytecode     string  `json:"bytecode"`
}

func (c HardhatCompiledContract) ToCompiledContract() (CompiledContract, error) {
	strippedHex := strings.TrimPrefix(c.Bytecode, "0x")
	hexBytes, err := hex.DecodeString(strippedHex)
	if err != nil {
		return CompiledContract{}, fmt.Errorf("failed to decode hex string: %w", err)
	}

	return CompiledContract{
		ABI: c.ABI,
		Bin: hexBytes,
	}, nil
}
