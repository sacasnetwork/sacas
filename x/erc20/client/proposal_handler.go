// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/sacas/sacas/blob/main/LICENSE)

package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/sacasnetwork/sacas/v11/x/erc20/client/cli"
)

var (
	RegisterCoinProposalHandler          = govclient.NewProposalHandler(cli.NewRegisterCoinProposalCmd)
	RegisterERC20ProposalHandler         = govclient.NewProposalHandler(cli.NewRegisterERC20ProposalCmd)
	ToggleTokenConversionProposalHandler = govclient.NewProposalHandler(cli.NewToggleTokenConversionProposalCmd)
)
