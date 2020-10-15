// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client to communicate with Product-Model-Toolkit server.
type Client struct {
	baseURL    string
	HTTPClient HTTPClient
}

// HTTPClient interface to simplify testing
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewClient returns a HTTP client.
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 3 * time.Minute,
		},
	}
}

// GetServerVersion returns the semantic version of the PMT server.
func (c *Client) GetServerVersion() (string, error) {
	url := fmt.Sprintf("%s/version", c.baseURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	return buf.String(), nil
}

const locHeader = "Location"

// PostResult send result JSON to PMT server and returns the path to the created resource.
func (c *Client) PostResult(url string, input io.Reader) (string, error) {
	completeURL := c.baseURL + url

	req, err := http.NewRequest(http.MethodPost, completeURL, input)
	if err != nil {
		return "", err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	return res.Header.Get(locHeader), nil
}
