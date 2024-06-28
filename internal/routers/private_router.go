package routers

import (
	"github.com/gofiber/fiber/v3"
)

var _ Router[*fiber.App] = (*PrivateRouter)(nil)

// PrivateRouter router
type PrivateRouter struct {
}

// NewPrivateRouter creates router
func NewPrivateRouter() *PrivateRouter {
	return &PrivateRouter{}
}

// Setup router
func (r *PrivateRouter) Setup(app *fiber.App) {
}
