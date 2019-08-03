package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	port   = 5432
	user   = "postgres"
	dbname = "postgres"
)

// User field (Model) defined
type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastNames string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Bucketlist field (Model) defined
type Bucketlist struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Created  string `json:"date_created"`
	Modified string `json:"date_modified"`
	User     *User  `json:"user"`
}

// BucketlistItem field (Model) defined
type BucketlistItem struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Created    string      `json:"date_created"`
	Modified   string      `json:"date_modified"`
	Done       bool        `json:"done"`
	Bucketlist *Bucketlist `json:"bucketlist"`
}

// HomeHandler handles request to the base path and
// return a simple welcome text
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "200",
		"message": "Welcome to Bucketlist API",
	})
}

func main() {

	password := os.Getenv("PGPASSWORD")
	host := os.Getenv("PGHOST")

	// setup DB connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to DB!")

	// Init Router
	router := mux.NewRouter()

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
