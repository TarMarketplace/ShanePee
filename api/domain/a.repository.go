package domain

import (
	"context"
)

type ARepository interface {
	FindMany(ctx context.Context) ([]A, error)
	Create(ctx context.Context, body A) error
}
