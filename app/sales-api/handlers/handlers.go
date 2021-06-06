package handlers

import (
	"github.com/CyganFx/ArdanLabs-Service/business/mid"
	"log"
	"net/http"
	"os"

	"github.com/CyganFx/ArdanLabs-Service/foundation/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	check := check{logger: log}

	app.Handle(http.MethodGet, "/readiness", check.readiness)

	return app
}
