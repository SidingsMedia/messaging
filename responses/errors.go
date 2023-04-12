// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package responses

import (
	"net/http"

	"api.sidingsmedia.com/models"
	"github.com/gin-gonic/gin"
)

func Send400(c *gin.Context) {
  c.JSON(
    http.StatusBadRequest,
    models.GeneralError{
      Code: http.StatusBadRequest,
      Message: "Request body was invalid. Please check the specification for the correct request format.",
    },
  )
}

func Send500(c *gin.Context) {
  c.JSON(
    http.StatusInternalServerError,
    models.GeneralError{
      Code: http.StatusInternalServerError,
      Message: "An unexpected error occured. Please retry the request in a moment.",
    },
  )
}
