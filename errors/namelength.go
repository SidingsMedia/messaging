// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package errors

import "fmt"

type NameLengthError struct {
	Name   string
	MaxLen int
}

func (e *NameLengthError) Error() string {
	return fmt.Sprintf("Name too long. %s (%d) exceeds max length %d.", e.Name, len(e.Name), e.MaxLen)
}

// Length of a users name was too long
func NewNameLengthError(name string, maxLen int) error {
	return &NameLengthError{
		Name:   name,
		MaxLen: maxLen,
	}
}
