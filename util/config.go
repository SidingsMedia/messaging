// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package util

import (
	"log"
	"os"
	"strconv"
)

// Runtime config
var (
	// Address to bind server to
	BindAddr string
	// Proxies from which to trust alternative client IP headers
	TrustedProxies []string
	// Email to send messages from
	EmailFrom string
	// Email to send messages to
	EmailTo string
	// Address or hostname of SMTP server
	SMTPAddr string
	// Port to submit mail on
	SMTPPort int
	// User to log into SMTP server with
	SMTPUsr string
	// Password to use to authenticate with SMTP server
	SMTPPwd string
)

// Environment variables
const (
	BindAddrEnv       = "BIND_ADDRESS"
	TrustedProxiesEnv = "TRUSTED_PROXIES"
	EmailFromEnv      = "EMAIL_FROM"
	EmailToEnv        = "EMAIL_TO"
	SMTPAddrEnv       = "SMTP_ADDRESS"
	SMTPPortEnv       = "SMTP_PORT"
	SMTPUsrEnv        = "SMTP_USER"
	SMTPPwdEnv        = "SMTP_PASSWORD"
)

// Defaults
const (
	DefaultBindAddr       = "[::]:3000"
	DefaultTrustedProxies = "*"
	DefaultSMTPPort       = 587
)

// Get the specified environment variable. If it doesn't exist, return
// the fallback instead
func SGetenv(key string, fallback string) string {
	val := os.Getenv(key)

	if len(val) == 0 {
		return fallback
	} else {
		return val
	}
}

// Get the specified environment variable as an integer. If is doesn't
// exist or cannot be converted to an int, return the fallback instead.
func IGetenv(key string, fallback int) int {
	val := os.Getenv(key)

	if len(val) == 0 {
		return fallback
	} else {
		res, err := strconv.Atoi(val)
		if err != nil {
			return fallback
		} else {
			return res
		}
	}
}

// Attempt to get the environment variable. If it is not set, log error
// and exit
func Mustgetenv(key string) string {
	val := os.Getenv(key)

	if len(val) == 0 {
		log.Fatalln("Failed to get", key, ". Environment variable not set")
	}
	return val
}
