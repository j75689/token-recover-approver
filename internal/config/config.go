package config

import (
	"os"
	"strings"

	"github.com/bnb-chain/airdrop-service/pkg/logger"

	"github.com/spf13/viper"
)

type Config struct {
	Logger LoggerConfig `mapstructure:"logger"`
	HTTP   HTTPConfig   `mapstructure:"http"`
}

type LoggerConfig struct {
	Level  string           `mapstructure:"level"`
	Format logger.LogFormat `mapstructure:"format"`
}

type HTTPConfig struct {
	Addr string `mapstructure:"addr"`
	Port uint16 `mapstructure:"port"`
}

func NewConfig(configPath string) (*Config, error) {
	var file *os.File
	file, _ = os.Open(configPath)

	v := viper.New()
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	/* default */
	v.SetDefault("logger.level", "INFO")
	v.SetDefault("logger.format", logger.ConsoleFormat)
	v.SetDefault("http.addr", "0.0.0.0")
	v.SetDefault("http.port", "8080")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.ReadConfig(file)

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return &config, err
	}

	return &config, nil
}
