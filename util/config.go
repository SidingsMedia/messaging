// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package util

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
