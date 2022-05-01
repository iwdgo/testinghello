package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	phrase = "Hello, Gopher! Also on Google Cloud ? (w/o billing)"
	port   = "8080"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// Not printing version during tests. Type varies with test method.
	/*
		if wType := reflect.TypeOf(w).String(); wType != "*httptest.ResponseRecorder" &&
			wType != "*http.response" {
			log.Println("running", runtime.Version())
		}
	*/
	fmt.Fprintln(w, phrase)
	fmt.Fprintln(w, os.Getenv("GAE_SERVICE"))
	fmt.Fprintln(w, os.Getenv("GAE_INSTANCE"))
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
