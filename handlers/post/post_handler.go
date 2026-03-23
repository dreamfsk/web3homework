package post

import (
	"com.dreamfsk/blog/commons/code"
	"com.dreamfsk/blog/models/common/response"
	"com.dreamfsk/blog/services/post"
	"github.com/gin-gonic/gin"
)

func (a *handler) NewPost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.NoAuth(code.Text(code.AuthorizationError), c)
		return
	}
	var req post.PostCreateReq
	req.UserId = userID.(uint)
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(code.Text(code.ParamBindError), c)
		return
	}
	ps, err := a.postService.CreatePost(&req)
	if err != nil {
		response.FailWithMessage(code.Text(code.PostCreateError), c)
		return
	}
	response.OkWithDetailed(PostResponse{
		ps.ID,
		ps.Title,
		ps.Content,
		ps.CreatedAt,
	}, "创建成功", c)
}

func (a *handler) PostPageList(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.NoAuth(code.Text(code.AuthorizationError), c)
		return
	}
	var req post.PostPageReq
	req.UserId = userID.(uint)
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(code.Text(code.ParamBindError), c)
		return
	}
	ps, err := a.postService.GetPostPageList(&req)
	if err != nil {
		response.FailWithMessage(code.Text(code.PostListError), c)
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
	response.OkWithDetailed(posts, "查询成功", c)
}

func (a *handler) GetPost(c *gin.Context) {
	var req post.PostSearchReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(code.Text(code.ParamBindError), c)
		return
	}
	ps, err := a.postService.GetPostById(req.ID)
	if err != nil {
		response.FailWithMessage(code.Text(code.PostDetailError), c)
		return
	}
	response.OkWithDetailed(PostResponse{
		ps.ID,
		ps.Title,
		ps.Content,
		ps.CreatedAt,
	}, "查询成功", c)
}

func (a *handler) RefreshPost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.FailWithMessage(code.Text(code.AuthorizationError), c)
		return
	}
	var req post.PostUpdateReq
	req.UserId = userID.(uint)
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(code.Text(code.ParamBindError), c)
		return
	}
	ps, err := a.postService.UpdatePost(&req)
	if err != nil {
		response.FailWithMessage(code.Text(code.PostUpdateError), c)
		return
	}
	response.OkWithDetailed(PostResponse{
		ps.ID,
		ps.Title,
		ps.Content,
		ps.CreatedAt,
	}, "更新成功", c)
}

func (a *handler) PostRemove(c *gin.Context) {
	var req post.PostSearchReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(code.Text(code.ParamBindError), c)
		return
	}
	rs, err := a.postService.DeletePost(req.ID)
	if err != nil {
		response.FailWithMessage(code.Text(code.PostDeleteError), c)
		return
	}
	response.OkWithDetailed(PostDeleteResponse{
		ID: rs,
	}, "删除成功", c)
}
