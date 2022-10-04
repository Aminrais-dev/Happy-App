package helper

import (
	"bytes"
	"capstone/happyApp/config"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"gopkg.in/gomail.v2"
)

var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL")
var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD")

type BodyEmail struct {
	Name  string
	Title string
	Date  string
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

func SendEmail(to, subject string, data interface{}) error {

	template, errPath := filepath.Abs("./helper/template/notif.html")
	if errPath != nil {
		panic("error get template")
	}
	result, _ := ParseTemplate(template, data)
	m := gomail.NewMessage()
	m.SetHeader("From", CONFIG_AUTH_EMAIL)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", result)
	// m.Attach(templateFile) // attach whatever you want
	d := gomail.NewDialer(config.SMTP_HOST, config.SMTP_PORT, CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD)
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
	return nil
}
