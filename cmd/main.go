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

package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/markbates/goth/providers/openidConnect"
	hc "github.com/micovery/oauth-react-spa/pkg/http"
	"net/http"
	_ "net/http/pprof"
	"os"
)

const AuthorityEnv = "AUTHORITY"
const ClientIdEnv = "CLIENT_ID"

func main() {

	port := os.Getenv("PORT")
	authority := os.Getenv(AuthorityEnv)
	clientId := os.Getenv(ClientIdEnv)

	if port == "" {
		port = "8080"
	}

	if authority == "" {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s environment variable is required", AuthorityEnv)
		os.Exit(1)
	}

	if clientId == "" {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s environment variable is required", ClientIdEnv)
		os.Exit(1)
	}

	httpCtl := hc.NewCtl(port)
	defer httpCtl.Quit()

	httpCtl.Echo.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().URL.Path == "/config.js" {
				c.Response().Header().Set("Content-Type", "text/javascript")
				return c.String(http.StatusOK, getConfigJs(authority, clientId))
			}
			return next(c)
		}
	})

	if err := httpCtl.Start(); err != nil {
		fmt.Printf("could not start HTTP server. %s\n", err.Error())
	}

	httpCtl.Wait()
}

func getConfigJs(authority string, clientId string) string {
	config := `
const config = {
  oidc: {
    authority: "` + authority + `",
    client_id: "` + clientId + `",
    redirect_uri: ` + "`${window.location.origin}/redirect`" + `,
  }
};`
	return config
}
