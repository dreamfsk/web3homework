package user

import (
	"com.dreamfsk/blog/repository/user"
	"gorm.io/gorm"
)

type Service interface {
	i()
	CreateUser(req *CreateUserReq) (*user.User, error)
	GetUserByID(id uint) (*user.User, error)
	Authenticate(username, password string) (*user.User, error)
	UpdateUser(id uint, req *UpdateUserReq) (*user.User, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{db: db}
}

func (s *service) i() {}
