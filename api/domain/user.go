package domain

type User struct {
	ID            int64         `json:"id" gorm:"primaryKey"`
	Email         string        `json:"email"`
	PasswordHash  string        `json:"-" swaggerignore:"true"`
	FirstName     *string       `json:"first_name"`
	LastName      *string       `json:"last_name"`
	Gender        *string       `json:"gender"`
	Tel           *string       `json:"tel"`
	Address       Address       `gorm:"embedded;embeddedPrefix:address_" json:"address"`
	PaymentMethod PaymentMethod `gorm:"embedded;embeddedPrefix:payment_" json:"payment_method"`
}

type Address struct {
	HouseNo  *string `json:"house_no"`
	District *string `json:"district"`
	Province *string `json:"province"`
	Postcode *string `json:"postcode"`
}

type PaymentMethod struct {
	CardNumber *string `json:"card_number"`
	ExpireDate *string `json:"expire_date"`
	CVV        *string `json:"cvv"`
	CardOwner  *string `json:"card_owner"`
}

type UserCreateBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(email string, passwordHash string) *User {
	return &User{
		ID:           GenID(),
		Email:        email,
		PasswordHash: passwordHash,
	}
}
