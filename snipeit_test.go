// Copyright 2020 Eurac Research. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package snipeit

import (
	"encoding/json"
	"fmt"
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

func TestLocations(t *testing.T) {
	mux.HandleFunc("/locations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeaders(t, r)
		testFormValues(t, r, map[string]string{
			"search": "Test",
		})

		fmt.Fprint(w, `{"total":1,"rows":[{"id": 1, "name": "Test"}]}`)
	})

	opt := &LocationOptions{
		Search: "Test",
	}
	locations, _, err := testClient.Locations(opt)
	if err != nil {
		t.Errorf("Locations returend error: %v", err)
	}

	var want = []*Location{{ID: 1, Name: "Test"}}
	if !reflect.DeepEqual(locations, want) {
		t.Errorf("Locations returned %v, want %+v", locations, want)
	}
}

func TestLocation(t *testing.T) {
	mux.HandleFunc("/locations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeaders(t, r)
		fmt.Fprint(w, `{"id": 1, "name": "Test"}`)
	})

	location, _, err := testClient.Location(1)
	if err != nil {
		t.Errorf("Location returned error: %v", err)
	}

	var want = &Location{ID: 1, Name: "Test"}
	if !reflect.DeepEqual(location, want) {
		t.Errorf("Location returned %v, want %+v", location, want)
	}
}

func TestHardware(t *testing.T) {
	mux.HandleFunc("/hardware", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeaders(t, r)
		testFormValues(t, r, map[string]string{
			"location_id": "1",
		})

		fmt.Fprint(w, `{"total":1, "rows": [{"id": 10, "name": "hardware", "location": 1}]}`)
	})

	opt := &HardwareOptions{
		LocationID: 1,
	}
	hardware, _, err := testClient.Hardware(opt)
	if err != nil {
		t.Errorf("Hardware returend error: %v", err)
	}

	var want = []*Hardware{{ID: 10, Name: "hardware", Location: 1}}
	if !reflect.DeepEqual(hardware, want) {
		t.Errorf("Hardware returend %v, want %+v", hardware, want)
	}
}

func TestCategories(t *testing.T) {
	mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeaders(t, r)
		testFormValues(t, r, map[string]string{
			"search": "Test",
		})

		fmt.Fprint(w, `{"total":1,"rows":[{"id": 1, "name": "Test"}]}`)
	})

	opt := &CategoryOptions{
		Search: "Test",
	}
	categories, _, err := testClient.Categories(opt)
	if err != nil {
		t.Errorf("Categories returend error: %v", err)
	}

	var want = []*Category{{ID: 1, Name: "Test"}}
	if !reflect.DeepEqual(categories, want) {
		t.Errorf("Categories returned %v, want %+v", categories, want)
	}
}

func TestCategory(t *testing.T) {
	mux.HandleFunc("/categories/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeaders(t, r)
		fmt.Fprint(w, `{"id": 1, "name": "Test"}`)
	})

	category, _, err := testClient.Category(1)
	if err != nil {
		t.Errorf("Categories returned error: %v", err)
	}

	var want = &Category{ID: 1, Name: "Test"}
	if !reflect.DeepEqual(category, want) {
		t.Errorf("Categories returned %v, want %+v", category, want)
	}
}

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
			t.Errorf("Header.Get(%q) returend %q, want %q", tc.key, got, tc.want)
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
