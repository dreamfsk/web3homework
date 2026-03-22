package repos

import (
	"com.dreamfsk/blog/repository/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB    *gorm.DB
	Model *models.User
}

func NewRepo(db *gorm.DB, u ...*models.User) *UserRepo {
	if len(u) == 0 {
		return &UserRepo{DB: db, Model: new(models.User)}
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
	var existingUser models.User
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
	var existingUser models.User
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

func (repo *UserRepo) GetUser() (u *models.User, err error) {
	up := repo.Model
	var ur models.User
	err = repo.DB.First(&ur, up.ID).Error
	return &ur, err
}

func (repo *UserRepo) UpdateUser() (u *models.User, err error) {
	up := repo.Model
	var ur models.User
	err = repo.DB.First(&ur, up.ID).Error
	return &ur, err
}

func (repo *UserRepo) UpdateUserByMap(m map[string]any) (u *models.User, err error) {
	up := repo.Model
	au := models.CurrentOperator(repo.DB)
	m["updated_by"] = au
	ur := models.User{}
	err = repo.DB.Model(&ur).Where("id = ?", up.ID).Updates(m).Error
	return &ur, err
}

func (repo *UserRepo) Save() (u *models.User, err error) {
	up := repo.Model
	err = repo.DB.Save(up).Error
	return up, err
}
