package pkg

import (
	"context"
	"fmt"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/config"
	eh "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/errors"
	nfm "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/middlewares/not_found"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/router"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/template/html/v2"
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

func (srv *Server) Run(ctx *context.Context) {
	slog.Info("server is listening")

	if err := srv.App.Listen(fmt.Sprintf(":%s", srv.AppConfig.Port)); err != nil {
		slog.Error("unable to start server")
	}
}

func (srv *Server) Shutdown() {
	slog.Info("shutdown service")
	if err := srv.App.Shutdown(); err != nil {
		slog.Error("unable to shutdown server")
	}
}

func NewApp(
	ac *config.AppConfig,
	eh *eh.ErrorsHandler,
	prr *router.PrivateRouter,
	pbr *router.PublicRouter,
	sr *router.SwaggerRouter,
) *Server {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	if ac.IsProduction() {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}
	slog.SetDefault(logger)

	app := fiber.New(fiber.Config{
		ErrorHandler: eh.Handle,
		Views:        html.New("./views", ".html"),
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Use(recover.New())

	prr.Setup(app)
	pbr.Setup(app)
	sr.Setup(app)

	// 404 handler
	app.Use(nfm.NewNotFound())

	return &Server{
		App:       app,
		AppConfig: ac,
	}
}
