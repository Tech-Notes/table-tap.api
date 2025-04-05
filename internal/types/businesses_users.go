package internal

type BusinessUser struct {
	ID         int    `json:"id" db:"id"`
	Email      string `json:"email" db:"email"`
	BusinessID int    `json:"business_id" db:"business_id"`
	Role       string `json:"role" db:"role"`
	RoleID     int    `json:"role_id" db:"role_id"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}
