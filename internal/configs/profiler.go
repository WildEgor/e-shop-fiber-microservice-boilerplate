package configs

import (
	"github.com/spf13/viper"
	"log/slog"
)

// ProfilerConfig holds profiler configurations
type ProfilerConfig struct {
	API string `mapstructure:"api"`
}

// NewProfilerConfig create profiler config
func NewProfilerConfig(c *Configurator) *ProfilerConfig {
	cfg := &ProfilerConfig{}

	updater := func() {
		if err := viper.UnmarshalKey("profiler", cfg); err != nil {
			slog.Error("profiler parse error", slog.Any("err", err))
			panic("profiler parse error")
		}
	}
	c.Register("logger config", updater)

	updater()

	return cfg
}
