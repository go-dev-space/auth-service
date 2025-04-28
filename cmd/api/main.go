package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/auth-service/internal/auth/application"
	"github.com/auth-service/internal/auth/infrastructure/database"
	"github.com/auth-service/internal/auth/infrastructure/store"
	"github.com/auth-service/internal/auth/interfaces"
	"github.com/auth-service/pkg/logs"
	"github.com/joho/godotenv"
)

//	@title			go-dev-space - [auth-service]
//	@version		1.0.6
//	@description	An example of an authorization microservice in Go according to DDD with EDA in a kubernetes cluster..
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	MIT License
//	@license.url	https://github.com/go-dev-space/auth-service?tab=MIT-1-ov-file

//	@host		localhost:8080
//	@BasePath	/v1

// @securityDefinitions.basic	BasicAuth
// @in							header
// @name						Authorization
// @description
func main() {

	// load env file
	if err := godotenv.Load("./.env"); err != nil {
		if err := godotenv.Load("./app/.env"); err != nil {
			panic(err)
		}
	}

	// initialize logger
	logger := logs.New()

	// database
	db, err := database.NewPostgresql(os.Getenv("POSTGRES_URI"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// store
	store := store.NewStoreServicePostrgresql(db)

	// use cases
	healthcheckUC := application.NewHealthcheckUseCase()
	registrationUC := application.NewRegistrationUserUseCase(store, logger)

	// system instance
	app := &system{
		config: &config{
			addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
			accessHeader: os.Getenv("ACCESS_HEADER"),
			swagger:      os.Getenv("SWAGGER_PATH"),
		},
		handler: &handler{
			healthcheck:  interfaces.NewHealthcheckHandler(logger, healthcheckUC),
			registration: interfaces.NewRegistrationHandler(logger, registrationUC),
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
