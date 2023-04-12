// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package util

import "regexp"

// IsValidEmail returns true if the email address is valid, false otherwise
func IsValidEmail(email string) bool {
    // A simple regular expression to validate email addresses
    pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
    match, err := regexp.MatchString(pattern, email)
    return err == nil && match
}
