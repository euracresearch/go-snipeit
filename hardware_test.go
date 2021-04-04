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

func TestHardware(t *testing.T) {
	mux.HandleFunc("/hardware", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeaders(t, r)
		testFormValues(t, r, map[string]string{
			"location_id": "1",
		})

		fmt.Fprint(w, `{"total":1, "rows": [{"id": 10, "name": "hardware", "location": {"id": 1}}]}`)
	})

	opt := &HardwareOptions{
		LocationID: 1,
	}
	hardware, _, err := testClient.Hardware(opt)
	if err != nil {
		t.Errorf("Hardware returned error: %v", err)
	}

	var want = []*Hardware{{ID: 10, Name: "hardware", Location: &Location{ID: 1}}}
	if !reflect.DeepEqual(hardware, want) {
		t.Errorf("Hardware returned %v, want %+v", hardware, want)
	}
}
