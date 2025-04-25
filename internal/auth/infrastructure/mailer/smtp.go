package mailer

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

type MailerService struct {
	Host     string
	Port     int
	Username string
	Password string
}

func New(host, username, password string, port int) *MailerService {
	return &MailerService{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func (s *MailerService) Send(to, subject, body string) error {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         s.Host,
	}

	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, s.Host)
	if err != nil {
		return err
	}
	defer client.Close()

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	if err := client.Auth(auth); err != nil {
		return err
	}

	if err := client.Mail(s.Username); err != nil {
		return err
	}

	if err := client.Rcpt(to); err != nil {
		return err
	}

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))
	writer, err := client.Data()
	if err != nil {
		return err
	}

	n, err := writer.Write(msg)
	if err != nil {
		return err
	}

	log.Println("bytes written: ", n)
	writer.Close()
	return nil
}
