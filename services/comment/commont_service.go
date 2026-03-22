package comment

import (
	"com.dreamfsk/blog/commons"
	"com.dreamfsk/blog/repository/comment"
	"com.dreamfsk/blog/utils"
)

func (a *service) CreateComment(req *CommentCreatReq) (*comment.Comment, error) {
	c := new(comment.Comment)
	c.PostID = req.PostId
	c.UserID = req.UserId
	c.Content = req.Content
	_, err := comment.NewCommentRepo(a.db, c).CreateComment()
	if err != nil {
		return nil, utils.ServiceError(commons.CommentCreateError)
	}
	return c, nil
}

func (a *service) GetCommentList(req *CommentPageReq) (*[]comment.Comment, error) {
	p := comment.Comment{
		PostID: req.PostId,
	}
	cs, err := comment.NewCommentRepo(a.db, &p).ListComment()
	if err != nil {
		return nil, utils.ServiceError(commons.PostListError)
	}
	return cs, nil
}
