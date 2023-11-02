package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

// While starting, the server is actively refusing connection.
const connectError = "connectex: No connection could be made because the target machine actively refused it."

var (
	target     *string
	client     = &http.Client{}
	wantPhrase = phrase()
)

func TestMain(m *testing.M) {
	target = flag.String("target", "localhost", "Set url of deployed application")
	flag.Parse()
	// Override port default using PORT environment variable
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	if *target == "localhost" {
		go func() {
			startServer()
		}()
		var err error
		d := 100 * time.Millisecond
		for {
			_, err = http.Get(getAppURL("").String())
			if err == nil {
				break
			}
			if !strings.Contains(err.Error(), connectError) {
				log.Fatalln(err)
			}
			log.Printf("waiting %s for %s:%s to start up\n", d, *target, port)
			time.Sleep(d)
		}
	}
	os.Exit(m.Run())
}

// getAppURL() returns a string containing the URL of the app as needed by http.
// The parameter s is the path
func getAppURL(s string) *url.URL {
	return &url.URL{Scheme: "http", Host: fmt.Sprintf("%s:%s", *target, port), Path: s}
}

// When the site is online, various methods to execute the request:
// - Using Get without client
// - Using Get of the client
// - Using Do of the client

func TestHelloOnlineNoClient(t *testing.T) {
	r, err := http.Get(getAppURL("").String())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			t.Log(err)
		}
	}()

	if r.StatusCode != 200 {
		t.Fatal("request failed:", r.StatusCode)
	}

	got, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal("reading body failed:", err)
	}

	if !bytes.Equal(got, wantPhrase) {
		t.Fatalf("got %s, want %s", got, wantPhrase)
	}
}

func TestHelloOnlineClientGet(t *testing.T) {
	r, err := client.Get(getAppURL("").String())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			t.Log(err)
		}
	}()

	if r.StatusCode != 200 {
		t.Fatalf("%s", r.Status)
	}

	got := make([]byte, r.ContentLength)
	b, err := r.Body.Read(got)
	if int64(b) != r.ContentLength {
		t.Fatalf("%d bytes lost", int64(b)-r.ContentLength)
	}
	if err != io.EOF {
		t.Fatalf("error reading body: %v and read %d", err, b)
	}

	if !bytes.Equal(got, wantPhrase) {
		t.Fatalf("got %s, want %s", got, wantPhrase)
	}
}

func TestHelloOnlineClientDo(t *testing.T) {
	req, err := http.NewRequest("GET", getAppURL("").String(), http.NoBody)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = req.Body.Close()
		if err != nil {
			t.Log(err)
		}
	}()

	r, err := client.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), connectError) {
			t.Skip("no connection")
		}
		t.Fatal(err)
	}

	// Reading the response
	got := make([]byte, r.ContentLength)
	b, err := r.Body.Read(got)
	if int64(b) != r.ContentLength {
		t.Fatal("data lost: ", r.ContentLength-int64(b))
	}
	if err != io.EOF {
		t.Fatalf("error reading body: %v and read %d", err, b)
	}

	if !bytes.Equal(got, wantPhrase) {
		t.Fatalf("got %s, want %s", got, wantPhrase)
	}
}

// When the site is offline, handler is called directly and request is built using httptest.
func TestHelloHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", getAppURL("").String(), http.NoBody)
	defer func() {
		err := r.Body.Close()
		if err != nil {
			t.Log(err)
		}
	}()

	hello(w, r)

	if w.Code != 200 {
		t.Fatalf("request failed with code: %d", w.Code)
	}

	if got := w.Body.Bytes(); !bytes.Equal(got, wantPhrase) {
		t.Fatalf("got %s, want %s", w.Body.String(), string(wantPhrase))
	}
}
