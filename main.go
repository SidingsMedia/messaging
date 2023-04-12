// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package main

import (
	"log"
	"strings"

	"github.com/SidingsMedia/api.sidingsmedia.com/controller"
	"github.com/SidingsMedia/api.sidingsmedia.com/service"
	"github.com/SidingsMedia/api.sidingsmedia.com/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
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
    smtpServer, err := InitialiseSMTP(&SMTPConfig{
        Host: util.SMTPAddr,
        Port: util.SMTPPort,
        User: util.SMTPUsr,
        Password: util.SMTPPwd,
    })

    if err != nil {
        log.Fatalf("Failed to connect to SMTP server. %s\n", err)
    }

    messagingService := service.NewMessagingService(smtpServer)

	engine := gin.Default()
    engine.Use(cors.Default())
    // Set our custom 404 handler
	// engine.NoRoute(func(c *gin.Context) {
	// 	responses.Send404(c)
	// })

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

func InitialiseSMTP(config *SMTPConfig) (gomail.SendCloser, error) {
    var s gomail.SendCloser
    var err error

    d := gomail.NewDialer(
        config.Host, config.Port,
        config.User, config.Password,
    )

    if s, err = d.Dial(); err != nil {
        return nil, err
    }

    return s, nil
}

type SMTPConfig struct {
	Host  string
	Port     int
	User     string
	Password string
}
