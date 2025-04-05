package main

import (
	"errors"
	"net/http"
)

func verify(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		hmacSignature := r.Header.Get("X-HMAC-SIGNATURE")
		if hmacSignature == "" {
			writeError(w, http.StatusUnauthorized, errors.New("missing HMAC signature"))
			return
		}

		h.ServeHTTP(w, r)
	})
}