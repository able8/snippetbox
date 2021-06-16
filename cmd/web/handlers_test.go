package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {

	// Create a new instance of our application struct. For now, this just
	// contains a couple of mock loggers (which discard anything writtern to them)
	app := &application{
		errorLog: log.New(ioutil.Discard, "", 0),
		infoLog:  log.New(ioutil.Discard, "", 0),
	}

	// This starts up a HTTPS server which listens on a randomly-chosen port.
	// Notice that we defer a call to ts.Close() to shut down the server when the test finishes.
	ts := httptest.NewTLSServer(app.routes())
	defer ts.Close()

	// The newwork address is contained in the ts.URL field.
	// This returns a http.Response struct containing the response.
	rs, err := ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	// And we can check that the response body writtern by the ping handler equals "ok".
	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}
