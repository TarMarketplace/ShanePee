package domain

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id int64, user map[string]any) error
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	FindUserByID(ctx context.Context, id int64) (*User, error)
	CreatePasswordChangeRequest(ctx context.Context, passwordChangeRequest *PasswordChangeRequest) error
	FindPasswordChangeRequestWithUserByID(ctx context.Context, id int64) (*PasswordChangeRequest, error)
	UpdateUserPasswordHash(ctx context.Context, id int64, passwordHash string) error
	DeletePasswordChangeRequestByID(ctx context.Context, id int64) error
}
