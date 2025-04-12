package internal

import (
	"context"

	"github.com/lib/pq"
	internal "github.com/table-tap/api/internal/types"
)

func BusinessUserPermissionsFromContext(ctx context.Context) []string {
	permissions := ctx.Value(internal.ContextKeyPermissions)
	if permissions, ok := permissions.(pq.StringArray); ok {
		return permissions
	}
	return nil
}

func BusinessIDFromContext(ctx context.Context) int64 {
	businessID := ctx.Value(internal.ContextKeyBusinessID)
	if businessID, ok := businessID.(int64); ok {
		return businessID
	}
	return 0
}

func TableIDFromContext(ctx context.Context) int64 {
	tableID := ctx.Value(internal.ContextKeyTableID)
	if tableID, ok := tableID.(int64); ok {
		return tableID
	}
	return 0
}