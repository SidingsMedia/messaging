// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package responses

import (
	"net/http"

	"api.sidingsmedia.com/models"
	"github.com/gin-gonic/gin"
)

// Send the standard response for a HTTP 400 Bad Request error
func Send400(c *gin.Context) {
  c.JSON(
    http.StatusBadRequest,
    models.GeneralError{
      Code: http.StatusBadRequest,
      Message: "Request body was invalid. Please check the specification for the correct request format.",
    },
  )
}

// Send the standard response for a HTTP 404 Not Found error
func Send404(c *gin.Context) {
  c.JSON(
    http.StatusNotFound,
    models.GeneralError{
      Code: http.StatusNotFound,
      Message: "Requested resource could not be located. Please double check your request.",
    },
  )
}

// Send the standard response for a HTTP 500 Internal Server Error error
func Send500(c *gin.Context) {
  c.JSON(
    http.StatusInternalServerError,
    models.GeneralError{
      Code: http.StatusInternalServerError,
      Message: "An unexpected error occured. Please retry the request in a moment.",
    },
  )
}
