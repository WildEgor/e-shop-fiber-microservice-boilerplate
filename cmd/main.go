package main

import (
	"context"
	"os/signal"
	"syscall"

	server "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal"
)

// @title		Swagger Doc
// @version		1.0
// @description	App
// @termsOfService	/
// @contact.name	mail
// @contact.url	/
// @contact.email	kartashov_egor96@mail.ru
// @license.name	MIT
// @license.url	http://www.apache.org/licenses/MIT.html
// @host			localhost:8888
// @BasePath		/
// @schemes		http
func main() {
	ctx, done := signal.NotifyContext(context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer done()

	srv, _ := server.NewServer()
	srv.Run(&ctx)

	<-ctx.Done()
	srv.Shutdown()
}
