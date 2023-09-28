package api

import (
	"github.com/arideep07/angelOneTest/constants"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// GetRouter is used to get the router configured with the middlewares and the routes
func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())

	// configure swagger
	router.GET(constants.SwaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))

	// configure actuator
	router.GET(constants.ActuatorRoute, actuator)

	// Configure Logger routes
	SetupLoggerRoutes(router)

	return router
}
