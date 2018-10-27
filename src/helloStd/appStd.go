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
	Phrase  = "Hello, Gopher! Also on gcloud ?"
)

/* init() is required as there is no main and appengine.Main() cannot be called */
func init() {
	http.HandleFunc("/", Hello)
}

/* Hander name is exported for test purposes. Import like "../" is not supported by dev_appserver.py*/
func Hello(w http.ResponseWriter, r *http.Request) {
	// log.Println("request on", r.RequestURI)
	// Not printing version during tests. Type varies with test method.
	if wType := reflect.TypeOf(w).String(); wType != "*httptest.ResponseRecorder" {
		log.Println("running", runtime.Version())
	}
	fmt.Fprint(w, Phrase)
}
