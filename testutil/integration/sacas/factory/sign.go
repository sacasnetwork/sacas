// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)
package factory

import (
	errorsmod "cosmossdk.io/errors"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sacasnetwork/sacas/v1/testutil/tx"
	evmtypes "github.com/sacasnetwork/sacas/v1/x/evm/types"
)

// SignMsgEthereumTx signs a MsgEthereumTx with the provided private key and chainID.
func (tf *IntegrationTxFactory) SignMsgEthereumTx(privKey cryptotypes.PrivKey, msgEthereumTx evmtypes.MsgEthereumTx) (evmtypes.MsgEthereumTx, error) {
	ethChainID := tf.network.GetEIP155ChainID()
	signer := gethtypes.LatestSignerForChainID(ethChainID)
	err := msgEthereumTx.Sign(signer, tx.NewSigner(privKey))
	if err != nil {
		return evmtypes.MsgEthereumTx{}, errorsmod.Wrap(err, "failed to sign transaction")
	}
	return msgEthereumTx, nil
}
