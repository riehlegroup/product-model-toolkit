// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
	"io/ioutil"
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

// newClient returns a HTTP client.
func newClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 3 * time.Minute,
		},
	}
}

// GetServerVersion returns the semantic version of the PMT server.
func (c *Client) getServerVersion() (string, error) {
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
func (c *Client) postResult(url string, input io.Reader) (string, error) {
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

func (c *Client) postData(url string, data []byte) (string, error) {
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
    req.Header.Set("Content-Type", "application/json")

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
		return "", nil
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
	return string(body), nil
}


// GetServerVersion returns the semantic version of the PMT server.
func (c *Client) GetProductId(id string) (*http.Response, error) {
	url := fmt.Sprintf("%s/products/%s", c.baseURL, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(res.Body)

	return res, nil
}