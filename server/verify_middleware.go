package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"os"
)

func verify(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		adminSecret := r.Header.Get("X-ADMIN-SECRET")
		if adminSecret != "" {
			verifyWithAdminSecret(h).ServeHTTP(w, r)
			return
		}
		
		hmacSignature := r.Header.Get("X-HMAC-SIGNATURE")
		if hmacSignature == "" {
			writeError(w, http.StatusUnauthorized, errors.New("missing HMAC signature"))
			return
		}

		verifyWithHMACSignature(h).ServeHTTP(w, r)
		return
	})
}

func verifyWithHMACSignature(h http.Handler) http.Handler {
	secretApiKey := os.Getenv("SECRET_API_KEY")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hmacSignature := generateHMACSignature(r.URL.Path, secretApiKey)

		if hmacSignature != r.Header.Get("X-HMAC-SIGNATURE") {
			writeError(w, http.StatusUnauthorized, errors.New("invalid HMAC signature"))
			return
		}

		h.ServeHTTP(w, r)
	})
}

func generateHMACSignature(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	return signature
}

func verifyWithAdminSecret(h http.Handler) http.Handler {
	adminSecret := os.Getenv("X_ADMIN_SECRET")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-ADMIN-SECRET") != adminSecret {
			writeError(w, http.StatusUnauthorized, errors.New("invalid admin secret"))
			return
		}

		h.ServeHTTP(w, r)
	})
}