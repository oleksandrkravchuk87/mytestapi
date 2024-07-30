package db

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

// NewDBClient returns a pointer to db client
func NewDBClient(connection string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", connection)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error verifying connection: %v", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
