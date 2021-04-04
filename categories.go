// Copyright 2020 Eurac Research. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package snipeit

import (
	"fmt"
	"net/http"
)

// CategoryOptions specifies a subset of optional query parameters for listing
// categories.
type CategoryOptions struct {
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
	Search string `url:"search,omitempty"`
	Sort   string `url:"sort,omitempty"`
	Order  string `url:"order,omitempty"`
}

// Category represents a Snipe-IT category.
type Category struct {
	ID                int64     `json:"id,omitempty"`
	Name              string    `json:"name,omitempty"`
	Image             string    `json:"image,omitempty"`
	CategoryType      string    `json:"category_type,omitempty"`
	Eula              bool      `json:"eula,omitempty"`
	CheckinEmail      bool      `json:"checkin_email,omitempty"`
	RequireAcceptance bool      `json:"require_acceptance,omitempty"`
	AssetsCount       int64     `json:"assets_count,omitempty"`
	AccessoriesCount  int64     `json:"accessories_count,omitempty"`
	ConsumablesCount  int64     `json:"consumables_count,omitempty"`
	ComponentsCount   int64     `json:"components_count,omitempty"`
	LicensesCount     int64     `json:"licenses_count,omitempty"`
	CreatedAt         Timestamp `json:"created_at,omitempty"`
	UpdatedAt         Timestamp `json:"updated_at,omitempty"`
	Actions           struct {
		Update bool `json:"update,omitempty"`
		Delete bool `json:"delete,omitempty"`
	} `json:"available_actions,omitempty"`
}

// Categories lists all categories.
//
// Snipe-IT API doc: https://snipe-it.readme.io/reference#categories-1
func (c *Client) Categories(opt *CategoryOptions) ([]*Category, *http.Response, error) {
	u, err := c.AddOptions("categories", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Total int64
		Rows  []*Category
	}
	resp, err := c.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Rows, resp, nil
}

// Category by ID.
//
// Snipe-IT API doc: https://snipe-it.readme.io/reference#category
func (c *Client) Category(id int64) (*Category, *http.Response, error) {
	u := fmt.Sprintf("categories/%d", id)

	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	l := new(Category)
	resp, err := c.Do(req, l)
	if err != nil {
		return nil, resp, err
	}

	return l, resp, nil
}
