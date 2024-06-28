//go:build wireinject
// +build wireinject

package pkg

import (
	"github.com/google/wire"
)

// ServerSet
var ServerSet = wire.NewSet(AppSet)

// NewServer
func NewServer() (*Server, error) {
	wire.Build(ServerSet)
	return nil, nil
}
