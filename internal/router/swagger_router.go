package router

import (
	"github.com/gofiber/fiber/v3"
)

type SwaggerRouter struct {
}

func NewSwaggerRouter() *SwaggerRouter {
	return &SwaggerRouter{}
}

// Setup func for describe group of API Docs routes.
func (sr *SwaggerRouter) Setup(app *fiber.App) {

}
