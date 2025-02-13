package dto

import "shanepee.com/api/domain"

type RegisterBody struct {
	Email    string `json:"email"    example:"johndoe@example.com"`
	Password string `json:"password" example:"VerySecurePassword"`
}

type PartialAddress struct {
	HouseNo  *string `json:"house_no,omitempty" example:"254"`
	District *string `json:"district,omitempty" example:"Pathumwan"`
	Province *string `json:"province,omitempty" example:"Bangkok"`
	Postcode *string `json:"postcode,omitempty" example:"10330"`
}

type PartialPaymentMethod struct {
	CardNumber *string `json:"card_number,omitempty" example:"4242424242424242"`
	ExpireDate *string `json:"expire_date,omitempty" example:"02/27"`
	CVV        *string `json:"cvv,omitempty" example:"132"`
	CardOwner  *string `json:"card_owner,omitempty" example:"Freddy Mercury"`
}

type UserUpdateBody struct {
	FirstName     *string               `json:"first_name,omitempty" example:"John"`
	LastName      *string               `json:"last_name,omitempty" example:"Doe"`
	Gender        *string               `json:"gender,omitempty" enum:"MALE,FEMALE,OTHER" example:"MALE"`
	Tel           *string               `json:"tel,omitempty" example:"0988888888"`
	Address       *PartialAddress       `json:"address,omitempty"`
	PaymentMethod *PartialPaymentMethod `json:"payment_method,omitempty"`
}

func (b *RegisterBody) IntoUser() *domain.User {
	return domain.NewUser(b.Email, b.Password)
}

func (b *UserUpdateBody) IntoMap() map[string]any {
	result := make(map[string]any)
	if b.FirstName != nil {
		result["first_name"] = b.FirstName
	}
	if b.LastName != nil {
		result["last_name"] = b.LastName
	}
	if b.Gender != nil {
		result["gender"] = b.Gender
	}
	if b.Tel != nil {
		result["tel"] = b.Tel
	}
	if b.Address != nil {
		if b.Address.HouseNo != nil {
			result["house_no"] = b.Address.HouseNo
		}
		if b.Address.District != nil {
			result["district"] = b.Address.District
		}
		if b.Address.Province != nil {
			result["province"] = b.Address.Province
		}
		if b.Address.Postcode != nil {
			result["postcode"] = b.Address.Postcode
		}
	}
	if b.PaymentMethod != nil {
		if b.PaymentMethod.CardNumber != nil {
			result["card_number"] = b.PaymentMethod.CardNumber
		}
		if b.PaymentMethod.ExpireDate != nil {
			result["expire_date"] = b.PaymentMethod.ExpireDate
		}
		if b.PaymentMethod.CVV != nil {
			result["cvv"] = b.PaymentMethod.CVV
		}
		if b.PaymentMethod.CardOwner != nil {
			result["card_owner"] = b.PaymentMethod.CardOwner
		}
	}
	return result
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

type RequestPasswordChangeBody struct {
	Email string `json:"email" example:"johndoe@example.com"`
}
