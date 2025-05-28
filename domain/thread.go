// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package domain

type Thread struct {
	Text     string   `json:"text"`
	Type     string   `json:"type"` // message, note, customer
	Customer Customer `json:"customer"`
}
