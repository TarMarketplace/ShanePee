package domain

import (
	"context"
	"errors"
)

var ErrUserNotFound error = errors.New("user not found")
var ErrPasswordChangeRequestNotFound error = errors.New("password change request not found")

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id int64, user map[string]interface{}) error
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	FindUserByID(ctx context.Context, id int64) (*User, error)
	CreatePasswordChangeRequest(ctx context.Context, passwordChangeRequest *PasswordChangeRequest) error
	FindPasswordChangeRequestWithUserByID(ctx context.Context, id int64) (*PasswordChangeRequest, error)
	UpdateUserPasswordHash(ctx context.Context, id int64, passwordHash string) error
	DeletePasswordChangeRequestByID(ctx context.Context, id int64) error
}
