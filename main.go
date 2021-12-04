package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DSN string `required:true`
}

var pool *sql.DB // Database connection pool.

func main() {
	var c Config
	fmt.Println("vim-go")
	err := envconfig.Process("myapp", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	pool, err = sql.Open("driver-name", c.DSN)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()

}
