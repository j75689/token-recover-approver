package tool

import (
	"github.com/rs/zerolog"

	"github.com/bnb-chain/token-recover-app/internal/config"
	"github.com/bnb-chain/token-recover-app/internal/store"
)

type Tool struct {
	logger *zerolog.Logger
	config *config.Config

	store store.GeneralStore
}

func newTool(
	logger *zerolog.Logger,
	config *config.Config,
	store store.GeneralStore,
) *Tool {
	return &Tool{
		logger: logger,
		config: config,
		store:  store,
	}
}
