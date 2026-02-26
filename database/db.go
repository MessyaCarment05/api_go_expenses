package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func GetConnection() {
	log.Println("Successfully Call GetConnection Function")
	var err error
	dsn := "root:@tcp(localhost:3306)/api_go_exercises?parseTime=true"
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}
	// cek db bisa pakai

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(50)
	DB.SetConnMaxIdleTime(5 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)
	log.Println("Connected to MySQL")
}
