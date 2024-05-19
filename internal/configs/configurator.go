package configs

import (
	"github.com/joho/godotenv"
	"log/slog"
)

// Configurator dummy
type Configurator struct{}

func NewConfigurator() *Configurator {
	c := &Configurator{}
	c.load()

	return c
}

// load Load env data from files (default: .env, .env.local)
func (c *Configurator) load() {
	if err := godotenv.Load(".env", ".env.local"); err != nil {
		slog.Error("error loading envs file", slog.Any("err", err))
		panic(err)
	}
}
