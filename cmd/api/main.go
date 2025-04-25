package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/auth-service/internal/auth/application"
	"github.com/auth-service/internal/auth/interfaces"
	"github.com/auth-service/pkg/logs"
	"github.com/joho/godotenv"
)

func main() {

	// load env file
	if err := godotenv.Load("./.env"); err != nil {
		if err := godotenv.Load("./app/.env"); err != nil {
			panic(err)
		}
	}

	// initialize logger
	logger := logs.New()

	// use cases
	healthcheckUC := application.NewHealthcheckUseCase()

	// system instance
	app := &system{
		config: &config{
			addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
			accessHeader: os.Getenv("ACCESS_HEADER"),
		},
		handler: &handler{
			healthcheck: interfaces.NewHealthcheckHandler(logger, healthcheckUC),
		},
		logger:    logger,
		waitgroup: &sync.WaitGroup{},
	}

	// start service
	app.logger.Info.Printf("service running%s\n", app.config.addr)
	if err := app.run(app.mount()); err != nil {
		panic(err)
	}

	// graceful shutdown
	app.shutdown()
}
