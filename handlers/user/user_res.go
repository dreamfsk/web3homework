package user

import (
	"com.dreamfsk/blog/repository/models"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        uint              `json:"id"`
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	CreatedAt models.CustomTime `json:"createdAt"`
}
