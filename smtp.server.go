package smtp

import (
	"errors"
	"net/smtp"
)

// Client ..
// smtp  client
type Client struct {
	Host string
	Port string
}

func (s *Client) address() string {
	return s.Host + ":" + s.Port
}

// Send ..
// send message
func (s *Client) Send(from, password string, to []string, message []byte) error {
	if from == "" {
		return errors.New("Invalid sender email")
	}

	if len(to) < 1 {
		return errors.New("Receiver is required")
	}
	auth := smtp.PlainAuth("", from, password, s.Host)
	return smtp.SendMail(s.address(), auth, from, to, message)
}
