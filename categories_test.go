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

func TestCategories(t *testing.T) {
	mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeaders(t, r)
		testFormValues(t, r, map[string]string{
			"search": "Test",
		})

		fmt.Fprint(w, `{"total":1,"rows":[{"id": 1, "name": "Test"}]}`)
	})

	opt := &CategoriesListOptions{
		Search: "Test",
	}
	categories, _, err := testClient.Categories.List(opt)
	if err != nil {
		t.Errorf("Categories returned error: %v", err)
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

	category, _, err := testClient.Categories.Get(1)
	if err != nil {
		t.Errorf("Categories returned error: %v", err)
	}

	var want = &Category{ID: 1, Name: "Test"}
	if !reflect.DeepEqual(category, want) {
		t.Errorf("Categories returned %v, want %+v", category, want)
	}
}
