package comment

import (
	"com.dreamfsk/blog/services/comment"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler interface {
	i()
	NewCommont(c *gin.Context)
	CommentList(c *gin.Context)
}

func NewHandler(db *gorm.DB, zl *zap.Logger) Handler {
	return &handler{
		commontService: comment.NewService(db), zl: zl,
	}
}

type handler struct {
	commontService comment.Service
	zl             *zap.Logger
}

func (h *handler) i() {}
