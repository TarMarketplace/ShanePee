package dto

import "shanepee.com/api/domain"

type RegisterBody struct {
	Email    string `json:"email"    example:"johndoe@example.com"`
	Password string `json:"password" example:"VerySecurePassword"`
}

type UserUpdateBody struct {
	FirstName     *string               `json:"first_name"`
	LastName      *string               `json:"last_name"`
	Gender        *string               `json:"gender"`
	Tel           *string               `json:"tel"`
	Address       *domain.Address       `json:"address"`
	PaymentMethod *domain.PaymentMethod `json:"payment_method"`
}

func (b *RegisterBody) IntoUser() *domain.User {
	return domain.NewUser(b.Email, b.Password)
}

func (b *UserUpdateBody) IntoMap() map[string]any {
	panic("Unimplemented")
}

type LoginBody struct {
	Email    string `json:"email" example:"johndoe@example.com"`
	Password string `json:"password" example:"VerySecurePassword"`
}

type ChangePasswordBody struct {
	RequestID   int64  `json:"request_id"`
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}
