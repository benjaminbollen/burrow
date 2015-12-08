package db

import (
	cfg "github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/go-config"
)

var config cfg.Config = nil

func init() {
	cfg.OnConfig(func(newConfig cfg.Config) {
		config = newConfig
	})
}