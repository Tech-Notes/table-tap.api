package internal

import (
	"net/http"

	types "github.com/table-tap/api/internal/types"
)

func genericError(err error) *types.Error {
	return &types.Error{Code: "generic", Message: err.Error()}
}

func WriteError(w http.ResponseWriter, status int, err error) {
	err2, ok := err.(*types.Error)
	if !ok {
		err2 = genericError(err)
	}
	WriteJSON(w, status, &types.ResponseBase{
		Status: types.ResponseStatusError,
		Error: err2,
	})
}