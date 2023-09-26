package main

import (
	"./api" // Import the api package using a relative path
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware
	r := gin.Default()

	// Call the SetupRoutes function from the api package to set up logger routes
	api.SetupRoutes(r)

	// Start the HTTP server and listen on port 8080
	r.Run(":8080")
}

/*package main

import (
	"fmt"
	"github.com/angel-one/go-example-project/api"
	"github.com/angel-one/go-example-project/business"
	"github.com/angel-one/go-example-project/constants"
	"github.com/angel-one/go-example-project/jobs"
	"github.com/angel-one/go-example-project/utils/configs"
	"github.com/angel-one/go-example-project/utils/database"
	"github.com/angel-one/go-example-project/utils/flags"
	"github.com/angel-one/go-example-project/utils/httpclient"
	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/go-utils/middlewares"
	"time"

	_ "github.com/angel-one/go-example-project/docs"
	_ "github.com/lib/pq"
)

// @title Go Example Project
// @version 1.0
// @description Go Example Project
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description "Bearer" followed by a space and JWT token (Specify Trade vs Non Trade JWt token).
// @contact.name Add Team Name
// @contact.email Team DL or Slack Group Name

// @BasePath /

func main() {
	initConfigs()
	startLogger()
	initHTTPClient()
	initDatabase()
	//initAsyncModules()
	defer closeDatabase()
	initHealthCheck()
	startRouter()
}

func initHealthCheck() {

}

func initConfigs() {
	// init configs
	configs.Init(flags.BaseConfigPath())
}

func startLogger() {
	// start logger
	loggerConfig, err := configs.Get(constants.LoggerConfig)
	if err != nil {
		log.Fatal(nil).Err(err).Msg("error getting logger config")
	}
	log.InitLogger(log.Level(loggerConfig.GetString(constants.LogLevelConfigKey)))
}

func initHTTPClient() {
	// get application configs
	applicationConfig, err := configs.Get(constants.ApplicationConfig)
	if err != nil {
		log.Fatal(nil).Err(err).Msg("error getting application config")
	}

	// init http client
	err = httpclient.Init(httpclient.Config{
		ConnectTimeout: time.Millisecond *
			applicationConfig.GetDuration(constants.HTTPConnectTimeoutInMillisKey),
		KeepAliveDuration: time.Millisecond *
			applicationConfig.GetDuration(constants.HTTPKeepAliveDurationInMillisKey),
		MaxIdleConnections: applicationConfig.GetInt(constants.HTTPMaxIdleConnectionsKey),
		IdleConnectionTimeout: time.Millisecond *
			applicationConfig.GetDuration(constants.HTTPIdleConnectionTimeoutInMillisKey),
		TLSHandshakeTimeout: time.Millisecond *
			applicationConfig.GetDuration(constants.HTTPTlsHandshakeTimeoutInMillisKey),
		ExpectContinueTimeout: time.Millisecond *
			applicationConfig.GetDuration(constants.HTTPExpectContinueTimeoutInMillisKey),
		Timeout: time.Millisecond *
			applicationConfig.GetDuration(constants.HTTPTimeoutInMillisKey),
	})
	if err != nil {
		log.Fatal(nil).Err(err).Msg("unable to initialize http client")
	}
}

func initDatabase() {
	// init database
	databaseConfig, err := configs.Get(constants.DatabaseConfig)
	if err != nil {
		log.Fatal(nil).Err(err).Msg("error getting database config")
	}
	err = database.InitDatabase(database.Config{
		//Server:                databaseConfig.GetString(constants.DatabaseServerConfigKey),
		Url:                  databaseConfig.GetString(constants.DatabaseUrlConfigKey),
		Name:                  databaseConfig.GetString(constants.DatabaseNameConfigKey),
		Username:              databaseConfig.GetString(constants.DatabaseUsernameConfigKey),
		Password:              databaseConfig.GetString(constants.DatabasePasswordConfigKey),
		MaxOpenConnections:    databaseConfig.GetInt(constants.DatabaseMaxOpenConnectionsKey),
		MaxIdleConnections:    databaseConfig.GetInt(constants.DatabaseMaxIdleConnectionsKey),
		ConnectionMaxLifetime: databaseConfig.GetDuration(constants.DatabaseConnectionMaxLifetimeInSecondsKey) * time.Second,
		ConnectionMaxIdleTime: databaseConfig.GetDuration(constants.DatabaseConnectionMaxIdleTimeInSecondsKey) * time.Second,
	})
	if err != nil {
		log.Fatal(nil).Err(err).Msg("unable to initialize database")
	}
}

func closeDatabase() {
	err := database.Close()
	if err != nil {
		log.Fatal(nil).Err(err).Msg("error closing database")
	}
}

func initAsyncModules() {

	jobsConfig, err := configs.Get(constants.JobsConfig)
	if err != nil {
		log.Fatal(nil).Err(err).Msg("error getting jobs config")
	}

	jobController := jobs.NewJobController(
		jobsConfig.GetString(constants.RedisConnectionString),
		jobsConfig.GetUint(constants.JobsTTLInHrs),
		jobsConfig.GetInt(constants.NumberOfWorkers))

	// Now pass this to every API module which needs jobs
	business.InitMaths(jobController)
	//.. other modules which need jobs here

	// At last spawn the worker
	go jobController.Start()

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
*/
