package controllers

import (
	"encoding/json"
	"net/http"
)

// HomeHandler handles request to the base path and
// return a simple welcome text
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":        "200",
		"message":       "Welcome to Bucketlist API",
		"documentation": "documentation_link_goes_here",
	})
}
