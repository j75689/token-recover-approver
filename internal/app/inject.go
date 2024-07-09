package app

import (
	"golang.org/x/sync/errgroup"

	"github.com/bnb-chain/token-recover-app/internal/config"
	"github.com/bnb-chain/token-recover-app/internal/module/http"
	"github.com/bnb-chain/token-recover-app/internal/module/tracker"
	"github.com/bnb-chain/token-recover-app/internal/store"
	"github.com/bnb-chain/token-recover-app/internal/version"

	"github.com/rs/zerolog"
)

type Modules string

func (m Modules) String() string {
	return string(m)
}

const (
	APIModule     Modules = "api"
	TrackerModule Modules = "tracker"
	BotModule     Modules = "bot"
)

type Application struct {
	logger       *zerolog.Logger
	config       *config.Config
	httpServer   *http.HttpServer
	eventTracker *tracker.EventTracker
	store        store.GeneralStore
}

func (application Application) Start(modules []Modules) error {
	application.logger.Info().Str("app_version", version.AppVersion).Str("git_commit", version.GitCommit).Str("git_commit_date", version.GitCommitDate).Msg("version info")
	eg := errgroup.Group{}
	for _, module := range modules {
		switch module {
		case APIModule:
			eg.Go(func() error {
				application.logger.Info().Msgf("http server listen %s:%d", application.config.HTTP.Addr, application.config.HTTP.Port)
				return application.httpServer.Run(application.config.HTTP)
			})
		case TrackerModule:
			eg.Go(func() error {
				return application.eventTracker.StartListeningTokenRecoverEvent()
			})
		case BotModule:
			//TODO
		}
	}

	eg.Go(func() error {
		if !application.config.Metrics.Enable {
			return nil
		}
		application.logger.Info().Msgf("metrics server listen %s:%d", application.config.Metrics.Addr, application.config.Metrics.Port)
		return application.httpServer.RunMetrics(application.config.Metrics)
	})

	return eg.Wait()
}

func (application Application) Stop() error {
	application.logger.Info().Msg("shutdown http server ...")
	if err := application.httpServer.Shutdown(); err != nil {
		return err
	}
	application.logger.Info().Msg("http server is closed")

	application.logger.Info().Msg("shutdown store ...")
	if err := application.store.Close(); err != nil {
		return err
	}
	application.logger.Info().Msg("store is closed")
	return nil
}

func newApplication(
	logger *zerolog.Logger,
	config *config.Config,
	store store.GeneralStore,
	eventTracker *tracker.EventTracker,
	httpServer *http.HttpServer,
) Application {
	return Application{
		logger:       logger,
		config:       config,
		store:        store,
		eventTracker: eventTracker,
		httpServer:   httpServer,
	}
}
