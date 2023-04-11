// SPDX-FileCopyrightText: 2023 Sidinggs Media
// SPDX-License-Identifier: MIT

package routes

import (
	"api.sidingsmedia.com/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterCommunications(router *gin.Engine) {
  router.POST("contact", controllers.SendEmail())
}
