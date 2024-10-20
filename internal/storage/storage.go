package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"url-sorter/internal/storage/database"
)

type Storage struct {
	*database.Queries
}

func New(cnfStorage string) (*Storage, error) {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, cnfStorage)
	if err != nil {
		return nil, fmt.Errorf("failed pgx connection: %w", err)
	}
	return &Storage{
		database.New(db),
	}, nil
}
