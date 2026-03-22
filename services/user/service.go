package user

import (
	"com.dreamfsk/blog/repository/models"
	"gorm.io/gorm"
)

type Service interface {
	i()
	CreateUser(req *CreateUserReq) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	Authenticate(username, password string) (*models.User, error)
	UpdateUser(id uint, req *UpdateUserReq) (*models.User, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{db: db}
}

func (s *service) i() {}
