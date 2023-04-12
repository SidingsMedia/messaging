// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package models

// Standardised error response schema
type GeneralError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
