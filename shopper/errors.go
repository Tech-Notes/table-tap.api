package shopper

import (
	types "github.com/table-tap/api/internal/types"
)

var (
	ErrInvalidTableToken *types.Error = &types.Error{Code: "invalid-table-token", Message: "Table token is not correct, please scan again."}
	ErrFailedRequestBody *types.Error = &types.Error{Code: "failed-request-body", Message: "Faild to read request body."}
	ErrRequiredShopID    *types.Error = &types.Error{Code: "required-shop-id", Message: "Shop id is required."}
)
