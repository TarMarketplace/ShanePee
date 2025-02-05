package domain

import (
	"context"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *User) error
}
