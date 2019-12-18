package provider

import (
	"log"
	"net/smtp"
)

func SendGmail(spu, spp, to, title, body string) {
	// spu:smtp provider user
	// spp:smtp provider password
	// to:send to someone
	// title:title for mail,like subject
	// body:message body for mail

	smtpHost := "smtp.gmail.com:587"
	from := spu
	msg := "From: " + spu + "\n" +
		"To: " + to + "\n" +
		"Subject: " + title + "\n\n" +
		body

	err := smtp.SendMail(
		smtpHost,
		smtp.PlainAuth("", spu, spp, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Printf("send mail to %s successful\n", to)
}
