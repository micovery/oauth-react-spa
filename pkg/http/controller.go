// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/micovery/oauth-react-spa/webapp"
	"github.com/ttys3/echo-pprof/v4"
	"gopkg.in/tomb.v2"
	"net/http"
	"os"
	"os/signal"
)

type Controller struct {
	port     string
	httpTomb *tomb.Tomb
	Echo     *echo.Echo
}

func NewCtl(port string) *Controller {
	httpCtl := &Controller{
		port: port,
	}

	if err := httpCtl.Init(); err != nil {
		panic(err)
	}

	return httpCtl
}

func (c *Controller) Init() (err error) {
	c.Echo, err = c.NewEcho(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) NewEcho(err error) (*echo.Echo, error) {

	echoServer := echo.New()
	echoHandler := echoServer
	echopprof.Wrap(echoHandler)

	//override server handler to intercept grpc-web requests (content-type: application/grpc-web)
	echoServer.Server = &http.Server{
		Addr: fmt.Sprintf(":%s", c.port),
		Handler: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			resp.Header().Set("Access-Control-Allow-Headers", "*")
			resp.Header().Set("Access-Control-Allow-Origin", "*")
			echoHandler.ServeHTTP(resp, req)
		}),
	}

	var httpFS http.FileSystem
	if httpFS, err = webapp.HTTPFileSystem(); err != nil {
		return nil, err
	}

	echoServer.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: httpFS,
		HTML5:      true,
	}))

	echoServer.HideBanner = true
	return echoServer, nil
}

func (c *Controller) Start() (err error) {
	if c.httpTomb != nil && c.httpTomb.Alive() {
		return errors.New("http already started")
	}

	c.httpTomb = &tomb.Tomb{}
	c.httpTomb.Go(func() error {
		fmt.Printf("⇨ http server started on [::]%s\n", c.Echo.Server.Addr)
		if err := c.Echo.Server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println("(http): server halted forcefully")
			return err
		}

		fmt.Println("(http): server halted gracefully")
		return nil
	})

	c.httpTomb.Go(func() error {
		<-c.httpTomb.Dying()
		if err := c.Echo.Shutdown(nil); err != nil {
			return err
		}
		return nil
	})

	go func() {
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, os.Interrupt)
		<-sigChan
		fmt.Println("Ctrl-C detected, exiting")
		if err := c.Stop(); err != nil {
			fmt.Printf("could not stop HTTP controller. %s\n", err.Error())
		}
	}()

	return nil
}

func (c *Controller) Stop() (err error) {
	if c.httpTomb == nil || !c.httpTomb.Alive() {
		return http.ErrServerClosed
	}

	c.httpTomb.Kill(nil)
	if err := c.httpTomb.Wait(); err != nil {
		return err
	}
	return nil
}

func (c *Controller) Quit() {
	if err := c.Stop(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("error while exiting http controller. %s\n", err.Error())
	}
}

func (c *Controller) Wait() {
	err := c.httpTomb.Wait()
	if err != nil {
		panic(err)
	}
}
