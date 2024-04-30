package router

import (
	hch "github.com/WildEgor/fibergo-microservice-boilerplate/internal/handlers/health-check"
	rch "github.com/WildEgor/fibergo-microservice-boilerplate/internal/handlers/readness-check"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"log/slog"
)

type PublicRouter struct {
	hch *hch.HealthCheckHandler
	rch *rch.ReadyCheckHandler
}

func NewPublicRouter(
	hh *hch.HealthCheckHandler,
	rch *rch.ReadyCheckHandler,
) *PublicRouter {
	return &PublicRouter{
		hh,
		rch,
	}
}

func (r *PublicRouter) SetupPublicRouter(app *fiber.App) {

	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			if err := r.rch.Handle(c); err != nil {
				slog.Error("NOT OK")
				return false
			}

			slog.Debug("OK")

			return true
		},
		ReadinessProbe: func(c *fiber.Ctx) bool {
			if err := r.hch.Handle(c); err != nil {
				slog.Error("NOT OK")
				return false
			}

			slog.Debug("OK")

			return true
		},
		LivenessEndpoint:  "/api/v1/livez",
		ReadinessEndpoint: "/api/v1/readyz",
	}))

}
