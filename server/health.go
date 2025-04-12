package main

import (
	"net/http"

	"github.com/table-tap/api/internal/types"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, types.SuccessResponse)
}