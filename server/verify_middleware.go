package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/table-tap/api/internal/types"
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

func generateHMACSignature(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	return signature
}

func verifyWithHMACSignature(h http.Handler) http.Handler {
	secretApiKey := os.Getenv("SECRET_API_KEY")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := "URL:" + r.URL.String()
		if strings.Contains(r.URL.Path, "notifications") {
			data = "URL:" + r.URL.Path
		}
		hmacSignature := generateHMACSignature(data, secretApiKey)

		if !hmac.Equal([]byte(hmacSignature), []byte(r.Header.Get("X-HMAC-SIGNATURE"))) {
			writeError(w, http.StatusUnauthorized, errors.New("invalid HMAC signature"))
			return
		}

		userEmail := r.Header.Get("X-USER-EMAIL")

		ctx, err := getContext(r.Context(), userEmail)
		if err != nil {
			writeError(w, http.StatusUnauthorized, err)
			return
		}

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func verifyWithAdminSecret(h http.Handler) http.Handler {
	adminSecret := os.Getenv("X_ADMIN_SECRET")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-ADMIN-SECRET") != adminSecret {
			writeError(w, http.StatusUnauthorized, errors.New("invalid admin secret"))
			return
		}

		userEmail := r.Header.Get("X-USER-EMAIL")

		ctx, err := getContext(r.Context(), userEmail)
		if err != nil {
			writeError(w, http.StatusUnauthorized, err)
			return
		}

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getContext(ctx context.Context, userEmail string) (context.Context, error) {
	if userEmail == "" {
		return nil, errors.New("missing user email")
	}

	businessUser, err := DBConn.GetLastActiveBusinessUserByEmail(ctx, userEmail)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	ctx = context.WithValue(ctx, types.ContextkeyUserEmail, userEmail)
	ctx = context.WithValue(ctx, types.ContextKeyUserID, businessUser.ID)
	ctx = context.WithValue(ctx, types.ContextKeyBusinessID, businessUser.BusinessID)
	ctx = context.WithValue(ctx, types.ContextKeyPermissions, businessUser.Permissions)

	return ctx, nil
}
