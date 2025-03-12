// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)

package types

// vesting events
const (
	EventTypeCreateClawbackVestingAccount = "create_clawback_vesting_account"
	EventTypeFundVestingAccount           = "fund_vesting_account"
	EventTypeClawback                     = "clawback"
	EventTypeUpdateVestingFunder          = "update_vesting_funder"

	AttributeKeyCoins       = "coins"
	AttributeKeyStartTime   = "start_time"
	AttributeKeyAccount     = "account"
	AttributeKeyFunder      = "funder"
	AttributeKeyNewFunder   = "new_funder"
	AttributeKeyDestination = "destination"
)
