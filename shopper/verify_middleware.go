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