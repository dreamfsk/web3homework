package utils

import (
	"com.dreamfsk/blog/commons"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

// AppError   app error
type AppError struct {
	Code    commons.ErrorCode // status code
	Message string            // message
	Err     error             // err
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func ApiError(code commons.ErrorCode) *AppError {
	return &AppError{
		Code:    code,
		Message: code.String(),
	}
}

func HandleError(c *gin.Context, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		Error(c, int(appErr.Code), appErr.Message)
		return
	}

	// 未知错误，记录日志但不暴露给客户端
	Error(c, 500, commons.ApiErr.String())
}
