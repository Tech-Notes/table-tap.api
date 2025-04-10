package main

import types "github.com/table-tap/api/internal/types"

var (
	ErrInvalidOrderStatus *types.Error = &types.Error{Code: "invalid-order-status", Message: "Order status is not correct."}
	ErrRequiredOrderID    *types.Error = &types.Error{Code: "required-order-id", Message: "Order id is required."}
	ErrFailedRequestBody  *types.Error = &types.Error{Code: "failed-request-body", Message: "Faild to read request body."}
)
