// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)
package ibc

import "errors"

var (
	ErrNoIBCVoucherDenom  = errors.New("denom is not an IBC voucher")
	ErrDenomTraceNotFound = errors.New("denom trace not found")
	ErrInvalidBaseDenom   = errors.New("invalid base denomination")
)
