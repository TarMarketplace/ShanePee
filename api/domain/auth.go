package domain

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
