package email

import (
	"context"
	"fmt"

	"shanepee.com/api/config"
	"shanepee.com/api/service"
)

type debugEmailSenderImpl struct {
	defaultName                   string
	defaultMailAddress            string
	resetPasswordFrontendEndpoint string
}

var _ service.EmailSender = &debugEmailSenderImpl{}

func NewDebug(cfg config.Config) *debugEmailSenderImpl {
	return &debugEmailSenderImpl{
		defaultName:                   cfg.Email.Name,
		defaultMailAddress:            cfg.Email.Address,
		resetPasswordFrontendEndpoint: cfg.ResetPasswordFrontendEndpoint,
	}
}

func (d *debugEmailSenderImpl) SendResetPasswordEmail(ctx context.Context, to string, token string, requestID int64) error {
	resetLink := fmt.Sprintf("%s?token=%s&request_id=%d", d.resetPasswordFrontendEndpoint, token, requestID)

	fmt.Printf("From: %s %s\n", d.defaultName, d.defaultMailAddress)
	fmt.Printf("To: %s\n", to)
	fmt.Printf("Subject: Reset Password\n")
	fmt.Printf("Body: Click the link to reset your password: %s\n", resetLink)

	return nil
}
