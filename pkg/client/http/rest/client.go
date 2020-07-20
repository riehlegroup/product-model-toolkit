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
	HTTPClient *http.Client
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
	res, err := c.HTTPClient.Get(url)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	return buf.String(), nil
}

// PostComposerResult send result JSON to PMT server.
func (c *Client) PostComposerResult(input io.Reader) error {
	url := fmt.Sprintf("%s/products/composer", c.baseURL)
	_, err := c.HTTPClient.Post(url, "application/json", input)
	if err != nil {
		return err
	}

	return nil
}
