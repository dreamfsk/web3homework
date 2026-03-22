package post

import "com.dreamfsk/blog/repository"

type PostResponse struct {
	ID        uint                  `json:"id"`
	Title     string                `json:"title"`
	Content   string                `json:"content"`
	CreatedAt repository.CustomTime `json:"createdAt"`
}
type PostUpResponse struct {
	ID        uint                  `json:"id"`
	Title     string                `json:"title"`
	Content   string                `json:"content"`
	UpdatedAt repository.CustomTime `json:"updatedAt"`
}

type PostDeleteResponse struct {
	ID int `json:"id"`
}
