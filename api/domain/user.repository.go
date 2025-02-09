package domain

import (
	"context"
	"errors"
)

var ErrUserNotFound error = errors.New("user not found")

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id int64, user map[string]interface{}) error
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	FindUserByID(ctx context.Context, id int64) (*User, error)
	CreatePasswordChangeRequest(ctx context.Context, passwordChangeRequest *PasswordChangeRequest) error
	FindPasswordChangeRequestWithUserByID(ctx context.Context, id int64) (*PasswordChangeRequest, error)
}
