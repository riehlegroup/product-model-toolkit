package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Instance struct {
	httpSrv *echo.Echo
	addr    string
}

// NewSrv creates a new REST server.
func NewSrv(address string) *Instance {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handleVersion)
	e.GET("/version", handleVersion)
	e.GET("/health", handleHealth)

	return &Instance{
		httpSrv: e,
		addr:    address,
	}
}

// Start starts the REST server.
func (srv *Instance) Start() {
	srv.httpSrv.Start(srv.Addr())
}

// Addr return the address with port of the REST server.
func (srv *Instance) Addr() string {
	return srv.addr
}
