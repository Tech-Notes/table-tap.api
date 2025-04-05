package db

import (
	"context"

	internal "github.com/table-tap/api/internal/types"
)

func (db *DB) GetLastActiveBusinessUserByEmail(ctx context.Context, email string) (*internal.BusinessUser, error) {

	query := `
	SELECT bu.id,
	u.email,
	bu.business_id,
	bu.role,
	bu.role_id,
	bu.created_at,
	bu.updated_at
	FROM businesses_users bu
	INNER JOIN users u 
	ON bu.user_id = u.id
	WHERE email = $1
	ORDER BY updated_at DESC
	LIMIT 1
	`

	businessUser := &internal.BusinessUser{}

	err := db.GetContext(ctx, businessUser, query, email)
	if err != nil {
		return nil, err
	}

	return businessUser, nil
}
