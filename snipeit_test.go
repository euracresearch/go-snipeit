// Copyright 2020 Eurac Research. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package snipeit

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"
)

const testToken = "premature optimization is the root of all evil (or at least most of it) in programming"

var (
	mux        *http.ServeMux // mux is the HTTP request multiplexer used with the test server.
	testClient *Client
)

func TestTimestampJSONUnmarshal(t *testing.T) {
	in := `{
		"datetime":"2019-05-21 21:37:40",
	     	"formatted":"2019-05-21 21:37"
	}`

	var got Timestamp
	if err := json.Unmarshal([]byte(in), &got); err != nil {
		t.Errorf("json.Unmarshal of Timestamp type failed: %v", err)
	}

	want := time.Date(2019, time.May, 21, 21, 37, 40, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("Timestamp is %v, want %v", got, want)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeaders(t *testing.T, r *http.Request) {
	h := []struct {
		key  string
		want string
	}{
		{"Content-Type", "application/json"},
		{"Accept", "application/json"},
		{"Authorization", "Bearer " + testToken},
	}

	for _, tc := range h {
		if got := r.Header.Get(tc.key); got != tc.want {
			t.Errorf("Header.Get(%q) returned %q, want %q", tc.key, got, tc.want)
		}
	}
}

func testFormValues(t *testing.T, r *http.Request, values map[string]string) {
	want := url.Values{}
	for k, v := range values {
		want.Set(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}

func TestMain(m *testing.M) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)

	var err error
	testClient, err = NewClient(server.URL, testToken)
	if err != nil {
		log.Fatal(err)
	}

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
