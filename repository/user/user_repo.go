package user

import (
	"com.dreamfsk/blog/repository"
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
func (a *User) BeforeCreate(tx *gorm.DB) error {
	user := repository.CurrentOperator(tx)
	a.Audit.CreatedBy = user
	a.Audit.UpdatedBy = user
	return nil
}

func (a *User) BeforeUpdate(tx *gorm.DB) error {
	user := repository.CurrentOperator(tx)
	a.Audit.UpdatedBy = user
	return nil
}

func (a *User) BeforeDelete(tx *gorm.DB) error {
	user := repository.CurrentOperator(tx)
	a.Audit.DeletedBy = user
	return nil
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

func (repo *UserRepo) UpdateUserByMap(m map[string]any) (u *User, err error) {
	up := repo.Model
	au := repository.CurrentOperator(repo.DB)
	m["updated_by"] = au
	ur := User{}
	err = repo.DB.Model(&ur).Where("id = ?", up.ID).Updates(m).Error
	return &ur, err
}

func (repo *UserRepo) Save() (u *User, err error) {
	up := repo.Model
	err = repo.DB.Save(up).Error
	return up, err
}
