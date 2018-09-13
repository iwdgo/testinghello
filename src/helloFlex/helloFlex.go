package main

// Package is main as required by flexible usage of GAE
import (
	"fmt"
	"google.golang.org/appengine"
	"log"
	"net/http"
	"runtime"
)

var (
	Address = "localhost:8080"
	Phrase  = "Hello, Gopher! Not using gcloud ?"
)

// If local, go run hello.go
// You can test your gcloud sdk running dev_appserver.py appFlex.yaml
//  or dev_appserver.py .\helloFlex\appFlex.yaml
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

// Main is here as flex requires it
func main() {
	if appengine.IsDevAppServer() {
		log.Println("running dev_appserver.py")
		// Not working to start. Still fails during deployment
		appengine.Main()
	} else {
		http.ListenAndServe(Address, nil)
	}
}
