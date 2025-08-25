package mailer

import (
	"bytes"
	"embed"
	"html/template"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

//go:embed templates/*
var templateFs embed.FS

type Mailer struct {
	dialer *gomail.Dialer
	sender string
	Logger echo.Logger
}
type EmailData struct {
	AppName string
	Subject string
	Meta    interface{}
}

func NewMailer(logger echo.Logger) Mailer {

	mailPort, err := strconv.Atoi(os.Getenv("MAIL_PORT"))

	if err != nil {
		logger.Fatal(err.Error())
	}

	mailHost := os.Getenv("MAIL_HOST")
	mailUser := os.Getenv("MAIL_USERNAME")
	mailPass := os.Getenv("MAIL_PASSWORD")

	mailSender := os.Getenv("MAIL_SENDER")

	dialer := gomail.NewDialer(mailHost, mailPort, mailUser, mailPass)

	return Mailer{
		dialer: dialer,
		sender: mailSender,
		Logger: logger,
	}
}

func (m Mailer) Send(recipient string, templateFile string, data EmailData) error {

	go func(recipient string, templateFile string, data EmailData) {

		absolutePath := "templates/" + templateFile

		templateData, err := template.ParseFS(templateFs, absolutePath)

		if err != nil {
			m.Logger.Fatal(err.Error())
			return
		}

		data.AppName = os.Getenv("APP_NAME")

		subject := new(bytes.Buffer)

		err = templateData.ExecuteTemplate(subject, "subject", data)

		if err != nil {

			m.Logger.Fatal(err.Error())
			return
		}

		htmlBody := new(bytes.Buffer)

		err = templateData.ExecuteTemplate(htmlBody, "htmlBody", data)

		if err != nil {

			m.Logger.Fatal(err.Error())
			return
		}

		err = templateData.Execute(subject, data)

		if err != nil {
			m.Logger.Fatal(err.Error())
			return
		}

		mailMessage := gomail.NewMessage()
		mailMessage.SetHeader("From", m.sender)
		mailMessage.SetHeader("To", recipient)
		mailMessage.SetHeader("Subject", subject.String())
		mailMessage.SetBody("text/html", htmlBody.String())

		err = m.dialer.DialAndSend(mailMessage)

		if err != nil {
			m.Logger.Fatal(err.Error())
			return
		}

	}(recipient, templateFile, data)

	return nil
}
