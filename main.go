package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/andela-sjames/go-bucketlist-api/auth"
	"github.com/gorilla/mux"
)

// HomeHandler handles request to the base path and
// return a simple welcome text
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "200",
		"message": "Welcome to Bucketlist API",
	})
}

func main() {
	// Init Router
	router := mux.NewRouter()

	router.Use(auth.JWTAuthenticationMiddleware) //attach JWT auth middleware

	// Route Handlers / Endpoints
	router.HandleFunc("/", HomeHandler)

	// server block defined here
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server on port :8000")
	log.Fatal(srv.ListenAndServe())

}
