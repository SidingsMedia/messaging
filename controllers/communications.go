package controllers

import (
	"log"
	"net/http"

	"api.sidingsmedia.com/configs"
	"api.sidingsmedia.com/models"
	"api.sidingsmedia.com/responses"
	"api.sidingsmedia.com/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gopkg.in/gomail.v2"
)

var validate = validator.New()

func SendEmail() gin.HandlerFunc {
  return func(c *gin.Context) {
    var message  models.Message

    if err := c.BindJSON(&message); err != nil {
      c.JSON(
        http.StatusBadRequest,
        responses.GeneralError{
          Code: http.StatusBadRequest,
          Message: "Request body was invalid",
        },
      )
      return
    }

    if err := validate.Struct(&message); err != nil {
      c.JSON(
        http.StatusBadRequest,
        responses.GeneralError{
          Code: http.StatusBadRequest,
          Message: "Request body was invalid",
        },
      )
      return
    }

    if !util.IsValidEmail(message.Email) {
      c.JSON(
        http.StatusBadRequest,
        responses.GeneralError{
          Code: http.StatusBadRequest,
          Message: "Request body was invalid",
        },
      )
      return
    }

    m := gomail.NewMessage()
    m.SetHeader("From", configs.FromAddr)
    m.SetHeader("Reply-To", message.Email)
    m.SetHeader("To", configs.ToAddr)
    m.SetHeader("Subject", message.Subject)
    m.SetBody("text/plain", message.Message)

    d := gomail.NewDialer(
      configs.SMTPServer,
      configs.SMTPPort,
      configs.SMTPUser,
      configs.SMTPPassword,
    )

    // Send the email
    if err := d.DialAndSend(m); err != nil {
      log.Println(err)
      log.Println(configs.SMTPPassword)
        c.JSON(
          http.StatusInternalServerError,
          responses.GeneralError{
            Code: http.StatusInternalServerError,
            Message: "Failed to send email due to unexpected error",
          },
        )
        return
    }
  }
}
