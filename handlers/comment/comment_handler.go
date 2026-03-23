package comment

import (
	"com.dreamfsk/blog/commons/code"
	"com.dreamfsk/blog/models/common/response"
	"com.dreamfsk/blog/services/comment"
	"github.com/gin-gonic/gin"
)

func (a *handler) NewCommont(c *gin.Context) {
	var req comment.CommentCreatReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(code.Text(code.ParamBindError), c)
		return
	}
	cr, err := a.commontService.CreateComment(&req)
	if err != nil {
		response.FailWithMessage(code.Text(code.CommentCreateError), c)
		return
	}
	response.OkWithDetailed(CommentRes{
		cr.ID,
		cr.Content,
		cr.CreatedAt,
	}, "创建成功", c)
}

func (a *handler) CommentList(c *gin.Context) {
	var req comment.CommentPageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(code.Text(code.ParamBindError), c)
		return
	}
	cs, err := a.commontService.GetCommentList(&req)
	if err != nil {
		response.FailWithMessage(code.Text(code.CommentListError), c)
		return
	}
	var comments = []CommentRes{}
	for _, p := range *cs {
		comments = append(comments, CommentRes{
			p.ID,
			p.Content,
			p.CreatedAt})
	}
	response.OkWithDetailed(comments, "查询成功", c)
}
