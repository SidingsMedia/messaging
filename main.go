// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package main

import (
	"api.sidingsmedia.com/configs"
	"api.sidingsmedia.com/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  configs.LoadRuntime()
  router.Use(cors.Default())

  routes.RegisterCommunications(router)

  router.Run(configs.BindAddr)
}
