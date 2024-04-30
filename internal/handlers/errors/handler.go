package error_handler

import (
	"errors"
	"github.com/gofiber/fiber/v3"
)

type ErrorsHandler struct {
}

func NewErrorsHandler() *ErrorsHandler {
	return &ErrorsHandler{}
}

func (hch *ErrorsHandler) Handle(ctx fiber.Ctx, err error) error {
	sc := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		sc = e.Code
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	if e != nil {
		ctx.Status(sc)
		return nil
	}

	return nil
}
