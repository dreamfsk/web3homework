package user

import (
	"com.dreamfsk/blog/repository"
	"com.dreamfsk/blog/repository/comment"
	"com.dreamfsk/blog/repository/post"
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
	PostCount int64                 `json:"post_count"`
	Posts     []post.Post           `json:"posts,omitempty" gorm:"foreignKey:UserID"`
	Comments  []comment.Comment     `json:"comments,omitempty" gorm:"foreignKey:UserID"`
}
