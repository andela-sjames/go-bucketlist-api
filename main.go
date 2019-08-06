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
	userSubRoutes := router.PathPrefix("/api/user").Subrouter()

	// attach JWT auth middleware
	router.Use(auth.JWTAuthenticationMiddleware)

	// Route Handlers / Endpoints
	router.HandleFunc("/", controllers.HomeHandler)

	// Define API sub routes
	userSubRoutes.HandleFunc("/signup", controllers.CreateUserHandler).Methods("POST")
	userSubRoutes.HandleFunc("/login", controllers.AuthenticateHandler).Methods("POST")

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
