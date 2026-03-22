package router

import (
	"com.dreamfsk/blog/config"
	"com.dreamfsk/blog/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func Router(r *gin.Engine) {

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		utils.Success(c, gin.H{
			"status": "ok",
		})
	})

	// 初始化数据库
	db, err := config.GetDB()
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	cfg := config.Env()
	setRoutApi(r, db, cfg)
}
