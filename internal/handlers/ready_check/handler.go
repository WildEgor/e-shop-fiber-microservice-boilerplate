package ready_check_handler

import (
	"github.com/gofiber/fiber/v3"
)

// ReadyCheckHandler represent ready check
type ReadyCheckHandler struct {
}

// NewReadyCheckHandler creates new handler
func NewReadyCheckHandler() *ReadyCheckHandler {
	return &ReadyCheckHandler{}
}

// Handle ready
// HealthCheck 		godoc
// @Summary			Ready check service
// @Description		Ready check service
// @Tags			Health Controller
// @Accept			json
// @Produce			json
// @Success			200
// @Router			/api/v1/readyz [get]
func (hch *ReadyCheckHandler) Handle(ctx fiber.Ctx) error {
	// TODO: Add own checks
	return nil
}
