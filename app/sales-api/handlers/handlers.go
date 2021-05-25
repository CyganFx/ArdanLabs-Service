package handlers

import (
	"log"
	"net/http"
	"os"

	
	"gitlab.com/tleuzhan13/se1903Service/foundation/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {
	app := web.NewApp(shutdown)

	check := check{logger: log}

	app.Handle(http.MethodGet,"/readiness", check.readiness)

	return app
}