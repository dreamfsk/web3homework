package user

import (
	"com.dreamfsk/blog/repository"
	"gorm.io/gorm"
)

type User struct {
	ID        uint                  `json:"id" gorm:"primaryKey"`
	Username  string                `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Email     string                `json:"email" gorm:"uniqueIndex;not null;size:100"`
	Password  string                `json:"-" gorm:"not null"`
	CreatedAt repository.CustomTime `json:"created_at"`
	UpdatedAt repository.CustomTime `json:"updated_at"`
	DeletedAt gorm.DeletedAt        `json:"-" gorm:"index"`
	Audit     repository.Audit      `gorm:"embedded"`
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
