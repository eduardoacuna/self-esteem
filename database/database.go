package database

import (
	"context"
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/eduardoacuna/self-esteem/log"
)

var connection *sql.DB

func Setup(ctx context.Context, user, dbname, host, port string) error {
	log.Info(ctx, "database setup", "user", user, "dbname", dbname, "host", host, "port", port)
	parameters := fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=disable", user, dbname, host, port)
	var err error
	connection, err = sql.Open("postgres", parameters)
	if err != nil {
		return err
	}
	// Test that the connection has been established
	err = connection.Ping()
	if err != nil {
		return err
	}
	return nil
}

func Connection() *sql.DB {
	return connection
}

func Close(ctx context.Context) {
	log.Warn(ctx, "closing database connection")
	if connection != nil {
		connection.Close()
	}
}
