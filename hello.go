// Package main starts an elementary server with one handler when invoked.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = "8080"

func phrase() []byte {
	b := []byte("Hello, Gopher! Are you online ?")
	// If deployed on Google Cloud, append some configuration variables
	if s := os.Getenv("GAE_SERVICE"); s != "" {
		b = append(b, fmt.Sprintf("\n%s\n%s", s, os.Getenv("GAE_INSTANCE"))...)
	}
	return b
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, string(phrase()))
}

func startServer() {
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	log.Printf("Listening on port %s", port)
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	startServer()
}
