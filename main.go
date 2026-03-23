package main

import (
	"com.dreamfsk/blog/config"
	"com.dreamfsk/blog/middleware"
	"com.dreamfsk/blog/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 加载配置
	cfg := config.Env()
	// 创建 Gin 引擎
	r := gin.Default()
	// 全局中间件
	//r.Use(gin.Logger())
	r.Use(middleware.CORS())

	router.Router(r)

	// 启动服务器
	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
