package domain

type User struct {
	Email        string `json:"email"`
	PasswordHash []byte `json:"-" swaggerignore:"true"`
}

type UserCreateBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(email string, passwordHash []byte) *User {
	return &User{
		email,
		passwordHash,
	}
}
