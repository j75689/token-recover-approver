package injection

import (
	"github.com/rs/zerolog"

	"github.com/bnb-chain/token-recover-app/internal/config"
	"github.com/bnb-chain/token-recover-app/internal/version"
	"github.com/bnb-chain/token-recover-app/pkg/logger"
)

func InitLogger(config *config.Config) (*zerolog.Logger, error) {
	return logger.NewLogger(config.Logger.Level, config.Logger.Format, logger.WithStr("app_id", version.APPName))
}
