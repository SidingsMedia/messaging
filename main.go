// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package main

import (
	"log"

	"api.sidingsmedia.com/routes"
	"api.sidingsmedia.com/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	log.Println("Fetching environment variables")

	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}

	util.BindAddr = util.SGetenv(util.BindAddrEnv, util.DefaultBindAddr)
	util.EmailFrom = util.Mustgetenv(util.EmailFromEnv)
	util.EmailTo = util.Mustgetenv(util.EmailToEnv)
	util.SMTPAddr = util.Mustgetenv(util.SMTPAddrEnv)
	util.SMTPPort = util.IGetenv(util.SMTPPortEnv, util.DefaultSMTPPort)
	util.SMTPUsr = util.Mustgetenv(util.SMTPUsrEnv)
	util.SMTPPwd = util.Mustgetenv(util.SMTPPwdEnv)
}

func main() {
  router := gin.Default()
  router.Use(cors.Default())

  routes.RegisterCommunications(router)

  router.Run(util.BindAddr)
}
