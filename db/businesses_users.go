package db

import (
	"context"

	"github.com/table-tap/api/internal/types"
)

func (db *DB) GetLastActiveBusinessUserByEmail(ctx context.Context, email string) (*types.BusinessUser, error) {

	query := `
	WITH permissions AS (
		SELECT 
		rp.role_id,
		COALESCE(ARRAY_AGG(p.name), '{}') AS permissions
		FROM role_permissions rp
		INNER JOIN permissions p
		ON rp.permission_id = p.id
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
	INNER JOIN permissions p
	ON bu.role_id = p.role_id
	WHERE email = $1
	ORDER BY updated_at DESC
	LIMIT 1
	`

	businessUser := &types.BusinessUser{}

	err := db.GetContext(ctx, businessUser, query, email)
	if err != nil {
		return nil, err
	}

	return businessUser, nil
}
