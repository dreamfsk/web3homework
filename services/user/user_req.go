package user

type CreateUserReq struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserReq struct {
	Email string `json:"email" binding:"omitempty,email"`
}
