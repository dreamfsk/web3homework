package router

import (
	"com.dreamfsk/blog/config"
	"com.dreamfsk/blog/handlers/comment"
	"com.dreamfsk/blog/handlers/post"
	"com.dreamfsk/blog/handlers/user"
	"com.dreamfsk/blog/middleware"
	"com.dreamfsk/blog/middleware/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func setRoutApi(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	// 加载配置
	// *zap.Logger
	lg := logger.Zap(db)
	zap.ReplaceGlobals(lg)
	// User business
	userHandler := user.NewHandler(db, cfg, lg)
	// 公开路由
	public := r.Group("/api/v1")
	{
		public.POST("/user/register", userHandler.Register)
		public.POST("/user/login", userHandler.Login)
	}

	// 需要认证的路由
	protected := r.Group("/api/v1")
	protected.Use(middleware.Auth(&cfg.JWT))
	{
		protected.GET("/user/me", userHandler.GetProfile)
		protected.PUT("/user/update", userHandler.UpdateProfile)

		postHandler := post.NewHandler(db, lg)
		protected.POST("/post/add", postHandler.NewPost)
		protected.PUT("/post/update", postHandler.RefreshPost)
		protected.POST("/post/page", postHandler.PostPageList)
		protected.GET("/post/get", postHandler.GetPost)
		protected.GET("/post/delete", postHandler.PostRemove)

		commentHandler := comment.NewHandler(db, lg)
		protected.POST("/comment/add", commentHandler.NewCommont)
		protected.POST("/comment/list", commentHandler.CommentList)
	}

}
