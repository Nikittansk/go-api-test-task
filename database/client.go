package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Nikittansk/go-api-test-task/config"
	_ "github.com/lib/pq"
)

var (
	DB      *sql.DB
	dbError error
)

func Connect(cfg *config.Config) {
	connectingString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Schema)

	DB, dbError = sql.Open("postgres", connectingString)
	if dbError != nil {
		log.Fatal(dbError.Error())
		return
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err.Error())
		return
	}

	DB.SetMaxOpenConns(cfg.MaxConnections)

	log.Println("Connected to PostgreSQL!")
}
