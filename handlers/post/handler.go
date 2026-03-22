package post

import (
	"com.dreamfsk/blog/services/post"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	CreatePost(c *gin.Context)
	GetPostById(c *gin.Context)
	UpdatePost(c *gin.Context)
	GetPostPageList(c *gin.Context)
	DeletePost(c *gin.Context)
}

func NewHandler(db *gorm.DB) Handler {
	return &handler{postService: post.NewService(db)}
}

type handler struct {
	postService post.Service
}

func (h *handler) i() {}
