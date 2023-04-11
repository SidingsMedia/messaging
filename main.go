// SPDX-FileCopyrightText: 2023 Sidinggs Media
// SPDX-License-Identifier: MIT

package main

import (
	"api.sidingsmedia.com/configs"
	"api.sidingsmedia.com/routes"
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  configs.LoadRuntime()

  routes.RegisterCommunications(router)

  router.Run(configs.BindAddr)
}
