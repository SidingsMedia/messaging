// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package service

import (
	"fmt"

	"github.com/SidingsMedia/api.sidingsmedia.com/model"
	"github.com/SidingsMedia/api.sidingsmedia.com/util"
	"gopkg.in/gomail.v2"
)

type MessagingService interface {
    SendEmail(message *model.Message) error
}

type messagingService struct {
    smtpServer gomail.SendCloser
}

func (service *messagingService) SendEmail(message *model.Message) error {
        m := gomail.NewMessage()
		m.SetHeader("From", util.EmailFrom)
		m.SetHeader("Reply-To", message.Email)
		m.SetHeader("To", util.EmailTo)
		m.SetHeader("Subject", message.Subject)
		body := fmt.Sprintf(
			"From: %s\n\nMessage:\n\n%s",
			message.Name,
			message.Message,
		)
		m.SetBody("text/plain", body)

        return gomail.Send(service.smtpServer, m)
}

func NewMessagingService(smtpServer gomail.SendCloser) MessagingService {
    return &messagingService{
        smtpServer: smtpServer,
    }
}
