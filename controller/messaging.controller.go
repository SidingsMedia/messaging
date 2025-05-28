// SPDX-FileCopyrightText: 2023-2025 Sidings Media
// SPDX-License-Identifier: MIT

package controller

import (
	"errors"
	"log"
	"net/http"

	internalerrors "github.com/SidingsMedia/messaging/errors"
	"github.com/SidingsMedia/messaging/model"
	"github.com/SidingsMedia/messaging/service"
	"github.com/SidingsMedia/messaging/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MessagingController interface {
	Contact(ctx *gin.Context)
    HealthCheck(ctx *gin.Context)
}

type messagingController struct {
	service service.MessagingService
}

func (controller messagingController) SendMessage(ctx *gin.Context) {
	request := &model.Message{}
	if err := ctx.ShouldBind(request); err != nil && errors.As(err, &validator.ValidationErrors{}){
		util.SendBadRequestFieldNames(ctx, err.(validator.ValidationErrors))
		return
	} else if err != nil {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, model.GeneralError{
            Code: http.StatusBadRequest,
            Message: "Request was malformed",
        })
        return
    }

	if err := controller.service.SendMessage(request); err != nil {
        var e *internalerrors.NameLengthError
        if (errors.As(err, &e)) {
            ctx.AbortWithStatusJSON(http.StatusBadRequest, model.GeneralError{
                Code: http.StatusBadRequest,
                Message: err.Error(),
            })
            return
        }

        log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.GeneralError{
            Code: http.StatusInternalServerError,
            Message: "An unexpected error occurred",
        })
        return
	}
    ctx.Status(http.StatusCreated)
}

func (controller messagingController) HealthCheck(ctx *gin.Context) {
    if err := controller.service.HealthCheck(); err != nil {
        ctx.String(http.StatusServiceUnavailable, "unhealthy")
        ctx.Abort()
        return
    }
    ctx.String(http.StatusOK, "healthy")
}

func NewMessagingController(engine *gin.Engine, messagingService service.MessagingService) {
    controller := &messagingController{
        service: messagingService,
    }
    api := engine.Group("messaging")
    {
        api.POST("contact", controller.SendMessage)
        api.GET("health", controller.HealthCheck)
    }
}
