package main

import types "github.com/table-tap/api/internal/types"

var (
	ErrInvalidOrderStatus *types.Error = &types.Error{Code: "invalid-order-status", Message: "Order status is not correct."}
	ErrRequiredOrderID    *types.Error = &types.Error{Code: "required-order-id", Message: "Order id is required."}
	ErrFailedRequestBody  *types.Error = &types.Error{Code: "failed-request-body", Message: "Faild to read request body."}
	ErrUnauthorizedAccess *types.Error = &types.Error{Code: "unauthorize-access", Message: "Permission denied."}
	ErrUserNotFound       *types.Error = &types.Error{Code: "user-not-found", Message: "User not found."}
	ErrInvalidCredentials *types.Error = &types.Error{Code: "invalid-credentials", Message: "Invalid credentials."}
	ErrRequiredTableID    *types.Error = &types.Error{Code: "required-table-id", Message: "Table is required."}
)
