package utils

import (
	"context"

	"github.com/lib/pq"
	types "github.com/table-tap/api/internal/types"
)

func BusinessUserPermissionsFromContext(ctx context.Context) []string {
	permissions := ctx.Value(types.ContextKeyPermissions)
	if permissions, ok := permissions.(pq.StringArray); ok {
		return permissions
	}
	return nil
}

func BusinessIDFromContext(ctx context.Context) int64 {
	businessID := ctx.Value(types.ContextKeyBusinessID)
	if businessID, ok := businessID.(int64); ok {
		return businessID
	}
	return 0
}

func TableIDFromContext(ctx context.Context) int64 {
	tableID := ctx.Value(types.ContextKeyTableID)
	if tableID, ok := tableID.(int64); ok {
		return tableID
	}
	return 0
}

func TableNoFromContext(ctx context.Context) int64 {
	tableNo := ctx.Value(types.ContextKeyTableNo)
	if tableNo, ok := tableNo.(int64); ok {
		return tableNo
	}
	return 0
}
