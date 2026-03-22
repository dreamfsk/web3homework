package comment

import (
	"com.dreamfsk/blog/services/comment"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler interface {
	i()
	NewCommont(c *gin.Context)
	CommentList(c *gin.Context)
}

func NewHandler(db *gorm.DB) Handler {
	return &handler{
		commontService: comment.NewService(db),
	}
}

type handler struct {
	commontService comment.Service
}

func (h *handler) i() {}
