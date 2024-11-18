package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World\n")
	}

	http.HandleFunc("/hello", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
