// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package tool

import (
	"github.com/bnb-chain/token-recover-app/internal/config"
	"github.com/bnb-chain/token-recover-app/internal/injection"
)

// Injectors from wire.go:

func Initialize(configPath string) (*Tool, error) {
	configConfig, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}
	logger, err := injection.InitLogger(configConfig)
	if err != nil {
		return nil, err
	}
	store, err := injection.InitStore(configConfig, logger)
	if err != nil {
		return nil, err
	}
	tool := newTool(logger, configConfig, store)
	return tool, nil
}
