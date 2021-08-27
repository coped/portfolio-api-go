package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"net/http"
)

type message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func Contact(c *gin.Context) {
	var request message
	if err := c.BindJSON(&request); err != nil {
		return
	}

	const (
		Sender     = "no-reply@coped.dev"
		SenderName = "coped.dev API"
		Recipient  = "dennisaaroncope@gmail.com" // Must be verified
		SmtpUser   = "SmtpUser"                  // Replace
		SmtpPass   = "SmtpPass"                  // Replace
		Host       = "email-smtp.us-west-2.amazonaws.com"
		Port       = 587
		Subject    = "Message from https://coped.dev!"
		TextBody   = "This email was sent with Amazon SES using the Gomail package."
		CharSet    = "UTF-8"
	)

	m := gomail.NewMessage()
	m.SetBody("text/plain", TextBody)
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(Sender, SenderName)},
		"To":      {Recipient},
		"Subject": {Subject},
	})

	// Send the email.
	d := gomail.NewPlainDialer(Host, Port, SmtpUser, SmtpPass)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err) // Return error code
	} else {
		fmt.Println("Email sent!") // Return success
	}
}

func Index(c *gin.Context) {
	greeting := map[string]string{
		"message": "Welcome to the https://coped.dev API!",
	}
	c.IndentedJSON(http.StatusOK, greeting)
}
