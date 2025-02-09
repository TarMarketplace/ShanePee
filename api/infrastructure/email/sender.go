package email

import (
	"errors"

	"shanepee.com/api/config"
	"shanepee.com/api/service"
)

type Provider string

const (
	DebugProvider    Provider = "debug"
	SendgridProvider Provider = "sendgrid"
)

func NewSenderFromConfig(cfg config.Config) (service.EmailSender, error) {
	switch Provider(cfg.Email.Provider) {
	case DebugProvider:
		return NewDebug(cfg), nil
	case SendgridProvider:
		return NewSendgrid(cfg), nil
	default:
		return nil, errors.New("Invalid email provider")
	}
}
