package domain

type PasswordResetRequest struct {
	ID        int64  `gorm:"primaryKey"`
	UserID    int64  `gorm:"not null"`
	TokenHash string `gorm:"not null"`
	User      User   `gorm:"constraint:OnDelete:CASCADE;"`
}

func NewPasswordResetRequest(tokenHash string, userID int64) *PasswordResetRequest {
	return &PasswordResetRequest{
		ID:        GenID(),
		UserID:    userID,
		TokenHash: tokenHash,
	}
}
