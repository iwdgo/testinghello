// Package main starts an elementary server with one handler when invoked.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	port = "8080"
)

func phrase() string {
	// TODO Phrase will differ from online and offline. Some configuration is useful
	return fmt.Sprintf("Hello, Gopher! Are you online ?\n%s\n%s\n", os.Getenv("GAE_SERVICE"), os.Getenv("GAE_INSTANCE"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, phrase())
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
