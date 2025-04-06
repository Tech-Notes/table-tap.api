package db

import (
	"context"

	internal "github.com/table-tap/api/internal/types"
)

func (db *DB) GetLastActiveBusinessUserByEmail(ctx context.Context, email string) (*internal.BusinessUser, error) {

	query := `
	WITH permissions AS (
		SELECT 
		p.id AS permission_id,
		COALESCE(ARRAY_AGG(p.name), '{}') AS permissions
		FROM permissions p
		GROUP BY 1
	)
	SELECT bu.id,
	u.email,
	u.password,
	bu.business_id,
	bu.role,
	bu.role_id,
	bu.created_at,
	bu.updated_at,
	p.permissions
	FROM businesses_users bu
	INNER JOIN users u 
	ON bu.user_id = u.id
	INNER JOIN role_permissions rp
	ON bu.role_id = rp.role_id
	LEFT JOIN permissions p
	ON rp.permission_id = p.permission_id
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
