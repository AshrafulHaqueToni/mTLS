package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// set up handler to listen to root path
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("new request")
		fmt.Println(w, "hello world")
	})

	// server on port 9090 of local host
	server := http.Server{
		Addr:    ":9090",
		Handler: handler,
	}

	if err := server.ListenAndServeTLS("/home/appscodepc/go/src/github.com/AshrafulHaqueToni/mTLS/certs/server.crt",
		"/home/appscodepc/go/src/github.com/AshrafulHaqueToni/mTLS/certs/server.key"); err != nil {
		log.Fatalf("error listening to server: %v", err)
	}
}
