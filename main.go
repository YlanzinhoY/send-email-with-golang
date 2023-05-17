package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"github.com/go-gomail/gomail"
)

func email() {

	g := gomail.NewDialer("smtp.gmail.com", 587, "your email", "your password")
	m := gomail.NewMessage()

	htmlContent, err := ioutil.ReadFile("./index.html")
	if err != nil {
		log.Fatal(err)
	}

	m.SetHeader("From", "email")
	m.SetHeader("To", "Email")
	m.SetHeader("Subject", "Title")
	m.SetBody("text/html", string(htmlContent))

	if err := g.DialAndSend(m); err != nil {
		log.Fatal(err)
	}

}

func main() {

	email()
	fmt.Println("Send Email.")

}
