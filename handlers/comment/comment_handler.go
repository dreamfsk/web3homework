package comment

import (
	"com.dreamfsk/blog/config/validation"
	"com.dreamfsk/blog/services/comment"
	"com.dreamfsk/blog/utils"
	"github.com/gin-gonic/gin"
)

func (a *handler) NewCommont(c *gin.Context) {
	var req comment.CommentCreatReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, validation.Error(err))
		return
	}
	cr, err := a.commontService.CreateComment(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, CommentRes{
		cr.ID,
		cr.Content,
		cr.CreatedAt,
	})
	return
}

func (a *handler) CommentList(c *gin.Context) {
	var req comment.CommentPageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, validation.Error(err))
		return
	}
	cs, err := a.commontService.GetCommentList(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	var comments = []CommentRes{}
	for _, p := range *cs {
		comments = append(comments, CommentRes{
			p.ID,
			p.Content,
			p.CreatedAt})
	}
	utils.Success(c, comments)
}
