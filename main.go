package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

// Your available domain names can be found here:
// (https://app.mailgun.com/app/domains)
var yourDomain string = "***sandbox1b4f79d8ff50472487ffdd638292d793.mailgun.org"

//"sandboxb35142e463314da09b5baa7fd0878ffb.mailgun.org" // e.g. mg.yourcompany.com

// You can find the Private API Key in your Account Menu, under "Settings":
// (https://app.mailgun.com/app/account/security)
var privateAPIKey string = "***47c46f002b5121fb6408d78aab109051-1b3a03f6-07cd747d" //"210edf5ee4c770ebcba833e747bc5eea-1b3a03f6-5e0578d0"

func main() {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	//When you have an EU-domain, you must specify the endpoint:
	//mg.SetAPIBase("https://api.eu.mailgun.net/v3")
	//sender := "duc0103999@gmail.com"
	sender := "hoangkimanhduc01031999@gmail.com"
	//sender := "Mailgun Sandbox<postmaster@sandboxb35142e463314da09b5baa7fd0878ffb.mailgun.org>"
	subject := "Hello Hoang Kim Anh Duc"
	body := "Hello from Mailgun Go! 123456789"
	recipient := "duc0103999@gmail.com"

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
