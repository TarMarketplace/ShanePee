package email

import (
	"log"

	"shanepee.com/api/config"
	"shanepee.com/api/service"
)

type Provider string

const (
	DebugProvider    Provider = "debug"
	SendgridProvider Provider = "sendgrid"
)

func NewEmailSender(cfg config.Config) service.EmailSender {
	switch Provider(cfg.Email.Provider) {
	case DebugProvider:
		return NewDebug(cfg)
	case SendgridProvider:
		return NewSendgrid(cfg)
	default:
		log.Print("Not found email config, using debug provider")
		return NewDebug(cfg)
	}
}
