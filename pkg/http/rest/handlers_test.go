// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"encoding/json"
	"errors"
	importing2 "github.com/osrgroup/product-model-toolkit/pkg/services/importing"
	querying2 "github.com/osrgroup/product-model-toolkit/pkg/services/querying"
	version2 "github.com/osrgroup/product-model-toolkit/pkg/services/version"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/osrgroup/product-model-toolkit/pkg/db/memory"
)

const basePath = "/api/v1"

func TestHandler(t *testing.T) {
	e := echo.New()
	v1 := e.Group("/api/v1")

	repo := new(memory.DB)
	Handler(v1, querying2.NewService(&mockDB{}), importing2.NewService(repo))
}

func TestHandleEntryPoint(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, basePath+"/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handleEntryPoint(c)

	if err != nil {
		t.Errorf(err.Error())
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code to be %v, but got %v", http.StatusOK, rec.Code)
	}
}

func TestHandleVersion(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, basePath+"/version", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handleVersion(c)

	if err != nil {
		t.Errorf(err.Error())
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code to be %v, but got %v", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != version2.Name() {
		t.Errorf("Expected body to be '%v', but got '%v'", version2.Name(), rec.Body.String())
	}
}

func TestHandleHealth(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, basePath+"/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handleHealth(c)

	if err != nil {
		t.Errorf(err.Error())
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code to be %v, but got %v", http.StatusOK, rec.Code)
	}

	expectedBody := "{\"status\":\"UP\"}\n"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body to be '%v', but got '%v'", expectedBody, rec.Body.String())
	}
}

func TestFindAllProducts(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, basePath+"/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	q := inMemQueryingService()

	handler := findAllProducts(q)
	err := handler(c)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code to be %v, but got %v", http.StatusOK, rec.Code)
	}

	var prods []model.Product
	err = json.Unmarshal(rec.Body.Bytes(), &prods)
	if err != nil {
		t.Errorf("Expected unmarshaling JSON without an error, but got %v", err.Error())
	}

	if len(prods) != 2 {
		t.Errorf("Expected %v products as result, but go %v", 2, len(prods))
	}
}

func TestFindAllProducts_Error(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, basePath+"/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	q := querying2.NewService(&mockDB{})
	handler := findAllProducts(q)
	handler(c)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code to be %v, but got %v", http.StatusInternalServerError, rec.Code)
	}
}

func TestFindProductByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, basePath+"/products/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	q := inMemQueryingService()

	handler := findProductByID(q)
	handler(c)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code to be %v, but got %v", http.StatusOK, rec.Code)
	}

	var prod model.Product
	err := json.Unmarshal(rec.Body.Bytes(), &prod)
	if err != nil {
		t.Errorf("Expected unmarshaling JSON without an error, but got %v", err.Error())
	}

	expectedID := 2
	if prod.ID != expectedID {
		t.Errorf("Expected product ID to be %v, but got %v", expectedID, prod.ID)
	}
}

func TestFindProductByID_NotExisting(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, basePath+"/products/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("473638")

	q := inMemQueryingService()

	handler := findProductByID(q)
	err := handler(c)
	if err != nil {
		t.Errorf(err.Error())
	}

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status code to be %v, but got %v", http.StatusNotFound, rec.Code)
	}
}

func TestFindProductByID_Error(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, basePath+"/products/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("$!*")

	q := inMemQueryingService()

	handler := findProductByID(q)
	err := handler(c)
	if err != nil {
		t.Errorf(err.Error())
	}

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code to be %v, but got %v", http.StatusInternalServerError, rec.Code)
	}
}

func inMemQueryingService() querying2.Service {
	repo := new(memory.DB)
	repo.AddSampleData()

	return querying2.NewService(repo)
}

type mockDB struct{}

func (db *mockDB) FindAllProducts() ([]model.Product, error) {
	return nil, errors.New("some error")
}

func (db *mockDB) FindProductByID(id int) (model.Product, error) {
	return model.Product{}, errors.New("some error")
}
