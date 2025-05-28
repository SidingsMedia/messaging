// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package domain

type Conversation struct {
	Type      string   `json:"type"` // email, phone, chat
	MailboxId int      `json:"mailboxId"`
	Subject   string   `json:"subject"`
	Customer  Customer `json:"customer"`
	Threads   []Thread `json:"threads"`
}
