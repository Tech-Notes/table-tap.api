package shopper

import (
	"context"
	"errors"
	"net/http"

	"github.com/table-tap/api/internal/types"
)

func verify(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tableToken := r.Header.Get("TABLE-TOKEN")
		if tableToken == "" {
			writeError(w, http.StatusBadRequest, ErrInvalidTableToken)
			return
		}

		ctx, err := getContext(ctx, tableToken)
		if err != nil {
			if errors.Is(err, ErrInvalidTableToken) {
				writeError(w, http.StatusBadRequest, ErrInvalidTableToken)
			}
			writeError(w, http.StatusUnauthorized, err)
			return
		}

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getContext(ctx context.Context, tableToken string) (context.Context, error) {

	table, err := DBConn.GetTableByToken(ctx, tableToken)
	if err != nil {
		return nil, ErrInvalidTableToken
	}

	if table != nil && table.ID == 0 {
		return nil, ErrInvalidTableToken
	}
	ctx = context.WithValue(ctx, types.ContextKeyBusinessID, table.BusinessID)
	ctx = context.WithValue(ctx, types.ContextKeyTableID, table.ID)
	ctx = context.WithValue(ctx, types.ContextKeyTableNo, table.TableNo)

	return ctx, nil
}
