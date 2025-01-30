package mail

import (
	"fmt"

	"github.com/jordan-wright/email"
)

const (
	smtpServerAddress = "localhost:1025"
)

type EmailSender interface {
	SendEmail(subject, content string, to, cc, bcc, attachFiles []string) error
}

type GmailSender struct {
	name             string
	fromEmailAddress string
}

func NewGmailSender(name, fromEmailAddress string) EmailSender {
	return &GmailSender{
		name:             name,
		fromEmailAddress: fromEmailAddress,
	}
}

func (sender *GmailSender) SendEmail(subject, content string, to, cc, bcc, attachFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)
		}
	}

	return e.Send(smtpServerAddress, nil)
}
