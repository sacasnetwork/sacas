// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/sacas/sacas/blob/main/LICENSE)

package types

import (
	errorsmod "cosmossdk.io/errors"
)

// errors
var (
	ErrBlockedAddress = errorsmod.Register(ModuleName, 2, "blocked address")
)
