// Copyright 2020 Eurac Research. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package snipeit provides a client for communicating with the Snipe-IT API and
// defines Snipe-IT specific data types.
package snipeit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

// A Client manages communication with the Snipe-IT API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.
	token  string       // Snipe-IT API personal API token.

	BaseURL *url.URL
}

// NewClient returns a new Snipe-IT API client with provided base URL. If base
// URL does not have a trailing slash, one is added automatically.
func NewClient(baseURL, token string) (*Client, error) {
	if baseURL == "" {
		return nil, errors.New("a baseURL must be provided")
	}
	if token == "" {
		return nil, errors.New("a token must be provided")
	}

	baseEndpoint, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}

	c := new(Client)
	c.client = http.DefaultClient
	c.token = "Bearer " + token
	c.BaseURL = baseEndpoint
	return c, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(strings.TrimPrefix(urlStr, "/"))
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.token)

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// If StatusCode is not in the 200 range something went wrong, return the
	// response but do not process it's body.
	if c := resp.StatusCode; 200 > c || c > 299 {
		return resp, nil
	}

	defer resp.Body.Close()
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

// AddOptions adds the parameters in opt as URL query parameters to s. opt must
// be a struct whose fields may contain "url" tags.
func (c *Client) AddOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// Timestamp is a custom time type for parsing Snipe-ITs API updated_at and
// created_at JSON values.
type Timestamp struct {
	time.Time
}

type apiTimestamp struct {
	Datetime string `json:"datetime"`
	Format   string `json:"formatted"`
}

func (ts *Timestamp) UnmarshalJSON(b []byte) error {
	d := &apiTimestamp{}
	if err := json.Unmarshal(b, d); err != nil {
		return err
	}
	if d.Datetime == "" {
		return nil
	}
	const format = "2006-01-02 15:04:05"
	t, err := time.Parse(format, d.Datetime)
	if err != nil {
		return fmt.Errorf("go-snipeit: can not parse api timestamp: %v", err)
	}
	ts.Time = t
	return nil
}
