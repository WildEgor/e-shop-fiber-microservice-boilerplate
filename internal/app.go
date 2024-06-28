package pkg

import (
	"context"
	"fmt"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/configs"
	eh "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/errors"
	nfm "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/middlewares/not_found"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/routers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/google/wire"
	"github.com/grafana/pyroscope-go"
	"log/slog"
	"time"
)

// AppSet link main app deps
var AppSet = wire.NewSet(
	NewApp,
	configs.Set,
	routers.Set,
)

// Server represents the main server configuration.
type Server struct {
	App            *fiber.App
	Configurator   *configs.Configurator
	AppConfig      *configs.AppConfig
	ProfilerConfig *configs.ProfilerConfig
	Pyro           *pyroscope.Profiler
}

// Run start service with deps
func (srv *Server) Run(ctx context.Context) {
	slog.Info("server is listening")

	go func() {
		slog.Info("watch config")
		select {
		default:
			srv.Configurator.Watch()
		case <-ctx.Done():
			slog.Info("stop watch config")
			return
		}
	}()

	pr := func(ac *configs.AppConfig) {
		slog.Info("changer called")

		if ac.IsDebug() && srv.Pyro == nil {
			slog.Info("pyro start")
			pyro, err := pyroscope.Start(pyroscope.Config{
				ApplicationName: srv.AppConfig.Name,
				ServerAddress:   srv.ProfilerConfig.API,
				Logger:          pyroscope.StandardLogger,
				ProfileTypes: []pyroscope.ProfileType{
					pyroscope.ProfileCPU,
					pyroscope.ProfileAllocObjects,
					pyroscope.ProfileAllocSpace,
					pyroscope.ProfileInuseObjects,
					pyroscope.ProfileInuseSpace,
				},
			})

			if err != nil {
				slog.Error("cannot start pyro client", slog.Any("err", err))
			}

			srv.Pyro = pyro
		}

		if !ac.IsDebug() {
			slog.Info("pyro stop")
			if srv.Pyro != nil {
				err := srv.Pyro.Stop()
				if err != nil {
					slog.Warn("cannot stop pyro client", slog.Any("err", err))
					return
				}
			}
		}

	}

	srv.AppConfig.OnChanged(pr)

	if err := srv.App.Listen(fmt.Sprintf(":%s", srv.AppConfig.HTTPPort), fiber.ListenConfig{
		DisableStartupMessage: false,
		EnablePrintRoutes:     false,
		OnShutdownSuccess: func() {
			slog.Debug("success shutdown service")
		},
	}); err != nil {
		slog.Error("unable to start server")
	}
}

// Shutdown graceful shutdown
func (srv *Server) Shutdown(ctx context.Context) {
	slog.Info("shutdown service")

	if srv.Pyro != nil {
		srv.Pyro.Stop()
	}

	if err := srv.App.Shutdown(); err != nil {
		slog.Error("unable to shutdown server")
	}
}

// NewApp init app
func NewApp(
	ac *configs.AppConfig,
	lc *configs.LoggerConfig, // init logger
	pc *configs.ProfilerConfig,
	c *configs.Configurator,
	eh *eh.ErrorsHandler,
	prr *routers.PrivateRouter,
	pbr *routers.PublicRouter,
	sr *routers.SwaggerRouter,
) *Server {
	app := fiber.New(fiber.Config{
		AppName:      ac.Name,
		ErrorHandler: eh.Handle,
		Views:        html.New("./assets", ".html"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Authorization, Connection, Access-Control-Allow-Origin",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))
	app.Use(recover.New())

	prr.Setup(app)
	pbr.Setup(app)
	sr.Setup(app)

	// 404 handler
	app.Use(nfm.NewNotFound())

	return &Server{
		App:            app,
		AppConfig:      ac,
		ProfilerConfig: pc,
		Configurator:   c,
	}
}
