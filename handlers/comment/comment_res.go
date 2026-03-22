package comment

import (
	"com.dreamfsk/blog/repository/models"
)

type CommentRes struct {
	ID        uint              `json:"id" gorm:"primaryKey"`
	Content   string            `json:"content" gorm:"type:text;not null"`
	CreatedAt models.CustomTime `json:"createdAt"`
}
