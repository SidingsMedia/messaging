// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package controllers

import (
	"fmt"
	"log"

	"api.sidingsmedia.com/models"
	"api.sidingsmedia.com/responses"
	"api.sidingsmedia.com/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gopkg.in/gomail.v2"
)

var validate = validator.New()

// Endpoint handler to send email. Sends an email using the configured
// SMTP server to the configured EmailTo address from the EmailFrom
// address.
func SendEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var message models.Message

		if err := c.BindJSON(&message); err != nil {
			responses.Send400(c)
			return
		}

		if err := validate.Struct(&message); err != nil {
			responses.Send400(c)
			return
		}

		if !util.IsValidEmail(message.Email) {
			responses.Send400(c)
			return
		}

		m := gomail.NewMessage()
		m.SetHeader("From", util.EmailFrom)
		m.SetHeader("Reply-To", message.Email)
		m.SetHeader("To", util.EmailTo)
		m.SetHeader("Subject", message.Subject)
        // TODO: Use templates for message body instead of simple
        // formatting.
		origin := c.GetHeader("Origin")
		body := fmt.Sprintf(
			"From: %s\nOrigin: %s\n\nMessage:\n\n%s",
			message.Name,
			origin,
			message.Message,
		)
		m.SetBody("text/plain", body)

		d := gomail.NewDialer(
			util.SMTPAddr,
			util.SMTPPort,
			util.SMTPUsr,
			util.SMTPPwd,
		)

		// Send the email
		if err := d.DialAndSend(m); err != nil {
			log.Println(err)
			responses.Send500(c)
			return
		}
	}
}
