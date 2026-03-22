package post

import (
	"com.dreamfsk/blog/repository/models"
)

type PostResponse struct {
	ID        uint              `json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	CreatedAt models.CustomTime `json:"createdAt"`
}
type PostUpResponse struct {
	ID        uint              `json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	UpdatedAt models.CustomTime `json:"updatedAt"`
}

type PostDeleteResponse struct {
	ID int `json:"id"`
}
