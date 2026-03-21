package router

import (
	"com.dreamfsk/blog/config"
	"com.dreamfsk/blog/handlers/user"
	"com.dreamfsk/blog/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setRoutApi(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	// 加载配置
	// User business
	userHandler := user.NewHandler(db, cfg)
	// 公开路由
	public := r.Group("/api/v1")
	{
		public.POST("/users/register", userHandler.Register)
		public.POST("/users/login", userHandler.Login)
	}

	// 需要认证的路由
	protected := r.Group("/api/v1")
	protected.Use(middleware.Auth(&cfg.JWT))
	{
		protected.GET("/users/me", userHandler.GetProfile)
		protected.PUT("/users/me", userHandler.UpdateProfile)
	}

}
