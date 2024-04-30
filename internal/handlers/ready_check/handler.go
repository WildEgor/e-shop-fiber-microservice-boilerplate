package ready_check_handler

import (
	"github.com/gofiber/fiber/v2"
)

type ReadyCheckHandler struct {
}

func NewReadCheckHandler() *ReadyCheckHandler {
	return &ReadyCheckHandler{}
}

// HealthCheck godoc
//
//	@Summary		Readness check service
//	@Description	Readness check service
//	@Tags			Health Controller
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/api/v1/readyz [get]
func (hch *ReadyCheckHandler) Handle(c *fiber.Ctx) error {
	// Add own checks
	return nil
}
