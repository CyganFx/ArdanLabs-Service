// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"github.com/CyganFx/ArdanLabs-Service/business/auth"
	"github.com/CyganFx/ArdanLabs-Service/business/mid"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"

	"github.com/CyganFx/ArdanLabs-Service/foundation/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth, db *sqlx.DB) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	cg := checkGroup{
		build: build,
		db:    db,
	}

	app.Handle(http.MethodGet, "/readiness", cg.readiness, mid.Authenticate(a), mid.Authorize(auth.RoleAdmin))

	return app
}
