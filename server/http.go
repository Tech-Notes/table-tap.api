package main

import (
	"encoding/json"
	"net/http"

	internal "github.com/table-tap/api/internal/types"
)

func readJSON(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(dst)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Write(b)
}

func genericError(err error) *internal.Error {
	return &internal.Error{Code: "generic", Message: err.Error()}
}

func writeError(w http.ResponseWriter, status int, err error) {

	writeJSON(w, status, &internal.ResponseBase{
		Status: internal.ResponseStatusError,
		Error: genericError(err),
	})
}