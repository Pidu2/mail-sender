package mail

import (
	"log"
	"net/smtp"
	"os"

	"golang.org/x/net/context"
)

var APP_MAIL = os.Getenv("APP_MAIL")
var APP_PASSWORD = os.Getenv("APP_PASSWORD")

type Server struct {
}

func (s *Server) SendMessage(ctx context.Context, message *Message) (*Success, error) {
	log.Printf("Received message to send from client: %s", message.Text)
	_, err := send(message.Subj, message.Text, message.Recipients)
	if err != nil {
		return &Success{Success: false}, err
	} else {
		return &Success{Success: true}, nil
	}
}

func send(subj string, text string, recipients []string) (result bool, e error) {
	// Configuration
	from := APP_MAIL
	password := APP_PASSWORD
	to := recipients
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: " + subj + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		text + "\r\n")

	// Create authentication
	log.Printf("from=%s recipients=%v", from, to)
	log.Printf("Authenticating...\n")
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	log.Printf("Sending message...\n")
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil

}
