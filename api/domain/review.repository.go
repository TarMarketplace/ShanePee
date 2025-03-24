package domain

import (
	"context"
)

type ReviewRepository interface {
	CreateReview(ctx context.Context, review *Review) error
}
