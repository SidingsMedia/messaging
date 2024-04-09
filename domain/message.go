// SPDX-FileCopyrightText: 2024 Sidings Media
// SPDX-License-Identifier: MIT

package domain

// Message recieved from the contact form
type Message struct {
	Alert       bool   `json:"alert"`
	Autorespond bool   `json:"autorespond"`
	Source      string `json:"source"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Subject     string `json:"subject"`
	Message     string `json:"message"`
    Attachments []string `json:"attachments"` // Just a placeholder
}
