package configs

import (
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

var (
	LogJsonFormat string = "json"
)

var logLevelToSlogLevel = map[string]slog.Leveler{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
}

// ProfilerConfig holds logger configurations
type LoggerConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

func NewLoggerConfig(c *Configurator) *LoggerConfig {
	cfg := &LoggerConfig{}

	updater := func() {
		if err := viper.UnmarshalKey("logger", cfg); err != nil {
			slog.Error("logger parse error", slog.Any("err", err))
			panic("logger parse error")
		}

		slog.Info("config", slog.Any("value", cfg))

		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevelToSlogLevel[cfg.Level],
		}))
		if cfg.Format == LogJsonFormat {
			logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: logLevelToSlogLevel[cfg.Level],
			}))
		}
		slog.SetDefault(logger)
	}
	c.Register("logger config", updater)

	updater()

	return cfg
}
