package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	addr := ":5001"
	handler := http.HandlerFunc(MyGreeterHandler)

	fmt.Printf("Starting server on %s\n", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	log.Fatal(server.ListenAndServe())
}
