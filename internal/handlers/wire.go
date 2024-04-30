package handlers

import (
	eh "github.com/WildEgor/fibergo-microservice-boilerplate/internal/handlers/errors"
	hch "github.com/WildEgor/fibergo-microservice-boilerplate/internal/handlers/health-check"
	rch "github.com/WildEgor/fibergo-microservice-boilerplate/internal/handlers/readness-check"
	"github.com/google/wire"
)

var HandlersSet = wire.NewSet(
	eh.NewErrorsHandler,
	hch.NewHealthCheckHandler,
	rch.NewReadCheckHandler,
)
