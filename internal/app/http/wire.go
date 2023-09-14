//go:build wireinject
// +build wireinject

//The build tag makes sure the stub is not built in the final build.

package http

import (
	"github.com/bnb-chain/airdrop-service/internal/config"
	"github.com/bnb-chain/airdrop-service/internal/delivery/http"
	"github.com/bnb-chain/airdrop-service/internal/wireset"

	"github.com/google/wire"
)

func Initialize(configPath string) (Application, error) {
	wire.Build(
		newApplication,
		config.NewConfig,
		wireset.InitLogger,
		http.NewHttpServer,
	)
	return Application{}, nil
}
