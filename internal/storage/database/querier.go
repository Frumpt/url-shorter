// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"context"
)

type Querier interface {
	DeleteURL(ctx context.Context, arg *DeleteURLParams) error
	GetURL(ctx context.Context, arg *GetURLParams) (string, error)
	SaveURL(ctx context.Context, arg *SaveURLParams) (*Url, error)
}

var _ Querier = (*Queries)(nil)
