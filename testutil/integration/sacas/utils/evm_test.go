package utils_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sacasnetwork/sacas/v1/contracts"
	testfactory "github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/factory"
	testhandler "github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/grpc"
	testkeyring "github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/keyring"
	testnetwork "github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/network"
	"github.com/sacasnetwork/sacas/v1/testutil/integration/sacas/utils"
	evmtypes "github.com/sacasnetwork/sacas/v1/x/evm/types"
	"github.com/stretchr/testify/require"
)

func TestGetERC20Balance(t *testing.T) {
	keyring := testkeyring.New(1)
	network := testnetwork.NewUnitTestNetwork(
		testnetwork.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	handler := testhandler.NewIntegrationHandler(network)
	factory := testfactory.New(network, handler)

	sender := keyring.GetKey(0)
	mintAmount := big.NewInt(100)

	// Deploy an ERC-20 contract
	erc20Addr, err := factory.DeployContract(
		sender.Priv,
		evmtypes.EvmTxArgs{},
		testfactory.ContractDeploymentData{
			Contract:        contracts.ERC20MinterBurnerDecimalsContract,
			ConstructorArgs: []interface{}{"TestToken", "TT", uint8(18)},
		},
	)
	require.NoError(t, err, "failed to deploy contract")
	require.NoError(t, network.NextBlock(), "failed to advance block")

	balance, err := utils.GetERC20Balance(network, erc20Addr, sender.Addr)
	require.NoError(t, err, "failed to get ERC20 balance")
	require.Equal(t, common.Big0.Int64(), balance.Int64(), "expected no balance before minting")

	// Mint some tokens
	_, err = factory.ExecuteContractCall(
		sender.Priv,
		evmtypes.EvmTxArgs{
			To: &erc20Addr,
		},
		testfactory.CallArgs{
			ContractABI: contracts.ERC20MinterBurnerDecimalsContract.ABI,
			MethodName:  "mint",
			Args:        []interface{}{sender.Addr, mintAmount},
		},
	)
	require.NoError(t, err, "failed to mint tokens")

	require.NoError(t, network.NextBlock(), "failed to advance block")

	balance, err = utils.GetERC20Balance(network, erc20Addr, sender.Addr)
	require.NoError(t, err, "failed to get ERC20 balance")
	require.Equal(t, mintAmount.Int64(), balance.Int64(), "expected different balance after minting")
}
