package main

import (
	"net/http"

	internal "github.com/table-tap/api/internal"
)

func readJSON(r *http.Request, dst any) error {
	return internal.ReadJSON(r, dst)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	internal.WriteJSON(w, status, data)
}

func writeError(w http.ResponseWriter, status int, err error) {
	internal.WriteError(w, status, err)
}