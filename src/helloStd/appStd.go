package helloStd

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

var (
	Address = "localhost:8080"
	Phrase  = "Hello, Gopher! Not using gcloud ?"
)

/* init() is required as there is no main and appengine.Main() is not called */
func init() {
	http.HandleFunc("/", Hello)
}

/* Hander name is exported for test purposes. Import like "../" are not supported by dev_appserver.py*/
func Hello(w http.ResponseWriter, r *http.Request) {
	// log.Println("request on", r.RequestURI)
	// Not printing version during tests
	if wType := reflect.TypeOf(w).String(); wType != "*httptest.ResponseRecorder" &&
		wType != "*http.response" {
		log.Println("running", runtime.Version())
	}
	fmt.Fprint(w, Phrase)
}
