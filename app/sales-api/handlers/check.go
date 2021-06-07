package handlers

import (
	"context"
	"github.com/CyganFx/ArdanLabs-Service/foundation/database"
	"github.com/CyganFx/ArdanLabs-Service/foundation/web"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type checkGroup struct {
	build string
	db    *sqlx.DB
}

func (cg checkGroup) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := "ok"
	statusCode := http.StatusOK
	if err := database.StatusCheck(ctx, cg.db); err != nil {
		status = "db not ready"
		statusCode = http.StatusInternalServerError
	}

	health := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	return web.Respond(ctx, w, health, statusCode)
}
