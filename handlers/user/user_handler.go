package user

import (
	"com.dreamfsk/blog/config/validation"
	"com.dreamfsk/blog/repository"
	"com.dreamfsk/blog/services/user"
	"com.dreamfsk/blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        uint                  `json:"id"`
	Username  string                `json:"username"`
	Email     string                `json:"email"`
	CreatedAt repository.CustomTime `json:"created_at"`
}

// Register 注册新用户
// @Summary      注册新用户
// @Description  注册新用户
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        username  body     string     true  "用户名"
// @Param        email     body     string     true  "邮箱"
// @Param        password  body     string     true  "密码"
// @Success      200     {object}  Response{data=UserResponse{models.User}}  "成功"
// @Failure      400     {object}  Response  "请求参数错误"
// @Failure      500     {object}  Response  "服务器内部错误"
// @Router       /api/v1/users/register [post]
func (h *handler) Register(c *gin.Context) {
	var req user.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, validation.Error(err))
		return
	}
	u, err := h.userService.CreateUser(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	})
}

// Login 用户登录
// @Summary      用户登录
// @Description  用户登录
// @Tags         用户登录
// @Accept       json
// @Produce      json
// @Param        username    body     string     true  "用户名"
// @Param        password   body     string     true  "密码"
// @Success      200     {object}  Response{data=UserResponse{models.User}}  "成功"
// @Failure      400     {object}  Response  "请求参数错误"
// @Failure      500     {object}  Response  "服务器内部错误"
// @Router       /api/v1/users/login [post]
func (h *handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, validation.Error(err))
		return
	}

	u, err := h.userService.Authenticate(req.Username, req.Password)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	token, err := utils.GenerateToken(h.jwtSecret, u.ID, u.Username)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, gin.H{
		"token": token,
		"user": UserResponse{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
		},
	})
}

// GetProfile 用户查询
// @Summary      用户查询
// @Description  用户查询
// @Tags         用户查询
// @Accept       json
// @Produce      json
// @Param        userID    query     string     true  "用户ID"
// @Success      200     {object}  Response{data=UserResponse{models.User}}  "成功"
// @Failure      400     {object}  Response  "请求参数错误"
// @Failure      500     {object}  Response  "服务器内部错误"
// @Router       /api/v1/users/me [GET]
func (h *handler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	u, err := h.userService.GetUserByID(userID.(uint))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	})
}

// UpdateProfile 用户更新
// @Summary      用户更新
// @Description  用户更新
// @Tags         用户更新
// @Accept       json
// @Produce      json
// @Param        userID    body     string     true  "用户ID"
// @Param        email    body     string     true  "用户邮箱"
// @Success      200     {object}  Response{data=UserResponse{models.User}}  "成功"
// @Failure      400     {object}  Response  "请求参数错误"
// @Failure      500     {object}  Response  "服务器内部错误"
// @Router       /api/v1/users/me [PUT]
func (h *handler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req user.UpdateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, validation.Error(err))
		return
	}

	u, err := h.userService.UpdateUser(userID.(uint), &req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	})
}
