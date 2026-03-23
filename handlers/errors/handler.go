package errors

import (
	"com.dreamfsk/blog/services/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler interface {
	i()
	CreateSysError(c *gin.Context)
	DeleteSysError(c *gin.Context)
	DeleteSysErrorByIds(c *gin.Context)
	UpdateSysError(c *gin.Context)
	FindSysError(c *gin.Context)
	GetSysErrorList(c *gin.Context)
	GetSysErrorSolution(c *gin.Context)
}

type handler struct {
	errService system.Service
	zl         *zap.Logger
}

func NewHandler(db *gorm.DB, zl *zap.Logger) Handler {
	return &handler{
		errService: system.NewService(db),
		zl:         zl,
	}
}
func (h *handler) i() {}
