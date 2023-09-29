package main

import (
	"context"
	"fmt"

	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/go-utils/middlewares"
	"github.com/angel-one/nbu-logger-service/api"
	"github.com/angel-one/nbu-logger-service/constants"
	"github.com/angel-one/nbu-logger-service/utils/flags"
)

func main() {
	//set up logger
	startLogger()
	// Start the HTTP server and listen on port
	startRouter()
}

func startLogger() {
	log.InitLogger(log.Level(constants.InfoLevel))
}

func startRouter() {
	ctx := context.Background()
	// get router
	router := api.GetRouter(middlewares.Logger(middlewares.LoggerMiddlewareOptions{}))
	// now start router
	err := router.Run(fmt.Sprintf(":%d", flags.Port()))
	if err != nil {
		log.Fatal(ctx).Err(err).Msg("error starting router")
	}
}
