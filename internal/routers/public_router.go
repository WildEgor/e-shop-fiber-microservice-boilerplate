package routers

import (
	hch "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/health_check"
	ph "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/ping"
	rch "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/ready_check"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"log/slog"
)

var _ Router[*fiber.App] = (*PublicRouter)(nil)

// PublicRouter router
type PublicRouter struct {
	hch *hch.HealthCheckHandler
	rch *rch.ReadyCheckHandler
	ph  *ph.PingCheckHandler
}

// NewPublicRouter creates router
func NewPublicRouter(
	hch *hch.HealthCheckHandler,
	rch *rch.ReadyCheckHandler,
	ph *ph.PingCheckHandler,
) *PublicRouter {
	return &PublicRouter{
		hch,
		rch,
		ph,
	}
}

// Setup router
func (r *PublicRouter) Setup(app *fiber.App) {
	api := app.Group("/api", limiter.New(limiter.Config{
		Max:                    10,
		SkipSuccessfulRequests: true,
	}))
	v1 := api.Group("/v1")

	// TODO: use for testing only (remove it)
	v1.Get("/ping", r.ph.Handle)

	v1.Get("/livez", healthcheck.NewHealthChecker(healthcheck.Config{
		Probe: func(ctx fiber.Ctx) bool {
			if err := r.hch.Handle(ctx); err != nil {
				slog.Error("error not healthy")
				return false
			}

			slog.Debug("is healthy")

			return true
		},
	}))
	v1.Get("/readyz", healthcheck.NewHealthChecker(healthcheck.Config{
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
