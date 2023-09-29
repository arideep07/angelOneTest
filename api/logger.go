package api

import (
	"encoding/json"
	"net/http"

	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/nbu-logger-service/constants"
	"github.com/angel-one/nbu-logger-service/models"
	"github.com/gin-gonic/gin"
)

func SetupLoggerRoutes(router *gin.Engine) {
	// Define your logger-related routes here
	router.POST(constants.LoggerRoute, loggerHandler)
}

// SetupRoutes initializes and sets up the routes for the logger API.
func loggerHandler(c *gin.Context) {
	var logEntry models.LogEntry

	// Parse the JSON request body into a LogEntry struct
	if err := c.ShouldBindJSON(&logEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageJson, _ := json.Marshal(logEntry)
	log.Info(c).Msg(string(messageJson))
	// Respond with the logged entry and a status code of 200 (Created)
	c.JSON(http.StatusOK, logEntry)
}
