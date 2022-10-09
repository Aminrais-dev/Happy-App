package helper

import (
	"bytes"
	"capstone/happyApp/config"
	"fmt"
	"html/template"
	"os"
	"runtime"

	"gopkg.in/gomail.v2"
)

type BodyEmail struct {
	Name  string
	Event string
	Date  string
	Url   string
}

type Email struct {
	Name   string
	Status string
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}
	return buf.String(), nil
}

func SendEmail(to, subject, template string, data interface{}) error {

	var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL")
	var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD")

	result, _ := ParseTemplate(template, data)

	m := gomail.NewMessage()
	m.SetHeader("From", CONFIG_AUTH_EMAIL)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", result)
	d := gomail.NewDialer(config.SMTP_HOST, config.SMTP_PORT, CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD)

	err := d.DialAndSend(m)

	return err
}

func SendEmailNotif(to, subject string, data interface{}) {

	template := "./utils/helper/template/notif.html"

	runtime.GOMAXPROCS(10)
	go SendEmail(to, subject, template, data)
}

func SendEmailVerify(to, subject string, data interface{}) {

	template := "./utils/helper/template/verify.html"

	runtime.GOMAXPROCS(10)
	go SendEmail(to, subject, template, data)

}

func SendEmailTransNotif(to, subject string, data interface{}) {

	template := "./utils/helper/template/transnotif.html"

	runtime.GOMAXPROCS(10)
	go SendEmail(to, subject, template, data)

}
