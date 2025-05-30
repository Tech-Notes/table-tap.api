package types

type ContextKey string

const (
	ContextKeyUserID      ContextKey = "user_id"
	ContextkeyUserEmail   ContextKey = "user_email"
	ContextKeyBusinessID  ContextKey = "business_id"
	ContextKeyPermissions ContextKey = "permissions"
	ContextKeyTableID     ContextKey = "table_id"
	ContextKeyTableNo     ContextKey = "table_no"
)

type TokenClaim struct {
	UserID     int64  `json:"user_id"`
	UserEmail  string `json:"user_email"`
	BusinessID int64  `json:"business_id"`
	Role       string `json:"role"`
	RoleID     int64  `json:"role_id"`
}

type ActionSuccessResponseData struct {
	ID int64 `json:"id"`
}

type ActionSuccessResponse struct {
	ResponseBase
	Data *ActionSuccessResponseData `json:"data"`
}