package routers

import (
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers"
	"github.com/google/wire"
)

// Set acts like "controllers" for routing http or etc.
var Set = wire.NewSet(
	handlers.Set,
	NewPublicRouter,
	NewPrivateRouter,
	NewSwaggerRouter,
)
