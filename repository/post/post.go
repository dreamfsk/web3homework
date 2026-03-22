package post

import (
	"com.dreamfsk/blog/repository"
	"com.dreamfsk/blog/repository/comment"
	"com.dreamfsk/blog/repository/user"
)

type Post struct {
	ID            uint                  `json:"id" gorm:"primaryKey"`
	Title         string                `json:"title" gorm:"not null"`
	Content       string                `json:"content" gorm:"type:text;not null"`
	UserID        uint                  `json:"userId" gorm:"not null"`
	Audit         repository.Audit      `gorm:"embedded"`
	CreatedAt     repository.CustomTime `json:"createdAt"`
	UpdatedAt     repository.CustomTime `json:"updatedAt"`
	DeletedAt     repository.CustomTime `json:"-" gorm:"index"`
	CommentStatus string                `json:"commentStatus"`
	User          user.User             `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Comments      []comment.Comment     `json:"comments,omitempty" gorm:"foreignKey:PostID"`
}
