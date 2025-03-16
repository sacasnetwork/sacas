// Copyright SacDev0S . (SacaS)
package codec

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/sacasnetwork/sacas/v20/crypto/ethsecp256k1"
)

// RegisterInterfaces register the Sacas key concrete types.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PrivKey)(nil), &ethsecp256k1.PrivKey{})
}
