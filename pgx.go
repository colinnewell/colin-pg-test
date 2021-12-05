package main

import (
	"context"
	"fmt"
	"os"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("MYAPP_DSN"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

}
