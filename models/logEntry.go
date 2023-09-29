package models

type LogEntry struct {
	Type string `json:"type" binding:"required"`
	Data map[string]interface{}
}
