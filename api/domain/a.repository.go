package domain

import (
	"context"
)

type ARepository interface {
	FindMany(ctx context.Context) ([]A, error)
	FindOne(ctx context.Context, id int64) (*A, error)
	Create(ctx context.Context, body A) error
	Update(ctx context.Context, id int64, body map[string]interface{}) error
	Delete(ctx context.Context, id int64) error
}
