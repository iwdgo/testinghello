package main

// Unit test file resides in the directory of the tested code.
import (
	// "./helloFlex" is not supported as it is package main
	"./helloStd"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var url = "http://" + helloStd.Address

const helloFlexPhrase = "Hello, Gopher! Not using gcloud ?"

// When the site is online
func TestHelloOnlineNoClient(t *testing.T) {
	r, err := http.Get(url)
	if err != nil {
		t.Fatal("hello is not online\n", err)
		// hello is not online:
		// Get http://localhost:8080: dial tcp [::1]:8080:
		// connectex: No connection could be made because the target machine actively refused it.
	}
	if r.StatusCode != 200 {
		t.Fatal("hello returned error code", r.StatusCode)
	}
	// Cannot distinguish deploy method
	if got, err := ioutil.ReadAll(r.Body); err != nil {
		t.Fatal("reading body failed with ", err)
	} else if string(got) != helloStd.Phrase && string(got) != helloFlexPhrase {
		t.Fatalf("fail : got %q, want %q", got, helloFlexPhrase+" or "+helloStd.Phrase)
	}
}

func TestHelloOnlineClientGet(t *testing.T) {
	client := &http.Client{}

	r, err := client.Get(url) // "/" is the same page
	if err != nil {
		t.Fatal(url, "is unavailable: ", err)
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		t.Fatal("client get is not OK: ", r.StatusCode)
	}

	got := make([]byte, r.ContentLength)
	b, err := r.Body.Read(got)
	if int64(b) != r.ContentLength {
		t.Fatal("data lost")
	}
	if err != io.EOF {
		t.Fatal("error reading body: ", err, "and read", b)
	} else if string(got) != helloStd.Phrase && string(got) != helloFlexPhrase {
		t.Fatalf("fail : got %q, want %q", got, helloFlexPhrase+" or "+helloStd.Phrase)
	}
}

func TestHelloOnlineClientDo(t *testing.T) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, http.NoBody) // nil panics deferred close
	if err != nil {
		t.Fatal(err)
	}
	defer req.Body.Close()

	r, err := client.Do(req)
	if err != nil {
		t.Fatal("client do:", err)
	}

	got := make([]byte, r.ContentLength)
	b, err := r.Body.Read(got)
	if int64(b) != r.ContentLength {
		t.Fatal("data lost: ", r.ContentLength-int64(b))
	}
	if err != io.EOF {
		t.Fatal("error reading body: ", err, "and read", b)
	} else if string(got) != helloStd.Phrase && string(got) != helloFlexPhrase {
		t.Fatalf("fail : got %q, want %q", got, helloFlexPhrase+" or "+helloStd.Phrase)
	}
}

// When the site is offline, handler is called directly
func TestHelloHandler(t *testing.T) {
	r, err := http.NewRequest("GET", "", http.NoBody) // url can be "/"
	if err != nil {
		t.Fatal("New request failed with ", err)
	}
	defer r.Body.Close()

	w := httptest.NewRecorder() // to record the transaction
	helloStd.Hello(w, r)

	if w.Code != 200 {
		t.Fatalf("request failed with code: %d", w.Code)
	}

	got := w.Body.String()
	if want := fmt.Sprintf(helloStd.Phrase); got != want {
		t.Fatalf("wrong body returned: got %s, want %s", got, want)
	}
}

// Using httptest, url may not be empty
func TestHelloHandler2(t *testing.T) {
	r := httptest.NewRequest("GET", "/", http.NoBody)
	defer r.Body.Close()

	w := httptest.NewRecorder()
	helloStd.Hello(w, r) // helloFlex test is identical

	if w.Code != 200 {
		t.Fatalf("request failed with code: %d", w.Code)
	}

	got := w.Body.String()
	if want := fmt.Sprintf(helloStd.Phrase); got != want {
		t.Fatalf("wrong body returned: got %s, want %s", got, want)
	}
}
