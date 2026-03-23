package logger

import (
	"com.dreamfsk/blog/config"
	"com.dreamfsk/blog/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"os"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap(db *gorm.DB) (logger *zap.Logger) {
	if ok, _ := utils.PathExists(config.B_Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", config.B_Zap.Director)
		_ = os.Mkdir(config.B_Zap.Director, os.ModePerm)
	}
	levels := config.B_Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := NewZapCore(levels[i], db)
		cores = append(cores, core)
	}
	// 构建基础 logger（错误级别的入库逻辑已在自定义 ZapCore 中处理）
	logger = zap.New(zapcore.NewTee(cores...))
	// 启用 Error 及以上级别的堆栈捕捉，确保 entry.Stack 可用
	opts := []zap.Option{zap.AddStacktrace(zapcore.ErrorLevel)}
	if config.B_Zap.ShowLine {
		opts = append(opts, zap.AddCaller())
	}
	logger = logger.WithOptions(opts...)
	return logger
}
