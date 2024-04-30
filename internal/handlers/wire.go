package handlers

import (
	eh "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/errors"
	hch "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/health_check"
	rch "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/ready_check"
	"github.com/google/wire"
)

var HandlersSet = wire.NewSet(
	eh.NewErrorsHandler,
	hch.NewHealthCheckHandler,
	rch.NewReadyCheckHandler,
)
