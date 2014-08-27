package mail

import (
	"errors"
	"fmt"
	"net/smtp"
)

var notifyEmail = []string{""}

//	this has only been tested with gmail.com
func SendMail(subject string, msg string) error {
	auth := smtp.PlainAuth(
		"",
		"some.email@gmail.com",
		"password",
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"some.email@gmail.com",
		notifyEmail,
		[]byte("Subject: "+subject+"\r\n\r\n"+msg),
	)
	if err != nil {
		errorMsg := fmt.Sprintf("Error sending mail - %v - message: %v", err, msg)
		return errors.New(errorMsg)
	}

	return nil
}
