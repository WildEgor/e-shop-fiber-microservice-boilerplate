package configs

import (
	"github.com/caarlos0/env/v7"
	"log/slog"
)

// AppConfig holds the main app configurations
type AppConfig struct {
	Name     string `env:"APP_NAME" envDefault:"app"`
	HttpPort string `env:"APP_HTTP_PORT" envDefault:"8888"`
	Mode     string `env:"APP_MODE,required"`
}

func NewAppConfig(c *Configurator) *AppConfig {
	cfg := AppConfig{}

	if err := env.Parse(&cfg); err != nil {
		slog.Error("app config parse error", slog.Any("err", err))
	}

	return &cfg
}

// IsProduction Check is application running in production mode
func (ac AppConfig) IsProduction() bool {
	return ac.Mode != "develop"
}
