package main

import (
	"net/http"

	internal "github.com/table-tap/api/internal/types"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, internal.SuccessResponse)
}