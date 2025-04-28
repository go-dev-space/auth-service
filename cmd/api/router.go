package main

import (
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/auth-service/internal/auth/interfaces"
	"github.com/auth-service/pkg/logs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type system struct {
	config    *config
	handler   *handler
	waitgroup *sync.WaitGroup
	logger    *logs.Logwriter
}

type config struct {
	addr         string
	accessHeader string
	swagger      string
}

type handler struct {
	healthcheck  *interfaces.HealthcheckHandler
	registration *interfaces.RegistrationHandler
}

// mount returns an http.handler for server startup
func (app system) mount() http.Handler {

	r := chi.NewRouter()

	// middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// serve swagger.json static
	r.Handle("/swagger.json", http.FileServer(http.Dir("./docs")))

	r.Route("/v1", func(r chi.Router) {
		// swagger route
		r.Get("/swagger-ui/*", httpSwagger.Handler(
			httpSwagger.URL(app.config.swagger),
		))

		// mount the auth subrouter
		r.Mount("/auth", interfaces.NewRouter(
			*app.handler.healthcheck,
			*app.handler.registration,
			app.accessHeader,
		))
	})

	return r
}

// run starts the webserver, it needs an http.handler and returns an error
func (app system) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		IdleTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	}
	return srv.ListenAndServe()
}

// shutdown for a graceful application shutdown
func (app system) shutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGABRT, syscall.SIGTERM)
	<-quit
	app.cleanup()
	os.Exit(0)
}

// cleanup all tasks before shutdown
func (app system) cleanup() {
	app.logger.Info.Println("cleanup started...")
	app.waitgroup.Wait()
	app.logger.Info.Println("cleanup done")
}
