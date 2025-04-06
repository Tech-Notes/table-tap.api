package internal

import "github.com/lib/pq"

type BusinessUser struct {
	ID          int64          `json:"id" db:"id"`
	Email       string         `json:"email" db:"email"`
	Password    string         `json:"password" db:"password"`
	BusinessID  int64          `json:"business_id" db:"business_id"`
	Role        string         `json:"role" db:"role"`
	RoleID      int64          `json:"role_id" db:"role_id"`
	Permissions pq.StringArray `json:"permissions" db:"permissions"`
	CreatedAt   string         `json:"created_at" db:"created_at"`
	UpdatedAt   string         `json:"updated_at" db:"updated_at"`
}
