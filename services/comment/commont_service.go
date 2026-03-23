package comment

import (
	"com.dreamfsk/blog/repository/models"
	"com.dreamfsk/blog/repository/repos"
)

func (a *service) CreateComment(req *CommentCreatReq) (*models.Comment, error) {
	c := new(models.Comment)
	c.PostID = req.PostId
	c.UserID = req.UserId
	c.Content = req.Content
	_, err := repos.NewCommentRepo(a.db, c).CreateComment()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (a *service) GetCommentList(req *CommentPageReq) (*[]models.Comment, error) {
	p := models.Comment{
		PostID: req.PostId,
	}
	cs, err := repos.NewCommentRepo(a.db, &p).ListComment()
	if err != nil {
		return nil, err
	}
	return cs, nil
}
