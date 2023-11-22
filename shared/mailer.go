package shared

import (
	"log"
	"net/smtp"

	"github.com/jasonyangmh/sayakaya/config"
	"github.com/jasonyangmh/sayakaya/model"
)

type SMTPServer struct {
	Host string
	Port string
}

func (s *SMTPServer) Address() string {
	return s.Host + ":" + s.Port
}

func SendMail(cfg *config.Config, userPromo *model.UserPromo) {
	from := cfg.SMTPEmail
	password := cfg.SMTPPass
	to := []string{userPromo.User.Email}

	smtpServer := SMTPServer{Host: cfg.SMTPHost, Port: cfg.SMTPPort}

	message := []byte(userPromo.Code)

	auth := smtp.PlainAuth("", from, password, smtpServer.Host)

	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		log.Fatalf("failed to send email: %v", err)
		return
	}
}
