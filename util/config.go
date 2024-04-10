// SPDX-FileCopyrightText: 2023-2024 Sidings Media
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
	// Should the ticketing system alert on a new message?
	TicketShouldAlert bool
	// Should the ticketing system auto respond on a new message?
	TicketShouldAutorespond bool
	// Source for messages
	TicketSource string
)

// Environment variables
const (
	BindAddrEnv                = "BIND_ADDRESS"
	TrustedProxiesEnv          = "TRUSTED_PROXIES"
	TicketAPIURLEnv            = "TICKET_API_URL"
    TicketHealthURLEnv         = "TICKET_HEALTH_URL"
	TicketAPIKeyEnv            = "TICKET_API_KEY"
	TicketShouldAlertEnv       = "TICKET_SHOULD_ALERT"
	TicketShouldAutorespondEnv = "TICKET_SHOULD_AUTORESPOND"
	TicketSourceEnv            = "TICKET_SOURCE"
)

// Defaults
const (
	DefaultBindAddr                = "[::]:3000"
	DefaultTrustedProxies          = "*"
	DefaultTicketShouldAlert       = true
	DefaultTicketShouldAutorespond = true
	DefaultTicketSource            = "API"
)
