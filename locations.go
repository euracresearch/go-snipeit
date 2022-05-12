// Copyright 2020 Eurac Research. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package snipeit

import (
	"fmt"
	"net/http"
)

// LocationService handles communication with the location related methods of
// the SnipeIT-API.
//
// https://snipe-it.readme.io/reference/locations
type LocationService service

// Location represents a Snipe-IT location.
type Location struct {
	ID             int64     `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Image          string    `json:"image,omitempty"`
	Address        string    `json:"address,omitempty"`
	Address2       string    `json:"address2,omitempty"`
	City           string    `json:"city,omitempty"`
	State          string    `json:"state,omitempty"`
	Country        string    `json:"country,omitempty"`
	Zip            string    `json:"zip,omitempty"`
	AssetsAssigned int64     `json:"assigned_assets_count,omitempty"`
	Assets         int64     `json:"assets_count,omitempty"`
	Users          int64     `json:"users_count,omitempty"`
	Currency       string    `json:"currency,omitempty"`
	CreatedAt      Timestamp `json:"created_at,omitempty"`
	UpdatedAt      Timestamp `json:"updated_at,omitempty"`
	Parent         struct {
		ID   int64  `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"parent,omitempty"`
	Manager  string     `json:"manager,omitempty"`
	Children []Location `json:"children,omitempty"`
	Actions  struct {
		Update bool
		Delete bool
	} `json:"available_actions,omitempty"`
}

// LocationListOptions specifies a subset of optional query parameters for
// listing locations.
type LocationListOptions struct {
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
	Search string `url:"search,omitempty"`
	Sort   string `url:"sort,omitempty"`
	Order  string `url:"order,omitempty"`
}

// Locations lists all locations.
//
// Snipe-IT API doc: https://snipe-it.readme.io/reference#locations
func (s *LocationService) List(opt *LocationListOptions) ([]*Location, *http.Response, error) {
	u, err := s.client.AddOptions("locations", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Total int64
		Rows  []*Location
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Rows, resp, nil
}

// Location by ID.
//
// Snipe-IT API doc: https://snipe-it.readme.io/reference#locations-1
func (s *LocationService) Get(id int64) (*Location, *http.Response, error) {
	u := fmt.Sprintf("locations/%d", id)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	l := new(Location)
	resp, err := s.client.Do(req, l)
	if err != nil {
		return nil, resp, err
	}

	return l, resp, nil
}
