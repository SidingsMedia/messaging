// SPDX-FileCopyrightText: 2023-2025 Sidings Media
// SPDX-License-Identifier: MIT

package util

// Runtime config
var (
	// Address to bind server to
	BindAddr string
	// Proxies from which to trust alternative client IP headers
	TrustedProxies []string
	// URL for ticketing system API
	TicketAPIURL string
    // URL to use for ticketing system health checks
    TicketHealthURL string
	// API key for ticketing system
	TicketAPIKey string
	// ID of mailbox to send messages to
    TicketMailboxId int
)

// Environment variables
const (
	BindAddrEnv                = "BIND_ADDRESS"
	TrustedProxiesEnv          = "TRUSTED_PROXIES"
	TicketAPIURLEnv            = "TICKET_API_URL"
    TicketHealthURLEnv         = "TICKET_HEALTH_URL"
	TicketAPIKeyEnv            = "TICKET_API_KEY"
    TicketMailboxIdEnv         = "TICKET_MAILBOX_ID"
)

// Defaults
const (
	DefaultBindAddr                = "[::]:3000"
	DefaultTrustedProxies          = "*"
)
