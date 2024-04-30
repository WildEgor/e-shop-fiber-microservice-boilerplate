package main

import (
	"os"
	"os/signal"
	"syscall"

	server "github.com/WildEgor/fibergo-microservice-boilerplate/internal"
	log "github.com/sirupsen/logrus"
)

// @title		Swagger Doc
// @version		1.0
// @description	App
// @termsOfService	/
// @contact.name	mail
// @contact.url	/
// @contact.email	TODO
// @license.name	MIT
// @license.url	http://www.apache.org/licenses/MIT.html
// @host			localhost:8888
// @BasePath		/
// @schemes		http
func main() {
	sigCh := make(chan os.Signal, 1)
	doneCh := make(chan bool, 1)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		log.Printf("[Main] Recieve shutdown signal %s", sig)
		doneCh <- true
	}()

	srv, _ := server.NewServer()
	srv.Run()

	<-doneCh

	srv.Shutdown()
}
