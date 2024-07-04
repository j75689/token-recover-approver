//go:build wireinject
// +build wireinject

//The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"

	"github.com/bnb-chain/token-recover-app/internal/config"
	"github.com/bnb-chain/token-recover-app/internal/injection"
	"github.com/bnb-chain/token-recover-app/internal/module/approval"
	"github.com/bnb-chain/token-recover-app/internal/module/http"
)

func Initialize(configPath string) (Application, error) {
	wire.Build(
		newApplication,
		config.NewConfig,
		injection.InitLogger,
		injection.InitKeyManager,
		injection.InitStore,
		injection.InitMetrics,
		injection.InitPrometheusRegister,
		approval.NewApprovalService,
		http.NewHttpServer,
	)
	return Application{}, nil
}
