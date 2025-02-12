package domain

type User struct {
	ID            int64         `json:"id" gorm:"primaryKey"`
	Email         string        `json:"email" gorm:"unique;not null"`
	PasswordHash  string        `json:"-" swaggerignore:"true"`
	FirstName     *string       `json:"first_name"`
	LastName      *string       `json:"last_name"`
	Gender        *string       `json:"gender"`
	Tel           *string       `json:"tel"`
	Address       Address       `gorm:"embedded" json:"address"`
	PaymentMethod PaymentMethod `gorm:"embedded" json:"payment_method"`
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

type PasswordChangeRequest struct {
	ID        int64  `gorm:"primaryKey"`
	UserID    int64  `gorm:"not null"`
	TokenHash string `gorm:"not null"`
	User      User   `gorm:"constraint:OnDelete:CASCADE;"`
}

func NewPasswordChangeRequest(tokenHash string, userID int64) *PasswordChangeRequest {
	return &PasswordChangeRequest{
		ID:        GenID(),
		UserID:    userID,
		TokenHash: tokenHash,
	}
}

type UserCreateBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateBody struct {
	FirstName     *string        `json:"first_name"`
	LastName      *string        `json:"last_name"`
	Gender        *string        `json:"gender"`
	Tel           *string        `json:"tel"`
	Address       *Address       `gorm:"embedded" json:"address"`
	PaymentMethod *PaymentMethod `gorm:"embedded" json:"payment_method"`
}

func NewUser(email string, passwordHash string) *User {
	return &User{
		ID:           GenID(),
		Email:        email,
		PasswordHash: passwordHash,
	}
}
