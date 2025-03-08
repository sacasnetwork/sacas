// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0( Sacas Network )

package ante_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EvmAnteTestSuite struct {
	suite.Suite
}

func TestEvmAnteTestSuite(t *testing.T) {
	suite.Run(t, &EvmAnteTestSuite{})
}
