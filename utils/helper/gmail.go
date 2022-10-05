package helper

import (
	"bytes"
	"capstone/happyApp/config"
	"fmt"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

// var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL")
// var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD")

type BodyEmail struct {
	Name  string
	Event string
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

func SendEmail(to, subject, name, title, date string) error {

	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatal(err)
	// }

	var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL")
	var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD")

	template := "./utils/helper/template/notif.html"
	result, _ := ParseTemplate(template, BodyEmail{
		Name:  name,
		Event: title,
		Date:  date,
	})

	fmt.Println(result)
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
