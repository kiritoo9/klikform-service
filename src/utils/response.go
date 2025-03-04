package utils

import (
	"encoding/json"
	"net/http"
)

func SetResponse(w http.ResponseWriter, statusCode int, message string, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]any{
		"message": message,
		"data":    data,
	})
}
