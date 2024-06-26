package configs

import (
	"github.com/spf13/viper"
	"log/slog"
)

// AppConfig holds the main app configurations
type AppConfig struct {
	Name     string `mapstructure:"name"`
	Mode     string `mapstructure:"mode"`
	HttpPort string `mapstructure:"http_port"`
	changer  func(ac *AppConfig)
}

func NewAppConfig(c *Configurator) *AppConfig {
	cfg := &AppConfig{}

	updater := func() {
		if err := viper.UnmarshalKey("app", &cfg); err != nil {
			slog.Error("app config parse error")
			panic("logger parse error")
		}

		slog.Info("config", slog.Any("value", cfg))
	}
	c.Register("app config", func() {
		updater()
		cfg.changer(cfg)
	})

	updater()

	return cfg
}

func (ac *AppConfig) OnChanged(fn func(ac *AppConfig)) {
	ac.changer = fn
}

// IsProduction Check is application running in production mode
func (ac *AppConfig) IsProduction() bool {
	return ac.Mode == "production"
}

// IsDebug Check is application running in debug mode
func (ac *AppConfig) IsDebug() bool {
	return ac.Mode == "debug"
}
