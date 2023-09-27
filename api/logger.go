package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

var logs []LogEntry

// SetupRoutes initializes and sets up the routes for the logger API.
func logger(r *gin.Engine) {
	// Define an HTTP endpoint to log messages (POST request)
	r.POST("/log", func(c *gin.Context) {
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
	})
}
