package comment

import "com.dreamfsk/blog/repository"

type CommentRes struct {
	ID        uint                  `json:"id" gorm:"primaryKey"`
	Content   string                `json:"content" gorm:"type:text;not null"`
	CreatedAt repository.CustomTime `json:"createdAt"`
}
