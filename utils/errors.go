package utils

import (
	"com.dreamfsk/blog/commons"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppError   app error
type AppError struct {
	Code         int               // status code
	businessCode commons.ErrorCode // status code
	Message      string            // message
	Err          error             // err
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func ServiceError(businessCode commons.ErrorCode, code ...int) *AppError {
	var serveCode int
	if len(code) > 0 {
		serveCode = code[0]
	} else {
		serveCode = http.StatusInternalServerError
	}
	return &AppError{
		Code:         serveCode,
		businessCode: businessCode,
		Message:      businessCode.String(),
	}
}

func HandleError(c *gin.Context, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		Error(c, appErr.Code, fmt.Sprintf("%v:%v", appErr.businessCode.Code(), appErr.Message))
		return
	}

	// 未知错误，记录日志但不暴露给客户端
	Error(c, http.StatusInternalServerError, commons.ApiErr.String())
}
