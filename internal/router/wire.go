package router

import (
	"github.com/WildEgor/fibergo-microservice-boilerplate/internal/handlers"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(
	handlers.HandlersSet,
	NewPublicRouter,
	NewPrivateRouter,
	NewSwaggerRouter,
)
