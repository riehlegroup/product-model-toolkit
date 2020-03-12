package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

const basePath = "/api/v1"

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

	if rec.Body.String() != version.Name() {
		t.Errorf("Expected body to be '%v', but got '%v'", version.Name(), rec.Body.String())
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
