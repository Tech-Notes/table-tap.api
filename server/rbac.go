package main

import (
	"context"
	"net/http"
	"slices"

	"github.com/table-tap/api/internal/utils"
)

type PermissionName string

const (
	DashboardView PermissionName = "dashboard_view"
	CreateTable   PermissionName = "create_table"
)

func checkPermissionAccess(ctx context.Context, permission PermissionName) bool {
	businessUserPermissions := utils.BusinessUserPermissionsFromContext(ctx)
	return slices.Contains(businessUserPermissions, string(permission))
}

func authorize(permission PermissionName, h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if checkPermissionAccess(r.Context(), permission) {
			h.ServeHTTP(w, r)
			return
		}
		writeError(w, http.StatusForbidden, ErrUnauthorizedAccess)
	})
}

func authorizeHandler(permission PermissionName, h http.HandlerFunc) http.HandlerFunc {
	handler := authorize(permission, h)
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}
