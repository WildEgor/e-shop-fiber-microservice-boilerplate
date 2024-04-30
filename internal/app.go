package pkg

import (
	"fmt"
	"github.com/WildEgor/fibergo-microservice-boilerplate/internal/config"
	eh "github.com/WildEgor/fibergo-microservice-boilerplate/internal/handlers/errors"
	"github.com/WildEgor/fibergo-microservice-boilerplate/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/wire"
	"log/slog"
	"os"
)

var AppSet = wire.NewSet(
	NewApp,
	config.ConfigsSet,
	router.RouterSet,
)

// Server represents the main server configuration.
type Server struct {
	App       *fiber.App
	AppConfig *config.AppConfig
}

func (srv *Server) Run() {
	slog.Info("server is listening")

	if err := srv.App.Listen(fmt.Sprintf(":%s", srv.AppConfig.Port)); err != nil {
		slog.Error("unable to start server", "data", slog.Any("error", err))
	}
}

func (srv *Server) Shutdown() {
	slog.Info("shutdown service")
	if err := srv.App.Shutdown(); err != nil {
		slog.Error("unable to shutdown server", "data", slog.Any("error", err))
	}
}

func NewApp(
	ac *config.AppConfig,
	prr *router.PrivateRouter,
	pbr *router.PublicRouter,
	sr *router.SwaggerRouter,
	eh *eh.ErrorsHandler,
) *Server {

	app := fiber.New(fiber.Config{
		ErrorHandler: eh.Handle,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Use(recover.New())

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	if ac.IsProduction() {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}
	slog.SetDefault(logger)

	prr.SetupPrivateRouter(app)
	pbr.SetupPublicRouter(app)
	sr.SetupSwaggerRouter(app)

	return &Server{
		App:       app,
		AppConfig: ac,
	}
}
