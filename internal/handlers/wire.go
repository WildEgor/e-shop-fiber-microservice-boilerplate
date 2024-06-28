package handlers

import (
	eh "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/errors"
	hch "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/health_check"
	ph "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/ping"
	rh "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/ready_check"
	"github.com/google/wire"
)

// Set contains http/amqp/etc handlers (acts like facades)
var Set = wire.NewSet(
	eh.NewErrorsHandler,
	hch.NewHealthCheckHandler,
	rh.NewReadyCheckHandler,
	ph.NewPingHandler,
)
