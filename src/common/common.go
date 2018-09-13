package common

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

/* Used only by test */
func Hello(w http.ResponseWriter, r *http.Request) {
	// log.Println("request on", r.RequestURI)
	// Not printing version during tests when we use httptest recorder
	if reflect.TypeOf(w).String() == "*httptest.ResponseRecorder" {
		log.Println("running", runtime.Version())
	}
	fmt.Fprint(w, Phrase)
}
