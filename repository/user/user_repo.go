package user

import (
	"gorm.io/gorm"
)

type UserRepo struct {
	DB    *gorm.DB
	Model *User
}

func NewRepo(db *gorm.DB, u ...*User) *UserRepo {
	if len(u) == 0 {
		return &UserRepo{DB: db, Model: new(User)}
	}
	return &UserRepo{DB: db, Model: u[0]}
}

func (repo *UserRepo) Create() (ID uint, err error) {
	u := repo.Model
	err = repo.DB.Create(u).Error
	return u.ID, err
}

func (repo *UserRepo) CheckNameExists() error {
	u := repo.Model
	var existingUser User
	var err error
	if u.Username != "" {
		err = repo.DB.Where("username = ?", u.Username).First(&existingUser).Error
		if err == nil && existingUser.ID != 0 {
			u.ID = existingUser.ID
			u.Username = existingUser.Username
			u.Password = existingUser.Password
		}
	}
	return err
}
func (repo *UserRepo) CheckEmailExists() error {
	u := repo.Model
	var existingUser User
	var err error
	if u.Email != "" {
		// 检查邮箱是否已存在
		err = repo.DB.Where("email = ?", u.Email).First(&existingUser).Error
		if err == nil && existingUser.ID != 0 {
			u.ID = existingUser.ID
			u.Username = existingUser.Username
			u.Password = existingUser.Password
		}
	}
	return err
}

func (repo *UserRepo) GetUser() (u *User, err error) {
	up := repo.Model
	var ur User
	err = repo.DB.First(&ur, up.ID).Error
	return &ur, err
}

func (repo *UserRepo) UpdateUser() (u *User, err error) {
	up := repo.Model
	var ur User
	err = repo.DB.First(&ur, up.ID).Error
	return &ur, err
}

func (repo *UserRepo) Save() (u *User, err error) {
	up := repo.Model
	err = repo.DB.Save(up).Error
	return up, err
}
