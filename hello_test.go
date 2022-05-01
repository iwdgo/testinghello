package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

// When this word is found, the test is skipped as no connection was found with the server
const connectError = "connection"

var (
	target     = "localhost"
	client     = &http.Client{}
	wantPhrase = phrase
)

func TestMain(m *testing.M) {
	// Override port default using PORT environment variable
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	if target == "localhost" {
		go func() {
			startServer()
		}()
		wantPhrase = phrase + "\n\n\n"
	}
	os.Exit(m.Run())
}

// getAppURL() returns a string containing the URL of the app as needed by http.
// The parameter s is the path
func getAppURL(s string) *url.URL {
	u := url.URL{Scheme: "http", Host: fmt.Sprintf("%s:%s", target, port), Path: s}
	return &u
}

// When the site is online, various methods to execute the request:
// - Using Get without client
// - Using Get of the client
// - Using Do of the client
func TestHelloOnlineNoClient(t *testing.T) {
	r, err := http.Get(getAppURL("").String())
	if err != nil {
		if strings.Contains(err.Error(), connectError) {
			t.Skip(err.Error())
		}
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

	got, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal("reading body failed:", err)
	}

	testPhrase(t, string(got))
}

func TestHelloOnlineClientGet(t *testing.T) {
	r, err := client.Get(getAppURL("").String())
	if err != nil {
		if strings.Contains(err.Error(), connectError) {
			t.Skip(err.Error())
		}
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

	testPhrase(t, string(got))
}

func TestHelloOnlineClientDo(t *testing.T) {
	req, err := http.NewRequest("GET", getAppURL("").String(), http.NoBody)
	if err != nil {
		if strings.Contains(err.Error(), connectError) {
			t.Skip(err.Error())
		}
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

	testPhrase(t, string(got))
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

	testPhrase(t, w.Body.String())
}

func testPhrase(t *testing.T, s string) {
	if got := s; got != wantPhrase {
		t.Fatalf("got %s, want %s", got, wantPhrase)
	}
}
