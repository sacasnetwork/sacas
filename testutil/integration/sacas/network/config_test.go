// Copyright SacDev0S . (SacaS)

package network_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	grpchandler "github.com/sacasnetwork/sacas/v20/testutil/integration/sacas/grpc"
	testkeyring "github.com/sacasnetwork/sacas/v20/testutil/integration/sacas/keyring"
	"github.com/sacasnetwork/sacas/v20/testutil/integration/sacas/network"
	"github.com/sacasnetwork/sacas/v20/utils"
	"github.com/stretchr/testify/require"
)

func TestWithBalances(t *testing.T) {
	key1Balance := sdk.NewCoins(sdk.NewInt64Coin(utils.BaseDenom, 1e18))
	key2Balance := sdk.NewCoins(
		sdk.NewInt64Coin(utils.BaseDenom, 2e18),
		sdk.NewInt64Coin("other", 3e18),
	)

	// Create a new network with 2 pre-funded accounts
	keyring := testkeyring.New(2)
	balances := []banktypes.Balance{
		{
			Address: keyring.GetAccAddr(0).String(),
			Coins:   key1Balance,
		},
		{
			Address: keyring.GetAccAddr(1).String(),
			Coins:   key2Balance,
		},
	}
	nw := network.New(
		network.WithBalances(balances...),
	)
	handler := grpchandler.NewIntegrationHandler(nw)

	req, err := handler.GetAllBalances(keyring.GetAccAddr(0))
	require.NoError(t, err, "error getting balances")
	require.Len(t, req.Balances, 1, "wrong number of balances")
	require.Equal(t, balances[0].Coins, req.Balances, "wrong balances")

	req, err = handler.GetAllBalances(keyring.GetAccAddr(1))
	require.NoError(t, err, "error getting balances")
	require.Len(t, req.Balances, 2, "wrong number of balances")
	require.Equal(t, balances[1].Coins, req.Balances, "wrong balances")
}
