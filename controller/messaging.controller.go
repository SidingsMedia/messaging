// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/SidingsMedia/api.sidingsmedia.com/model"
	"github.com/SidingsMedia/api.sidingsmedia.com/service"
	"github.com/SidingsMedia/api.sidingsmedia.com/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MessagingController interface {
	Contact(ctx *gin.Context)
}

type messagingController struct {
	service service.MessagingService
}

func (controller messagingController) SendEmail(ctx *gin.Context) {
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

	if err := controller.service.SendEmail(request); err != nil {
        log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.GeneralError{
            Code: http.StatusInternalServerError,
            Message: "An unexpected error occured",
        })
	}
}

func NewMessagingController(engine *gin.Engine, messagingService service.MessagingService) {
    controller := &messagingController{
        service: messagingService,
    }
    api := engine.Group("messaging")
    {
        api.POST("contact", controller.SendEmail)
    }
}
