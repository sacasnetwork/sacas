// Copyright Tharsis Labs Ltd.(Sacas)
// SPDX-License-Identifier:ENCL-1.0(SacasINC)

//go:build rocksdb
// +build rocksdb

package main

import (
	"sort"

	"github.com/linxGnu/grocksdb"
	"github.com/spf13/cobra"

	versiondbclient "github.com/crypto-org-chain/cronos/versiondb/client"

	"github.com/sacasnetwork/sacas/v1/app"
	"github.com/sacasnetwork/sacas/v1/cmd/sacasd/opendb"
)

// ChangeSetCmd returns a Cobra command for interacting with change sets.
// NOTE: this is only included in builds with rocksdb
func ChangeSetCmd() *cobra.Command {
	keys, _, _ := app.StoreKeys()
	storeNames := make([]string, 0, len(keys))
	for name := range keys {
		storeNames = append(storeNames, name)
	}
	sort.Strings(storeNames)

	return versiondbclient.ChangeSetGroupCmd(versiondbclient.Options{
		DefaultStores:  storeNames,
		OpenReadOnlyDB: opendb.OpenReadOnlyDB,
		AppRocksDBOptions: func(sstFileWriter bool) *grocksdb.Options {
			return opendb.NewRocksdbOptions(nil, sstFileWriter)
		},
	})
	return nil
}
