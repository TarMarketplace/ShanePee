package email

import (
	"context"
	"fmt"

	"shanepee.com/api/config"
	"shanepee.com/api/service"
)

type debugEmailSenderImpl struct {
	defaultName                    string
	defaultMailAddress             string
	changePasswordFrontendEndpoint string
}

var _ service.EmailSender = &debugEmailSenderImpl{}

func NewDebug(cfg config.Config) *debugEmailSenderImpl {
	return &debugEmailSenderImpl{
		defaultName:                    cfg.Email.Name,
		defaultMailAddress:             cfg.Email.Address,
		changePasswordFrontendEndpoint: cfg.ChangePasswordFrontendEndpoint,
	}
}

func (d *debugEmailSenderImpl) SendChangePasswordEmail(ctx context.Context, to string, token string, requestID int64) error {
	resetLink := fmt.Sprintf("%s?token=%s&request_id=%d", d.changePasswordFrontendEndpoint, token, requestID)

	fmt.Printf("From: %s %s\n", d.defaultName, d.defaultMailAddress)
	fmt.Printf("To: %s\n", to)
	fmt.Printf("Subject: Change Password\n")
	fmt.Printf("Body: Click the link to change your password: %s\n", resetLink)

	return nil
}
