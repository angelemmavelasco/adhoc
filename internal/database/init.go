package database

import (
	"context"
	"database/sql"
	"fmt"
)

var schemaSQL string

func InitDB(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, schemaSQL); err != nil {
		return fmt.Errorf("error while initializing database %v", err)
	}
	return nil
}
