package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response business interface response
type Response struct {
	Code    int         `json:"code"`            // status code
	Message string      `json:"message"`         // message
	Data    interface{} `json:"data,omitempty"`  // data
	Error   interface{} `json:"error,omitempty"` // error message
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}
