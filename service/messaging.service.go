// SPDX-FileCopyrightText: 2023-2025 Sidings Media
// SPDX-License-Identifier: MIT

package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/SidingsMedia/messaging/domain"
	"github.com/SidingsMedia/messaging/errors"
	"github.com/SidingsMedia/messaging/model"
	"github.com/SidingsMedia/messaging/util"
)

type MessagingService interface {
	SendMessage(message *model.Message) error
	HealthCheck() error
}

type messagingService struct {
}

func (service *messagingService) SendMessage(message *model.Message) error {
	names := strings.Split(message.Name, " ")
	firstName := names[0]
	lastName := ""
	if len(names) > 1 {
		lastName = names[len(names)-1]
	}

	if len(firstName) > 40 {
        return errors.NewNameLengthError(firstName, 40)
	}

	if len(lastName) > 40 {
		return errors.NewNameLengthError(lastName, 40)
	}

	customer := domain.Customer{
		Email:     message.Email,
		FirstName: firstName,
		LastName:  lastName,
	}

	body := domain.Conversation{
		Type:      "email",
		MailboxId: util.TicketMailboxId,
		Subject:   "Query from contact form",
		Customer:  customer,
		Threads: []domain.Thread{{
			Text:     message.Message,
			Type:     "customer",
			Customer: customer,
		}},
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", util.TicketAPIURL, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-FreeScout-API-Key", util.TicketAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("got bad response from API: %d %s", resp.StatusCode, string(body))
	}
	return nil
}

func (service *messagingService) HealthCheck() error {
	resp, err := http.Get(util.TicketHealthURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("got bad response from health check: %d %s", resp.StatusCode, string(body))
	}
	return nil
}

func NewMessagingService() MessagingService {
	return &messagingService{}
}
