package user

import (
	"com.dreamfsk/blog/repository/models"
	"com.dreamfsk/blog/repository/repos"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

// CreateUser
func (s *service) CreateUser(req *CreateUserReq) (*models.User, error) {
	// 检查用户名是否已存在
	log.Printf("UserService.CreateUser user: %#v", req)
	repo := repos.NewRepo(s.db)
	u := repo.Model
	u.Username = req.Username
	var er error
	er = repo.CheckNameExists()
	if u.ID != 0 {
		log.Printf("UserService.CreateUser CheckExists err: %#v", er)
		return nil, er
	}
	u.Email = req.Email
	er = repo.CheckEmailExists()
	if u.ID != 0 {
		log.Printf("UserService.CreateUser CheckExists err: %#v", er)
		return nil, er
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// 创建用户
	u.Password = string(hashedPassword)
	id, err := repo.Create()
	if err != nil {
		return nil, err
	}
	log.Printf(" add user successful uId: %s", id)
	return u, nil
}

func (s *service) GetUserByID(id uint) (*models.User, error) {
	log.Printf("UserService.GetUserByID user: %#v", id)
	repo := repos.NewRepo(s.db)
	repo.Model.ID = id
	u, err := repo.GetUser()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return u, nil
}

func (s *service) Authenticate(username, password string) (*models.User, error) {
	log.Printf("UserService.Authenticate username: %v ,password: %v", username, password)
	repo := repos.NewRepo(s.db)
	repo.Model.Username = username
	u := repo.Model
	if err := repo.CheckNameExists(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *service) UpdateUser(id uint, req *UpdateUserReq) (*models.User, error) {
	u, err := s.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// 如果更新邮箱，检查是否已存在
	if req.Email != "" && req.Email == u.Email {
		return nil, err
	}
	u.Email = req.Email
	_, err = repos.NewRepo(s.db, u).Save()
	if err != nil {
		return nil, err
	}
	return u, nil
}
