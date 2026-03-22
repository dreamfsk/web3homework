package post

import (
	"com.dreamfsk/blog/config/validation"
	"com.dreamfsk/blog/services/post"
	"com.dreamfsk/blog/utils"
	"github.com/gin-gonic/gin"
)

func (a *handler) CreatePost(c *gin.Context) {
	var req post.PostCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, validation.Error(err))
		return
	}

	return
}

func (a *handler) GetPostPageList(c *gin.Context) {

	return
}

func (a *handler) GetPostById(c *gin.Context) {

	return
}

func (a *handler) UpdatePost(c *gin.Context) {

	return
}

func (a *handler) DeletePost(c *gin.Context) {

	return
}
