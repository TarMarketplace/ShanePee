package domain

type User struct {
	ID            int64         `json:"id" gorm:"primaryKey" example:"97"`
	Email         string        `json:"email" gorm:"unique;not null" example:"johndoe@example.com"`
	PasswordHash  string        `json:"-"`
	FirstName     *string       `json:"first_name" example:"John"`
	LastName      *string       `json:"last_name" example:"Doe"`
	Gender        *string       `json:"gender" example:"MALE"`
	Tel           *string       `json:"tel" example:"0988888888"`
	Address       Address       `gorm:"embedded" json:"address"`
	PaymentMethod PaymentMethod `gorm:"embedded" json:"payment_method"`
	Photo         *string       `json:"photo"`
}

type Address struct {
	HouseNo  *string `json:"house_no" example:"254"`
	District *string `json:"district" example:"Pathumwan"`
	Province *string `json:"province" example:"Bangkok"`
	Postcode *string `json:"postcode" example:"10330"`
}

type PaymentMethod struct {
	CardNumber *string `json:"card_number" example:"4242424242424242"`
	ExpireDate *string `json:"expire_date" example:"02/27"`
	CVV        *string `json:"cvv" example:"132"`
	CardOwner  *string `json:"card_owner" example:"Freddy Mercury"`
}

func NewUser(email string, passwordHash string) *User {
	return &User{
		ID:           GenID(),
		Email:        email,
		PasswordHash: passwordHash,
	}
}
