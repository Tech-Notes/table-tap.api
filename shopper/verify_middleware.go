package shopper

import (
	"context"
	"net/http"

	internal "github.com/table-tap/api/internal/types"
)

func verify(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tableToken := r.Header.Get("TABLE-TOKEN")
		if tableToken == "" {
			writeError(w, http.StatusBadRequest, ErrInvalidTableToken)
			return
		}

		table, err := DBConn.GetTableByToken(ctx, tableToken)
		if err != nil {
			writeError(w, http.StatusBadRequest, ErrInvalidTableToken)
			return
		}

		if table != nil && table.ID == 0 {
			writeError(w, http.StatusBadRequest, ErrInvalidTableToken)
			return
		}

		ctx = context.WithValue(ctx, internal.ContextKeyBusinessID, table.BusinessID)
		ctx = context.WithValue(ctx, internal.ContextKeyTableID, table.ID)
		
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func businessIDFromContext(ctx context.Context) int64 {
	businessID := ctx.Value(internal.ContextKeyBusinessID)
	if businessID, ok := businessID.(int64); ok {
		return businessID
	}
	return 0
}

func tableIDFromContext(ctx context.Context) int64 {
	tableID := ctx.Value(internal.ContextKeyTableID)
	if tableID, ok := tableID.(int64); ok {
		return tableID
	}
	return 0
}