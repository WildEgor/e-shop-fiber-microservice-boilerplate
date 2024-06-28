package ping_handler

import (
	"crypto/rand"
	"errors"
	"github.com/gofiber/fiber/v3"
)

// PingCheckHandler represent ping handler
type PingCheckHandler struct {
}

// NewPingHandler creates new handler
func NewPingHandler() *PingCheckHandler {
	return &PingCheckHandler{}
}

// Handle ping
// Ping 			godoc
// @Summary			Dummy ping
// @Description		Just random return 200 OK or error
// @Tags			Test Controller
// @Accept			json
// @Produce			json
// @Success			200
// @Router			/api/v1/ping [get]
func (hch *PingCheckHandler) Handle(ctx fiber.Ctx) error {
	c := 10
	b := make([]byte, c)
	n, err := rand.Read(b)
	if err != nil {
		return errors.New("error")
	}

	if n%2 == 0 {
		return errors.New("ooops")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "pong",
	})
}
