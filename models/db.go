package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // get the gorm postgres dialect
)

const (
	port = 5432
)

var db *gorm.DB

func init() {

	password := os.Getenv("PGPASSWORD")
	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSER")
	dbname := os.Getenv("PGDBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	conn, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{}, &Bucketlist{}, &BucketlistItem{})
}

// GetDB function defined to return DB instance
func GetDB() *gorm.DB {
	return db
}
