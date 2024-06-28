package configs

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log/slog"
)

// Configurator dummy
type Configurator struct {
	watchers []func()
}

// NewConfigurator create new Configurator
func NewConfigurator() *Configurator {
	c := &Configurator{
		watchers: make([]func(), 0),
	}
	c.load()

	return c
}

// Watch config
func (c *Configurator) Watch() {
	viper.OnConfigChange(func(e fsnotify.Event) {
		slog.Info(fmt.Sprintf("watchers len: %d", len(c.watchers)))

		for _, watcher := range c.watchers {
			watcher()
		}
	})
	viper.WatchConfig()
}

// Register watcher
func (c *Configurator) Register(name string, fn func()) {
	slog.Info("register watcher", slog.Any("value", name))
	c.watchers = append(c.watchers, fn)
}

// load Load env data from files (default: .env, .env.local)
func (c *Configurator) load() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		slog.Error("error loading config file", slog.Any("err", err))
		panic("error loading config file")
	}

	err := viper.MergeInConfig()
	if err != nil {
		slog.Error("error merge config file", slog.Any("err", err))
		return
	}
}
