package domain

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id int64, user map[string]any) error
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	FindUserByID(ctx context.Context, id int64) (*User, error)
	FindSellers(ctx context.Context) ([]*User, error)
	FindSellerByID(ctx context.Context, id int64) (*User, error)
	CreatePasswordResetRequest(ctx context.Context, passwordResetRequest *PasswordResetRequest) error
	FindPasswordResetRequestWithUserByID(ctx context.Context, id int64) (*PasswordResetRequest, error)
	UpdateUserPasswordHash(ctx context.Context, id int64, passwordHash string) error
	DeletePasswordResetRequestByID(ctx context.Context, id int64) error
}
