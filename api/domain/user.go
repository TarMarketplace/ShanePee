package domain

type User struct {
	ID            int64         `json:"id" gorm:"primaryKey" example:"97"`
	Email         string        `json:"email" gorm:"unique;not null" example:"johndoe@example.com"`
	PasswordHash  string        `json:"-"`
	FirstName     *string       `json:"first_name,omitempty" example:"John"`
	LastName      *string       `json:"last_name,omitempty" example:"Doe"`
	Gender        *string       `json:"gender,omitempty" example:"MALE"`
	Tel           *string       `json:"tel,omitempty" example:"0988888888"`
	Address       Address       `gorm:"embedded" json:"address"`
	PaymentMethod PaymentMethod `gorm:"embedded" json:"payment_method"`
	Photo         *string       `json:"photo,omitempty" example:"data:image/png;base64,mfkirjIDSFIj324if..."`
}

type Address struct {
	HouseNo  *string `json:"house_no,omitempty" example:"254"`
	District *string `json:"district,omitempty" example:"Pathumwan"`
	Province *string `json:"province,omitempty" example:"Bangkok"`
	Postcode *string `json:"postcode,omitempty" example:"10330"`
}

type PaymentMethod struct {
	CardNumber *string `json:"card_number,omitempty" example:"4242424242424242"`
	ExpireDate *string `json:"expire_date,omitempty" example:"02/27"`
	CVV        *string `json:"cvv,omitempty" example:"132"`
	CardOwner  *string `json:"card_owner,omitempty" example:"Freddy Mercury"`
}

func NewUser(email string, passwordHash string) *User {
	return &User{
		ID:           GenID(),
		Email:        email,
		PasswordHash: passwordHash,
	}
}
