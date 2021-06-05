package handlers

import (
	"context"
	"github.com/CyganFx/ArdanLabs-Service/foundation/web"
	"log"
	"net/http"
)

type check struct {
	logger *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	//id := web.Param(r, "id")

	//var u user.User
	//if err := web.Decode(r, &u); err != nil {
	//	return err
	//}

	status := struct {
		Status string
	}{
		Status: "Ok",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
