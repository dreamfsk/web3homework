package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null;size:100"`
	Password  string         `json:"-" gorm:"not null"`
	CreatedAt CustomTime     `json:"createdAt"`
	UpdatedAt CustomTime     `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Audit     Audit          `gorm:"embedded"`
	PostCount int64          `json:"postCount"`
	Posts     []Post         `json:"posts,omitempty" gorm:"foreignKey:UserID"`
	Comments  []Comment      `json:"comments,omitempty" gorm:"foreignKey:UserID"`
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
