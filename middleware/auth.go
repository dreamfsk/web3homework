package middleware

import (
	"com.dreamfsk/blog/commons/code"
	"com.dreamfsk/blog/config"
	"com.dreamfsk/blog/models/common/response"
	"com.dreamfsk/blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

func Auth(jwt *config.JWTConfig, zl *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Header 获取 Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			zl.Error(code.Text(code.AuthorizationNotFound))
			response.NoAuth(code.Text(code.AuthorizationNotFound), c)
			c.Abort()
			return
		}

		// 提取 Token（Bearer <token>）
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			zl.Error(code.Text(code.AuthorizationError))
			response.NoAuth(code.Text(code.AuthorizationError), c)
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 验证 Token
		claims, err := utils.ParseToken(tokenString, []byte(jwt.Secret))
		if err != nil {
			zl.Error(code.Text(code.AuthorizationInvaild))
			response.NoAuth(code.Text(code.AuthorizationInvaild), c)
			c.Abort()
			return
		}

		// 将用户信息存储到 Context
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
