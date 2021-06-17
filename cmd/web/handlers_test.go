package main

import (
	"bytes"
	"net/http"
	"testing"
)

// func TestPing(t *testing.T) {

// 	// Create a new instance of our application struct. For now, this just
// 	// contains a couple of mock loggers (which discard anything writtern to them)
// 	app := &application{
// 		errorLog: log.New(ioutil.Discard, "", 0),
// 		infoLog:  log.New(ioutil.Discard, "", 0),
// 	}

// 	// This starts up a HTTPS server which listens on a randomly-chosen port.
// 	// Notice that we defer a call to ts.Close() to shut down the server when the test finishes.
// 	ts := httptest.NewTLSServer(app.routes())
// 	defer ts.Close()

// 	// The newwork address is contained in the ts.URL field.
// 	// This returns a http.Response struct containing the response.
// 	rs, err := ts.Client().Get(ts.URL + "/ping")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if rs.StatusCode != http.StatusOK {
// 		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
// 	}

// 	// And we can check that the response body writtern by the ping handler equals "ok".
// 	defer rs.Body.Close()
// 	body, err := ioutil.ReadAll(rs.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if string(body) != "OK" {
// 		t.Errorf("want body to equal %q", "OK")
// 	}
// }

// V2

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")
	if code != http.StatusOK {
		t.Errorf("want %q; got %q", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}

}

func TestShowSnippet(t *testing.T) {
	// Create a new instance of our application struct which users the mocked dependencies.
	app := newTestApplication(t)

	// Establish a new test server for running end-to-end tests.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Set up some table-driven tests to check the responses.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/snippet/1", http.StatusOK, []byte("An old silent pond...")},
		{"Non-existent ID", "/snippet/99", http.StatusNotFound, nil},
		{"Negatve ID", "/snippet/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/snippet/1.23", http.StatusNotFound, nil},
		{"String ID", "/snippet/foo", http.StatusNotFound, nil},
		{"Enable ID", "/snippet/", http.StatusNotFound, nil},
		{"Trailing slash", "/snippet/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}
		})
	}

}

func TestSignupUser(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Make a GET /user/signup request and then extract the CSRF token from the response body.
	_, _, body := ts.get(t, "/user/signup")
	csrfToken := extractCSRFToken(t, body)

	// Log the CSRF token value in our test output.
	// To see the output from the t.Log() command you need to run
	// `go test ` with the -v(verbose) flag enabled.
	t.Log(csrfToken)
}
