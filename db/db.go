package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
	EncryptKey *[32]byte
}

func (db *DB) InsertxContext(ctx context.Context, query string, arg interface{}) (int64, error) {
	namedStmt, err := db.PrepareNamedContext(ctx, query)
	if err != nil {
		return 0, err
	}

	var id int64
	err = namedStmt.GetContext(ctx, &id, arg)
	if err != nil {
		return 0, err
	}

	return id, nil
}