package router

import (
	hch "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/health_check"
	rch "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/ready_check"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/limiter"
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

func (r *PublicRouter) Setup(app *fiber.App) {
	api := app.Group("/api", limiter.New(limiter.Config{
		Max:                    10,
		SkipSuccessfulRequests: true,
	}))
	v1 := api.Group("/v1")

	v1.Get("/ping", func(ctx fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "pong",
		})
	})

	app.Get("/api/v1/livez", healthcheck.NewHealthChecker(healthcheck.Config{
		Probe: func(ctx fiber.Ctx) bool {
			if err := r.hch.Handle(ctx); err != nil {
				slog.Error("error not healthy")
				return false
			}

			slog.Debug("is healthy")

			return true
		},
	}))
	app.Get("/api/v1/readyz", healthcheck.NewHealthChecker(healthcheck.Config{
		Probe: func(ctx fiber.Ctx) bool {
			if err := r.rch.Handle(ctx); err != nil {
				slog.Error("error not ready")
				return false
			}

			slog.Debug("is ready")

			return true
		},
	}))
}
