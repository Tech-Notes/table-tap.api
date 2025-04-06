package main

import (
	"context"
	"errors"
	"net/http"
	"slices"
)

type PermissionName string

const (
	DashboardView PermissionName = "dashboard_view"
)

func checkPermissionAccess(ctx context.Context, permission PermissionName) bool {
	businessUserPermissions := BusinessUserPermissionsFromContext(ctx)
	return slices.Contains(businessUserPermissions, string(permission))
}

func authorize(permission PermissionName, h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if checkPermissionAccess(r.Context(), permission) {
			h.ServeHTTP(w, r)
			return
		}
		writeError(w, http.StatusForbidden, errors.New("permission denied"))
	})
}

func authorizeHandler(permission PermissionName, h http.HandlerFunc) http.HandlerFunc {
	handler := authorize(permission, h)
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}
