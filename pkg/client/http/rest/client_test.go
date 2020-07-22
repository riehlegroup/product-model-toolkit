// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	urlShould := "api/v42/"
	c := NewClient(urlShould)

	if c.baseURL != urlShould {
		t.Errorf("Expected baseUrl to be %v, but got %v", urlShould, c.baseURL)
	}

	if c.HTTPClient == nil {
		t.Error("Expected HTTPClient to be not nil")
	}
}

func TestGetServerVersion(t *testing.T) {
	body := ioutil.NopCloser(bytes.NewReader([]byte("0.0.1")))
	c := &Client{
		baseURL:    "/someUrl/api/v1",
		HTTPClient: &mockHTTPClient{body: body},
	}

	v, _ := c.GetServerVersion()

	vShould := "0.0.1"
	if v != vShould {
		t.Errorf("Expected version to be %v, but got %v", vShould, v)
	}
}

func TestGetServerVersion_Error(t *testing.T) {
	c := &Client{
		baseURL:    "/someUrl/api/v1",
		HTTPClient: &mockHTTPClient{withErr: true},
	}

	_, err := c.GetServerVersion()
	if err == nil {
		t.Error("Expected to return error")
	}
}

func TestPostComposerResult(t *testing.T) {
	h := make(http.Header)
	locShould := "api/v1/products/42"
	h.Set("Location", locShould)

	c := &Client{
		baseURL:    "/someUrl/api/v1",
		HTTPClient: &mockHTTPClient{header: h},
	}

	loc, _ := c.PostComposerResult(nil)

	if loc != locShould {
		t.Errorf("Expected location header to be %v, but got %v", locShould, loc)
	}
}

func TestPostComposerResult_Error(t *testing.T) {
	c := &Client{
		baseURL:    "/someUrl/api/v1",
		HTTPClient: &mockHTTPClient{withErr: true},
	}

	_, err := c.PostComposerResult(nil)
	if err == nil {
		t.Error("Expected to return error")
	}
}

type mockHTTPClient struct {
	withErr bool
	body    io.ReadCloser
	header  http.Header
	doFunc  func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.withErr {
		return nil, errors.New("some error")
	}

	return &http.Response{
		Body:   m.body,
		Header: m.header,
	}, nil
}
