package controllers

import (
	"encoding/json"
	"net/http"
)

// HomeHandler handles request to the base path and
// return a simple welcome text
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "200",
		"message": "Welcome to Bucketlist API",
		"signup":  "/api/user/signup",
		"login":   "/api/user/login",
	})
}
