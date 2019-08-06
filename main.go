package main

import (
	"log"
	"net/http"
	"time"

	"github.com/andela-sjames/go-bucketlist-api/auth"
	"github.com/andela-sjames/go-bucketlist-api/controllers"
	"github.com/gorilla/mux"
)

func main() {
	// Init Router
	router := mux.NewRouter()

	router.Use(auth.JWTAuthenticationMiddleware) //attach JWT auth middleware

	// Route Handlers / Endpoints
	router.HandleFunc("/", controllers.HomeHandler)

	// server block defined here
	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8000",

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server on port :8000")
	log.Fatal(srv.ListenAndServe())

}
