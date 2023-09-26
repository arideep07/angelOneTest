package main

import (
	"fmt"

	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/go-utils/middlewares"
	"github.com/arideep07/angelOneTest/api"
	"github.com/arideep07/angelOneTest/utils/flags"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware
	r := gin.Default()

	// Call the SetupRoutes function from the api package to set up logger routes
	api.SetupRoutes(r)

	// Start the HTTP server and listen on port
	startRouter()
}

func startRouter() {
	// get router
	router := api.GetRouter(middlewares.Logger(middlewares.LoggerMiddlewareOptions{}))
	// now start router
	err := router.Run(fmt.Sprintf(":%d", flags.Port()))
	if err != nil {
		log.Fatal(nil).Err(err).Msg("error starting router")
	}
}
