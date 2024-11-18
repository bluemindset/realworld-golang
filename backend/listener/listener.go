package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	url := "postgres://stef:stefpass@localhost:5437/realworld"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stderr, "Successfully connecrted")
	}
	var id int
	var name string
	var email string

	err1 := conn.QueryRow(context.Background(), "SELECT id, username, email FROM users WHERE id = $1", 0).Scan(&id, &name, &email)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err1)
		return
	}
	fmt.Printf("User: ID=%d, Name=%s, Email=%s\n", id, name, email)

}
