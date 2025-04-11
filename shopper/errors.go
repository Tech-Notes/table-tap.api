package shopper

import (
	types "github.com/table-tap/api/internal/types"
)

var (
	ErrInvalidTableToken *types.Error = &types.Error{Code: "invalid-table-token", Message: "Table token is not correct, please scan again."}
)