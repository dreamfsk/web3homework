package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null;size:100"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null;size:255"`
	Password  string         `json:"-" gorm:"not null"`
	Audit     Audit          `gorm:"embedded"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	PostCount int64          `json:"post_count"`

	// 关联关系
	Posts    []Post    `json:"posts,omitempty" gorm:"foreignKey:UserID"`
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:UserID"`
}

func (a *User) BeforeCreate(tx *gorm.DB) error {
	user := CurrentOperator(tx)
	a.Audit.CreatedBy = user
	a.Audit.UpdatedBy = user
	return nil
}

func (a *User) BeforeUpdate(tx *gorm.DB) error {
	user := CurrentOperator(tx)
	a.Audit.UpdatedBy = user
	return nil
}

func (a *User) BeforeDelete(tx *gorm.DB) error {
	user := CurrentOperator(tx)
	a.Audit.DeletedBy = user
	return nil
}
