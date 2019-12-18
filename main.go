package main

import (
	"flag"
	"os"
	"sendMail/provider"
)

var (
	smtpProviderUser     string
	smtpProviderPassword string

	sendTo string
	title  string
	body   string
)

func main() {
	flag.StringVar(&smtpProviderUser, "spu", "", "smtp provider user")
	flag.StringVar(&smtpProviderPassword, "spp", "", "smtp provider password")
	flag.StringVar(&sendTo, "to", "", "send to someone")
	flag.StringVar(&title, "t", "", "mail title")
	flag.StringVar(&body, "b", "", "mail message body")

	flag.Parse()

	if smtpProviderUser == "" || smtpProviderPassword == "" || sendTo == "" || title == "" {
		flag.Usage()
		os.Exit(1)
	}

	provider.SendGmail(smtpProviderUser, smtpProviderPassword, sendTo, title, body)
}
