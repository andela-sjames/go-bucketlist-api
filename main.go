package main

import (
	"log"
	"net/http"
	"time"

	"github.com/andela-sjames/go-bucketlist-api/auth"
	"github.com/andela-sjames/go-bucketlist-api/views"
	"github.com/gorilla/mux"
)

func main() {
	// Init Router
	router := mux.NewRouter()
	userSubRoutes := router.PathPrefix("/api/auth").Subrouter()
	bucketlistSubRoutes := router.PathPrefix("/api/bucketlist").Subrouter()
	itemSubRoutes := router.PathPrefix("/api/bucketlist/{id:[0-9]+}/items").Subrouter()

	// attach JWT auth middleware
	router.Use(auth.JWTAuthenticationMiddleware)

	// Route Handlers / Endpoints
	router.HandleFunc("/", views.HomeHandler)

	// Define API sub routes
	userSubRoutes.HandleFunc("/signup", views.CreateUserHandler).Methods("POST")
	userSubRoutes.HandleFunc("/login", views.AuthenticateHandler).Methods("POST")

	bucketlistSubRoutes.HandleFunc("/", views.CreateBucketlistHandler).Methods("POST")
	bucketlistSubRoutes.HandleFunc("/", views.GetAllBucketlistHandler).Methods("GET")

	bucketlistSubRoutes.HandleFunc("/{id:[0-9]+}", views.GetBucketByIDlistHandler).Methods("GET")
	bucketlistSubRoutes.HandleFunc("/{id:[0-9]+}", views.UpdateBucketByIDlistHandler).Methods("PUT")
	// bucketlistSubRoutes.HandleFunc("/{id:[0-9]+}", views.DeleteBucketByIDlistHandler).Methods("DELETE")

	itemSubRoutes.HandleFunc("/", views.CreateItemHandler).Methods("POST")
	itemSubRoutes.HandleFunc("/{itemID:[0-9]+}", views.UpdateDeleteItemHandler).Methods("PUT", "DELETE")

	// TODO: Paginate the get routes.

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
