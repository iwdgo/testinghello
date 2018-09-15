package main

import (
	"fmt"
	"google.golang.org/appengine"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

var (
	Address = "localhost:8080"
	Phrase  = "Hello, Gopher! Not using gcloud ?"
)

// same handler in all packages but import ../common is not supported
func hello(w http.ResponseWriter, r *http.Request) {
	// log.Println("request on", r.RequestURI)
	// Not printing os version in testing log when running dev_appserver.py
	if fmt.Sprintf("%s", reflect.TypeOf(w).String()) != "*internal.context" {
		log.Println("running", runtime.Version())
	}
	fmt.Fprint(w, Phrase)
}

// Main is here as flex complains otherwise.
// You have to run Std without main as appengine.Main() fails in std mode
func main() {
	if appengine.IsDevAppServer() {
		if appengine.InstanceID() == "0" { // fails on std
			// running flex
			log.Print("running local dev_appserver.py in flexible environment")
			http.HandleFunc("/", hello)
			appengine.Main()
		} else {
			// dev_appserver Std but currently unusable
		}
	} else {
		// go run ./helloFlex/helloFlex.go
		http.HandleFunc("/", hello)
		http.ListenAndServe(Address, nil)
	}
}
