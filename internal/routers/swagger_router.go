package routers

import (
	_ "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/api/swagger"
	"github.com/gofiber/fiber/v3"
)

var _ Router[*fiber.App] = (*SwaggerRouter)(nil)

// SwaggerRouter router
type SwaggerRouter struct {
}

// NewSwaggerRouter creates new router
func NewSwaggerRouter() *SwaggerRouter {
	return &SwaggerRouter{}
}

// Setup func for describe group of API Docs routes.
func (sr *SwaggerRouter) Setup(app *fiber.App) {
	// TODO: fiber v3 not impl swagger middleware now
}
