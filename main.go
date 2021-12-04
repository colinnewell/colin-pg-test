package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type Config struct {
	DSN string `required:"true"`
}

var pool *sql.DB // Database connection pool.

// MYAPP_DSN="dbname=docker user=docker password=password host=postgres sslmode=disable"

func main() {
	var c Config
	err := envconfig.Process("myapp", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	pool, err = sql.Open("postgres", c.DSN)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()

	if err := pool.PingContext(context.Background()); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	fmt.Println("vim-go")
}
