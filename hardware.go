// Copyright 2020 Eurac Research. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package snipeit

import "net/http"

// HardwareOptions specifies a subset of optional query parameters
// for listing assets.
type HardwareOptions struct {
	Limit          int    `url:"limit,omitempty"`
	Offset         int    `url:"offset,omitempty"`
	Search         string `url:"search,omitempty"`
	OrderNumber    string `url:"order_number,omitempty"`
	Sort           string `url:"sort,omitempty"`
	Order          string `url:"order,omitempty"`
	ModelID        int    `url:"model_id,omitempty"`
	CategoryID     int    `url:"category_id,omitempty"`
	ManufacturerID int    `url:"manufacturer_id,omitempty"`
	CompanyID      int    `url:"company_id,omitempty"`
	LocationID     int    `url:"location_id,omitempty"`
	Status         string `url:"status,omitempty"`
	StatusID       string `url:"status_id,omitempty"`
}

// Hardware represents a Snipe-IT hardware object.
type Hardware struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	AssetTag string `json:"asset_tag,omitempty"`
	Serial   string `json:"serial,omitempty"`
	Model    struct {
		ID   int64  `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"model,omitempty"`
	ModelNumber string `json:"model_number,omitempty"`
	StatusLabel struct {
		ID         int64  `json:"id,omitempty"`
		Name       string `json:"name,omitempty"`
		StatusMeta string `json:"status_meta,omitempty"`
	} `json:"status_label,omitempty"`
	Category     *Category `json:"category,omitempty"`
	Manufacturer struct {
		ID   int64  `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"manufacturer,omitempty"`
	Supplier struct {
		ID   int64  `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"supplier,omitempty"`
	Notes       string    `json:"notes,omitempty"`
	OrderNumber string    `json:"order_number,omitempty"`
	Company     string    `json:"company,omitempty"`
	Location    *Location `json:"location,omitempty"`
	RtdLocation *Location `json:"rtd_location,omitempty"`
	Image       string    `json:"image,omitempty"`
	AssignedTo  struct {
		ID        int64  `json:"id,omitempty"`
		Username  string `json:"username,omitempty"`
		Name      string `json:"name,omitempty"`
		Firstname string `json:"first_name,omitempty"`
		Lastname  string `json:"last_name,omitempty"`
		Emplyee   string `json:"employee_number,omitempty"`
		Type      string `json:"type,omitempty"`
	} `json:"assigned_to,omitempty"`
	WarrantyMonths   interface{}   `json:"warranty_months,omitempty"`
	WarrantyExpires  interface{}   `json:"warranty_expires,omitempty"`
	CreatedAt        Timestamp     `json:"created_at,omitempty"`
	UpdatedAt        Timestamp     `json:"updated_at,omitempty"`
	DeletedAt        Timestamp     `json:"deleted_at,omitempty"`
	PurchaseDate     Timestamp     `json:"purchase_date,omitempty"`
	LastCheckout     Timestamp     `json:"last_checkout,omitempty"`
	ExpectedCheckin  Timestamp     `json:"expected_checkin,omitempty"`
	PurchaseCost     string        `json:"purchase_cost,omitempty"`
	UserCanCheckout  bool          `json:"user_can_checkout,omitempty"`
	CustomFields     []interface{} `json:"custom_fields,omitempty"`
	AvailableActions struct {
		Checkout bool `json:"checkout,omitempty"`
		Checkin  bool `json:"checkin,omitempty"`
		Clone    bool `json:"clone,omitempty"`
		Restore  bool `json:"restore,omitempty"`
		Update   bool `json:"update,omitempty"`
		Delete   bool `json:"delete,omitempty"`
	} `json:"available_actions,omitempty"`
}

// Hardware lists all Hardware.
//
// https://snipe-it.readme.io/reference#hardware-list
func (c *Client) Hardware(opt *HardwareOptions) ([]*Hardware, *http.Response, error) {
	u, err := c.AddOptions("hardware", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Total int64
		Rows  []*Hardware
	}
	resp, err := c.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Rows, resp, nil
}
