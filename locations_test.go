// Copyright 2020 Eurac Research. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package snipeit

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
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
		t.Errorf("Locations returned error: %v", err)
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
