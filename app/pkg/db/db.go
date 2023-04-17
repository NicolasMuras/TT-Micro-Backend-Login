package db

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {

	username := "postgres"
	password := "postgres"
	hostname := "core-testing.c4kgjhp0r4l4.us-east-2.rds.amazonaws.com"
	port := "5432"
	database := "postgres"

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, hostname, port, database)
	client, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error: Database Connection failed : %v", err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
