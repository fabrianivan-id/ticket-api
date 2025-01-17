package utils

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sqlx.DB

func InitDB() {
	var err error
	DB, err = sqlx.Open("postgres", "user=postgres dbname=support_system sslmode=disable")
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}
}
