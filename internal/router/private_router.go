package router

import (
	"github.com/gofiber/fiber/v2"
)

type PrivateRouter struct {
}

func NewPrivateRouter() *PrivateRouter {
	return &PrivateRouter{}
}

func (r *PrivateRouter) SetupPrivateRouter(app *fiber.App) {
	// v1 := app.Group("/api/v1")
}
