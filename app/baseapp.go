// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)

package app

import (
	"errors"
	"io"
)

// Close will be called in graceful shutdown in start cmd
func (app *Sacas) Close() error {
	errs := []error{app.BaseApp.Close()}

	// flush the versiondb
	if closer, ok := app.qms.(io.Closer); ok {
		errs = append(errs, closer.Close())
	}

	// mainly to flush memiavl
	if closer, ok := app.BaseApp.CommitMultiStore().(io.Closer); ok {
		errs = append(errs, closer.Close())
	}

	err := errors.Join(errs...)
	app.BaseApp.Logger().Info("Application gracefully shutdown", "error", err)
	return err
}
