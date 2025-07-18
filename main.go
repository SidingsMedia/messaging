// SPDX-FileCopyrightText: 2023-2025 Sidings Media
// SPDX-License-Identifier: MIT

package main

import (
	"log"
	"strings"

	"github.com/SidingsMedia/messaging/controller"
	"github.com/SidingsMedia/messaging/service"
	"github.com/SidingsMedia/messaging/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	log.Println("Fetching environment variables")

	// This will fail in a docker container. Perhaps we need to check if
	// we are in a container and only run if not.
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}

	// Server settings
	util.BindAddr = util.SGetenv(util.BindAddrEnv, util.DefaultBindAddr)
	util.TrustedProxies = strings.Split(
		util.SGetenv(util.TrustedProxiesEnv, util.DefaultTrustedProxies),
		",",
	)

    // Ticket server settings
    util.TicketAPIURL = util.SMustgetenv(util.TicketAPIURLEnv)
    util.TicketHealthURL = util.SMustgetenv(util.TicketHealthURLEnv)
    util.TicketAPIKey = util.SMustgetenv(util.TicketAPIKeyEnv)
    util.TicketMailboxId = util.IMustGetEnv(util.TicketMailboxIdEnv)
}

func main() {
    messagingService := service.NewMessagingService()

	engine := gin.Default()
    engine.Use(cors.Default())

    controller.NewMessagingController(engine, messagingService)

	// Set trusted proxies. If user has set it to * then we can just
	// ignore it as GIN trusts all by default
	if util.TrustedProxies[0] != "*" {
		if err := engine.SetTrustedProxies(util.TrustedProxies); err != nil {
			log.Fatalf("Failed to set trusted proxies. %s", err)
		}
		log.Printf("Trusting the following proxies: %s", util.TrustedProxies)
	}

	log.Printf("Starting server on %s\n", util.BindAddr)
	engine.Run(util.BindAddr)
}
