package healthcheck

import (
	"github.com/angel-one/go-example-project/healthcheck/sqlchecker"
	"github.com/angel-one/go-example-project/utils/database"
	"github.com/gin-gonic/gin"
	"github.com/hootsuite/healthchecks"
)


func HealthChecker(ctx *gin.Context) {
	// Define a StatusEndpoint at '/status/db' for a database dependency
	dbEndpoint := healthchecks.StatusEndpoint{
		Name: "The DB",
		Slug: "db",
		Type: "internal",
		IsTraversable: false,
		StatusCheck: sqlchecker.SQLDBStatusChecker{
			DB: database.Get(),
		},
		TraverseCheck: nil,
	}

	// Define the list of StatusEndpoints for your service
	statusEndpoints := []healthchecks.StatusEndpoint{ dbEndpoint }


	// Set the path for the about and version files
	var AboutFilePath = "resources/about.json"
	var VersionFilePath = "resources/version.txt"

	// Set up any service injected customData for /status/about response.
	// Values can be any valid JSON conversion and will override values set in about.json.
	var CustomData = make(map[string]interface{})

	healthCheckHandler := healthchecks.HandlerFunc(statusEndpoints, AboutFilePath, VersionFilePath, CustomData)

	healthCheckHandler(ctx.Writer, ctx.Request)
}
