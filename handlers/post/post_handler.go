package post

import (
	"com.dreamfsk/blog/commons"
	"com.dreamfsk/blog/config/validation"
	"com.dreamfsk/blog/services/post"
	"com.dreamfsk/blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *handler) NewPost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, commons.Unauthorized.String())
		return
	}
	var req post.PostCreateReq
	req.UserId = userID.(uint)
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, validation.Error(err))
		return
	}
	ps, err := a.postService.CreatePost(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, PostResponse{
		ps.ID,
		ps.Title,
		ps.Content,
		ps.CreatedAt,
	})
}

func (a *handler) PostPageList(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, commons.Unauthorized.String())
		return
	}
	var req post.PostPageReq
	req.UserId = userID.(uint)
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, validation.Error(err))
		return
	}
	ps, err := a.postService.GetPostPageList(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	var posts = []PostResponse{}
	for _, p := range *ps {
		posts = append(posts, PostResponse{
			p.ID,
			p.Title,
			p.Content,
			p.CreatedAt})
	}
	utils.Success(c, posts)
}

func (a *handler) GetPost(c *gin.Context) {
	ID, exists := c.Get("ID")
	if !exists {
		utils.Error(c, http.StatusInternalServerError, commons.PostNotFoundError.String())
		return
	}
	ps, err := a.postService.GetPostById(ID.(uint))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, PostResponse{
		ps.ID,
		ps.Title,
		ps.Content,
		ps.CreatedAt,
	})
}

func (a *handler) RefreshPost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, commons.Unauthorized.String())
		return
	}
	var req post.PostUpdateReq
	req.UserId = userID.(uint)
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, validation.Error(err))
		return
	}
	ps, err := a.postService.UpdatePost(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, PostUpResponse{
		ps.ID,
		ps.Title,
		ps.Content,
		ps.UpdatedAt,
	})
}

func (a *handler) PostRemove(c *gin.Context) {
	ID, exists := c.Get("ID")
	if !exists {
		utils.Error(c, http.StatusInternalServerError, commons.PostNotFoundError.String())
		return
	}
	rs, err := a.postService.DeletePost(ID.(uint))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, PostDeleteResponse{
		ID: rs,
	})
}
