package api

import (
	"net/http"
	"time"

	"github.com/arideep07/angelOneTest/constants"
	"github.com/gin-gonic/gin"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

func SetupLoggerRoutes(router *gin.Engine) {
	// Define your logger-related routes here
	router.POST(constants.LoggerRoute, loggerHandler)
}

var logs []LogEntry

// SetupRoutes initializes and sets up the routes for the logger API.
func loggerHandler(c *gin.Context) {
	var logEntry LogEntry

	// Parse the JSON request body into a LogEntry struct
	if err := c.ShouldBindJSON(&logEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add a timestamp to the log entry
	logEntry.Timestamp = time.Now()

	// Append the log entry to the in-memory log slice
	logs = append(logs, logEntry)

	// Respond with the logged entry and a status code of 201 (Created)
	c.JSON(http.StatusCreated, logEntry)
}
