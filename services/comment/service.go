package comment

import (
	"com.dreamfsk/blog/repository/models"
	"gorm.io/gorm"
)

type Service interface {
	i()
	CreateComment(req *CommentCreatReq) (*models.Comment, error)
	GetCommentList(req *CommentPageReq) (*[]models.Comment, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{db: db}
}

func (s service) i() {}
