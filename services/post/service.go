package post

import (
	"com.dreamfsk/blog/repository/post"
	"gorm.io/gorm"
)

type Service interface {
	i()
	CreatePost(req *PostCreateReq) (*post.Post, error)
	GetPostById(id uint) (*post.Post, error)
	UpdatePost(req *PostUpdateReq) (*post.Post, error)
	GetPostPageList(req *PostPageReq) (*[]post.Post, error)
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
