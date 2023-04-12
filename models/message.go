// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package models


type Message struct {
  Name string `json:"name" validate:"required"`
  Email string `json:"email" validate:"required"`
  Subject string `json:"subject" validate:"required"`
  Message string `json:"message" validate:"required"`
}
