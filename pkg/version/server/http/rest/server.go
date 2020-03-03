package rest

import (
	"context"
	"os"
	"os/signal"
	"time"

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
	v1 := e.Group("/api/v1")
	v1.GET("/", handleEntryPoint)
	v1.GET("/version", handleVersion)
	v1.GET("/health", handleHealth)

	return &Instance{
		httpSrv: e,
		addr:    address,
	}
}

// Start starts the REST server.
func (srv *Instance) Start() {
	srv.httpSrv.Start(srv.Addr())
}

// Shutdown perform a gracefully shutdown of the HTTP sever with a 10 seconds timeout
func (srv *Instance) Shutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.httpSrv.Shutdown(ctx); err != nil {
		srv.httpSrv.Logger.Fatal(err)
	}
}

// Addr return the address with port of the REST server.
func (srv *Instance) Addr() string {
	return srv.addr
}
