// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/bnb-chain/token-recover-app/internal/config"
	"github.com/bnb-chain/token-recover-app/internal/injection"
	"github.com/bnb-chain/token-recover-app/internal/module/approval"
	"github.com/bnb-chain/token-recover-app/internal/module/http"
	"github.com/bnb-chain/token-recover-app/internal/module/tracker"
)

// Injectors from wire.go:

func Initialize(configPath string) (Application, error) {
	configConfig, err := config.NewConfig(configPath)
	if err != nil {
		return Application{}, err
	}
	logger, err := injection.InitLogger(configConfig)
	if err != nil {
		return Application{}, err
	}
	generalStore, err := injection.InitStore(configConfig, logger)
	if err != nil {
		return Application{}, err
	}
	keyManager, err := injection.InitKeyManager(configConfig)
	if err != nil {
		return Application{}, err
	}
	eventTracker := tracker.NewEventTracker(configConfig, keyManager, generalStore, logger)
	registry := injection.InitPrometheusRegister()
	metrics := injection.InitMetrics(registry)
	approvalService, err := approval.NewApprovalService(configConfig, keyManager, generalStore, metrics, logger)
	if err != nil {
		return Application{}, err
	}
	tokenList, err := injection.InitTokenListProvider(configConfig, logger)
	if err != nil {
		return Application{}, err
	}
	httpServer := http.NewHttpServer(approvalService, generalStore, tokenList, registry, logger)
	application := newApplication(logger, configConfig, generalStore, eventTracker, httpServer)
	return application, nil
}
