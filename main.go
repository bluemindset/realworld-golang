package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	url := "postgres://stef:stefpass@localhost:5437/realworld"
	_, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stderr, "Successfully connecrted")
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("GET /api/hello/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
