package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPStatusCode struct {
	StatusCode int         `mapstruct:"status_code"`
	Message    string      `mapstruct:"message"`
	Data       interface{} `mapstruct:"data"`
}

func HTTPStatusCodeSuccess(c *gin.Context, statusCode int, data any) {
	c.JSON(http.StatusOK, HTTPStatusCode{
		StatusCode: statusCode,
		Message:    msg[statusCode],
		Data:       data,
	})
}

func HTTPStatusCodeError(c *gin.Context, statusCode int) {
	c.JSON(http.StatusBadRequest, HTTPStatusCode{
		StatusCode: statusCode,
		Message:    msg[statusCode],
		Data:       nil,
	})
}
