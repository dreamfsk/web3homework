package comment

import (
	"com.dreamfsk/blog/repository/comment"
	"gorm.io/gorm"
)

type Service interface {
	i()
	CreateComment(req *CommentCreatReq) (*comment.Comment, error)
	GetCommentList(req *CommentPageReq) (*[]comment.Comment, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{db: db}
}

func (s service) i() {}
