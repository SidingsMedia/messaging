// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package main

import (
	"log"
	"strings"

	"api.sidingsmedia.com/responses"
	"api.sidingsmedia.com/routes"
	"api.sidingsmedia.com/util"
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

	// Email settings
	util.EmailFrom = util.Mustgetenv(util.EmailFromEnv)
	util.EmailTo = util.Mustgetenv(util.EmailToEnv)
	util.SMTPAddr = util.Mustgetenv(util.SMTPAddrEnv)
	util.SMTPPort = util.IGetenv(util.SMTPPortEnv, util.DefaultSMTPPort)
	util.SMTPUsr = util.Mustgetenv(util.SMTPUsrEnv)
	util.SMTPPwd = util.Mustgetenv(util.SMTPPwdEnv)
}

func main() {
	router := gin.Default()

	// Set trusted proxies. If user has set it to * then we can just
	// ignore it as GIN trusts all by default
	if util.TrustedProxies[0] != "*" {
		if err := router.SetTrustedProxies(util.TrustedProxies); err != nil {
			log.Fatalf("Failed to set trusted proxies. %s", err)
		}
		log.Printf("Trusting the following proxies: %s", util.TrustedProxies)
	}

	// Set our custom 404 handler
	router.NoRoute(func(c *gin.Context) {
		responses.Send404(c)
	})

	log.Println("Registering middlewares")
	router.Use(cors.Default())

	log.Println("Registering routes")
	routes.MessagingRoutes(router)

	log.Printf("Starting server on %s\n", util.BindAddr)
	router.Run(util.BindAddr)
}
