package helloStd

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

var (
	Phrase = "Hello, Gopher! Not using gcloud ?"
)

// If local, go run hello.go
// Usage of init is required for dev_appserver.py
func init() {
	http.HandleFunc("/", hello)
}

// same handler in all packages
func hello(w http.ResponseWriter, r *http.Request) {
	// log.Println("request on", r.RequestURI)
	log.Println("running", runtime.Version())
	fmt.Fprint(w, Phrase)
}
