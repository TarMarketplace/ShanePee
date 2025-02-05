package domain

import (
	"context"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	UpdateUser(ctx context.Context, id int64, user map[string]interface{}) error
}
