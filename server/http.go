package main

import (
	"net/http"

	"github.com/table-tap/api/internal/httphelper"
)

func readJSON(r *http.Request, dst any) error {
	return httphelper.ReadJSON(r, dst)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	httphelper.WriteJSON(w, status, data)
}

func writeError(w http.ResponseWriter, status int, err error) {
	httphelper.WriteError(w, status, err)
}
