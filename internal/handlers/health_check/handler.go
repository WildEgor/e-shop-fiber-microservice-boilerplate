package health_check_handler

import (
	"github.com/gofiber/fiber/v3"
)

// HealthCheckHandler represent handler
type HealthCheckHandler struct {
}

// NewHealthCheckHandler creates new handler
func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

// Handle health
// HealthCheck 		godoc
// @Summary			Health check service
// @Description		Health check service
// @Tags			Health Controller
// @Accept			json
// @Produce			json
// @Success			200
// @Router			/api/v1/livez [get]
func (hch *HealthCheckHandler) Handle(ctx fiber.Ctx) error {
	// TODO: Add own checks
	return nil
}
