package post

import (
	"com.dreamfsk/blog/repository/models"
	"gorm.io/gorm"
)

type Service interface {
	i()
	CreatePost(req *PostCreateReq) (*models.Post, error)
	GetPostById(id uint) (*models.Post, error)
	UpdatePost(req *PostUpdateReq) (*models.Post, error)
	GetPostPageList(req *PostPageReq) (*[]models.Post, error)
	DeletePost(id uint) (int, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (p *service) i() {
}
