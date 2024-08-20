package response

import "github.com/gin-gonic/gin"

type HTTPStatusCode struct {
	StatusCode int         `mapstruct:"status_code"`
	Message    string      `mapstruct:"message"`
	Data       interface{} `mapstruct:"data"`
}

func HTTPStatusCodeSuccess(c *gin.Context, statusCode int, Message string, data any) {
	c.JSON(statusCode, HTTPStatusCode{
		StatusCode: statusCode,
		Message:    Message,
		Data:       data,
	})
}

func HTTPStatusCodeError(c *gin.Context, statusCode int, Message string) {
	c.JSON(statusCode, HTTPStatusCode{
		StatusCode: statusCode,
		Message:    Message,
		Data:       nil,
	})
}
