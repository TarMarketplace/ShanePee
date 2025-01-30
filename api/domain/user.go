package domain

type User struct {
	ID           int64  `json:"id" gorm:"primaryKey"`
	Email        string `json:"email"`
	PasswordHash string `json:"-" swaggerignore:"true"`
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
