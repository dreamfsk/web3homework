package comment

import (
	"com.dreamfsk/blog/repository"
	"com.dreamfsk/blog/repository/post"
	"com.dreamfsk/blog/repository/user"
)

type Comment struct {
	ID        uint                  `json:"id" gorm:"primaryKey"`
	Content   string                `json:"content" gorm:"type:text;not null"`
	UserID    uint                  `json:"user_id" gorm:"not null"`
	PostID    uint                  `json:"post_id" gorm:"not null"`
	Audit     repository.Audit      `gorm:"embedded"`
	CreatedAt repository.CustomTime `json:"created_at"`
	UpdatedAt repository.CustomTime `json:"updated_at"`
	DeletedAt repository.CustomTime `json:"-" gorm:"index"`
	User      user.User             `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Post      post.Post             `json:"post,omitempty" gorm:"foreignKey:PostID"`
}
