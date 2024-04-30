package router

import (
	"github.com/gofiber/fiber/v2"
)

type SwaggerRouter struct {
}

func NewSwaggerRouter() *SwaggerRouter {
	return &SwaggerRouter{}
}

// SetupSwaggerRouter func for describe group of API Docs routes.
func (sr *SwaggerRouter) SetupSwaggerRouter(app *fiber.App) {
	//cfg := swagger.Config{
	//	Title:    "Swagger Doc",
	//	BasePath: "/",
	//	Path:     "doc",
	//	FilePath: "./docs/swagger.json",
	//}
	//
	//app.Use(swagger.New(cfg))
}
