// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"github.com/CyganFx/ArdanLabs-Service/business/auth"
	"github.com/CyganFx/ArdanLabs-Service/business/data/user"
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

	app.Handle(http.MethodGet, "/readiness", cg.readiness)
	//app.Handle(http.MethodGet, "/liveness", cg.liveness)

	// Register user management and authentication endpoints.
	ug := userGroup{
		user: user.New(log, db),
		auth: a,
	}
	app.Handle(http.MethodGet, "/v1/users/:page/:rows", ug.query, mid.Authenticate(a), mid.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodGet, "/v1/users/:id", ug.queryByID, mid.Authenticate(a))
	app.Handle(http.MethodGet, "/v1/users/token/:kid", ug.token)
	app.Handle(http.MethodPost, "/v1/users", ug.create, mid.Authenticate(a), mid.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodPut, "/v1/users/:id", ug.update, mid.Authenticate(a), mid.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodDelete, "/v1/users/:id", ug.delete, mid.Authenticate(a), mid.Authorize(auth.RoleAdmin))

	return app
}
