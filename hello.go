package main

import (
	"fmt"
	"net/http"
)

// If local, go run hello.go
// You can test your gcloud sdk running dev_appserver.py ./hello/app.yam
const (
	address = "localhost:8080"
	url     = "http://" + address
	phrase  = "Hello, Gopher! Not using gcloud ?"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// log.Println("request on", r.RequestURI)
	fmt.Fprint(w, phrase)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(address, nil)
}
