package email

import (
	"context"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"shanepee.com/api/config"
	"shanepee.com/api/service"
)

type senderEmailSenderImpl struct {
	apiKey                         string
	defaultName                    string
	defaultMailAddress             string
	changePasswordFrontendEndpoint string
}

var _ service.EmailSender = &senderEmailSenderImpl{}

func NewSendgrid(cfg config.Config) *senderEmailSenderImpl {
	return &senderEmailSenderImpl{
		apiKey:                         cfg.Email.SendgridAPIKey,
		defaultName:                    cfg.Email.Name,
		defaultMailAddress:             cfg.Email.Address,
		changePasswordFrontendEndpoint: cfg.ChangePasswordFrontendEndpoint,
	}
}

func (s *senderEmailSenderImpl) SendChangePasswordEmail(ctx context.Context, toStr string, token string, requestID int64) error {
	resetLink := fmt.Sprintf("%s?token=%s&request_id=%d", s.changePasswordFrontendEndpoint, token, requestID)
	from := mail.NewEmail(s.defaultName, s.defaultMailAddress)
	subject := "Change Password"
	to := mail.NewEmail("", toStr)
	body := fmt.Sprintf("Click the link to change your password: %s", resetLink)
	message := mail.NewSingleEmailPlainText(from, subject, to, body)
	client := sendgrid.NewSendClient(s.apiKey)
	_, err := client.SendWithContext(ctx, message)
	return err
}
