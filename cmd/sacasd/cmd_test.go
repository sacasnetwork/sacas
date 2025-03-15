package main_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/sacasnetwork/sacas/v15/app"
	sacasd "github.com/sacasnetwork/sacas/v15/cmd/sacasd"
	"github.com/sacasnetwork/sacas/v15/utils"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := sacasd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",       // Test the init cmd
		"sacas-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, utils.TestnetChainID+"-1"),
	})

	err := svrcmd.Execute(rootCmd, "sacasd", app.DefaultNodeHome)
	require.NoError(t, err)
}

func TestAddKeyLedgerCmd(t *testing.T) {
	rootCmd, _ := sacasd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"keys",
		"add",
		"dev0",
		fmt.Sprintf("--%s", flags.FlagUseLedger),
	})

	err := svrcmd.Execute(rootCmd, "SACASD", app.DefaultNodeHome)
	require.Error(t, err)
}
