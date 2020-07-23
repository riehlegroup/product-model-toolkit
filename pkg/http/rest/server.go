// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/osrgroup/product-model-toolkit/pkg/importing"
	"github.com/osrgroup/product-model-toolkit/pkg/querying"
)

type Instance struct {
	httpSrv *echo.Echo
	addr    string
}

// NewSrv creates a new REST server.
func NewSrv(address string, qSrv querying.Service, iSrv importing.Service) *Instance {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	v1 := e.Group("/api/v1")
	Handler(v1, qSrv, iSrv)

	return &Instance{
		httpSrv: e,
		addr:    address,
	}
}

// Start starts the REST server.
func (srv *Instance) Start() {
	time.AfterFunc(300*time.Millisecond, srv.printRoutes)

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

// PrintRoutes prints all available routes to the standard logger.
func (srv *Instance) printRoutes() {
	var r string
	for _, v := range srv.httpSrv.Routes() {
		r += fmt.Sprintf("\t%s\t%s\n", v.Method, v.Path)
	}

	fmt.Printf("Available routes:\n%s\n", r)
}
