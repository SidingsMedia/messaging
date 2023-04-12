// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package routes

import (
	"api.sidingsmedia.com/controllers"
	"github.com/gin-gonic/gin"
)

// Register routes for the messaging endpoints
func MessagingRoutes(router *gin.Engine) {
	router.POST("messaging/contact", controllers.SendEmail())
}
