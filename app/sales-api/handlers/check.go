package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)


type check struct {
	logger *log.Logger
}

func(c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct{
		Status string
	}{
		Status: "Ok",
	}

	return json.NewEncoder(w).Encode(status)
}