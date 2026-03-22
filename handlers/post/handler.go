package post

import (
	"com.dreamfsk/blog/services/post"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	NewPost(c *gin.Context)
	GetPost(c *gin.Context)
	RefreshPost(c *gin.Context)
	PostPageList(c *gin.Context)
	PostRemove(c *gin.Context)
}

func NewHandler(db *gorm.DB) Handler {
	return &handler{postService: post.NewService(db)}
}

type handler struct {
	postService post.Service
}

func (h *handler) i() {}
