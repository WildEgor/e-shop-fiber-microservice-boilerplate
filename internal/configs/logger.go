package configs

import (
	"github.com/caarlos0/env/v7"
	"log/slog"
)

var lvls map[string]slog.Leveler = map[string]slog.Leveler{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
}

// LoggerConfig holds logger configurations
type LoggerConfig struct {
	Level  slog.Leveler
	Format string `env:"LOG_FORMAT" envDefault:"json"`
	level  string `env:"LOG_LEVEL" envDefault:"debug"`
}

func NewLoggerConfig(c *Configurator) *LoggerConfig {
	cfg := LoggerConfig{}

	if err := env.Parse(&cfg); err != nil {
		slog.Error("log config parse error", slog.Any("err", err))
	}

	cfg.Level = lvls[cfg.level]

	return &cfg
}

// IsJSON check format
func (c *LoggerConfig) IsJSON() bool {
	return c.Format != "json"
}
