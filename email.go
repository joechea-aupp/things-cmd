package main

import (
	"fmt"
	"net/smtp"
)

type Email struct {
	To      []string
	From    string
	Subject string
	Body    string
}

func (e *Email) SendEmail(c Credential, es EmailServer) error {
	if len(e.To) == 0 {
		return fmt.Errorf("no recipient")
	}

	if e.From == "" {
		return fmt.Errorf("no sender")
	}

	if e.Subject == "" {
		return fmt.Errorf("no subject")
	}

	if c.Username == "" || c.Password == "" {
		return fmt.Errorf("no credentials")
	}

	if es.SMTPHost == "" || es.SMTPPort == "" {
		return fmt.Errorf("no email server")
	}

	content := fmt.Sprintf("Subject: %s\r\n\r\n%s", e.Subject, e.Body)
	message := []byte(content)

	auth := smtp.PlainAuth("", e.From, c.Password, es.SMTPHost)
	err := smtp.SendMail(es.SMTPHost+":"+es.SMTPPort, auth, e.From, e.To, message)
	if err != nil {
		return err
	}

	return nil
}
