package injection

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	gormLogger "gorm.io/gorm/logger"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/bnb-chain/token-recover-app/internal/common"
	"github.com/bnb-chain/token-recover-app/internal/config"
	"github.com/bnb-chain/token-recover-app/internal/module/tracker"
	"github.com/bnb-chain/token-recover-app/internal/store"
	"github.com/bnb-chain/token-recover-app/internal/store/gorm"
	"github.com/bnb-chain/token-recover-app/internal/store/memory"
)

type StoreType string

const (
	MemoryStore StoreType = "memory"
	GORMStore   StoreType = "gorm"
)

type TokenListProviderType string

const (
	FromURL  TokenListProviderType = "url"
	FromFile TokenListProviderType = "file"
)

func initSDK(config *config.Config, logger *zerolog.Logger) {
	logger.Info().Str("chain_id", config.ChainID).Msg("init sdk config")
	sdkConfig := types.GetConfig()
	sdkConfig.SetBech32PrefixForAccount("bnb", "bnbp")
	if config.ChainID != common.MainnetChainID {
		sdkConfig.SetBech32PrefixForAccount("tbnb", "bnbp")
		logger.Debug().Str("chain_id", config.ChainID).Msg("set bech32 prefix to tbnb")
	}
}

func InitTokenListProvider(config *config.Config, logger *zerolog.Logger) (tracker.TokenList, error) {
	var (
		jsonData []byte
		err      error
	)
	switch TokenListProviderType(config.TokenList.Provider) {
	case FromURL:
		resp, err := http.Get(config.TokenList.URL)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		jsonData, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	case FromFile:
		jsonData, err = os.ReadFile(config.TokenList.FilePath)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid token list provider type")
	}
	var tokens []tracker.TokenInfo
	err = json.Unmarshal(jsonData, &tokens)
	if err != nil {
		return nil, err
	}
	tokenList := make(tracker.TokenList, len(tokens))
	for _, token := range tokens {
		tokenList[token.Symbol] = token
	}
	logger.Info().Int("token_count", len(tokenList)).Msg("init token list")
	return tokenList, nil
}

func InitStore(config *config.Config, logger *zerolog.Logger) (store.GeneralStore, error) {
	initSDK(config, logger)
	logger.Debug().Str("store_type", config.Store.Driver).Msg("init store")
	switch StoreType(config.Store.Driver) {
	case MemoryStore:
		memStore, err := memory.NewMemoryStore(
			config.Store.MemoryStore.MerkleProofs,
		)
		return memStore, err
	case GORMStore:
		sqlStore, err := gorm.NewSQLStore(
			config,
			gorm.SetConnMaxLifetime(config.Store.SqlStore.MaxLifetime),
			gorm.SetConnMaxIdleTime(config.Store.SqlStore.MaxIdleTime),
			gorm.SetMaxIdleConns(config.Store.SqlStore.MaxIdleConn),
			gorm.SetMaxOpenConns(config.Store.SqlStore.MaxOpenConn),
			gorm.SetLogLevel(gormLogger.LogLevel(config.Store.SqlStore.LogLevel)),
		)
		return sqlStore, err
	default:
		return nil, errors.New("invalid store type")
	}
}
