package user

import "com.dreamfsk/blog/repository"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        uint                  `json:"id"`
	Username  string                `json:"username"`
	Email     string                `json:"email"`
	CreatedAt repository.CustomTime `json:"createdAt"`
}
