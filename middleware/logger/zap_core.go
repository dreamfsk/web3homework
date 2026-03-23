package logger

import (
	"com.dreamfsk/blog/config"
	"com.dreamfsk/blog/repository/models"
	"com.dreamfsk/blog/services/system"
	"com.dreamfsk/blog/utils/ast"
	"com.dreamfsk/blog/utils/stacktrace"
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
)

type ZapCore struct {
	level zapcore.Level
	zapcore.Core
	db *gorm.DB
}

func NewZapCore(level zapcore.Level, db *gorm.DB) *ZapCore {
	entity := &ZapCore{level: level, db: db}
	syncer := entity.WriteSyncer()
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})
	entity.Core = zapcore.NewCore(config.B_Zap.Encoder(), syncer, levelEnabler)
	return entity
}

func (z *ZapCore) WriteSyncer(formats ...string) zapcore.WriteSyncer {
	cutter := NewCutter(
		config.B_Zap.Director,
		z.level.String(),
		config.B_Zap.RetentionDay,
		CutterWithLayout(time.DateOnly),
		CutterWithFormats(formats...),
	)
	if config.B_Zap.LogInConsole {
		multiSyncer := zapcore.NewMultiWriteSyncer(os.Stdout, cutter)
		return zapcore.AddSync(multiSyncer)
	}
	return zapcore.AddSync(cutter)
}

func (z *ZapCore) Enabled(level zapcore.Level) bool {
	return z.level == level
}

func (z *ZapCore) With(fields []zapcore.Field) zapcore.Core {
	return z.Core.With(fields)
}

func (z *ZapCore) Check(entry zapcore.Entry, check *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if z.Enabled(entry.Level) {
		return check.AddCore(entry, z)
	}
	return check
}

func (z *ZapCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	for i := 0; i < len(fields); i++ {
		if fields[i].Key == "business" || fields[i].Key == "folder" || fields[i].Key == "directory" {
			syncer := z.WriteSyncer(fields[i].String)
			z.Core = zapcore.NewCore(config.B_Zap.Encoder(), syncer, z.level)
		}
	}
	// 先写入原日志目标
	err := z.Core.Write(entry, fields)

	// 捕捉 Error 及以上级别日志并入库，且可提取 zap.Error(err) 的错误内容
	if entry.Level >= zapcore.ErrorLevel {
		// 避免与 GORM zap 写入互相递归：跳过由 gorm logger writer 触发的日志
		if strings.Contains(entry.Caller.File, "gorm_logger_writer.go") {
			return err
		}

		form := "后端"
		level := entry.Level.String()
		// 生成基础信息
		info := entry.Message

		// 提取 zap.Error(err) 内容
		var errStr string
		for i := 0; i < len(fields); i++ {
			f := fields[i]
			if f.Type == zapcore.ErrorType || f.Key == "error" || f.Key == "err" {
				if f.Interface != nil {
					errStr = fmt.Sprintf("%v", f.Interface)
				} else if f.String != "" {
					errStr = f.String
				}
				break
			}
		}
		if errStr != "" {
			info = fmt.Sprintf("%s | 错误: %s", info, errStr)
		}

		// 附加来源与堆栈信息
		if entry.Caller.File != "" {
			info = fmt.Sprintf("%s \n 源文件:%s:%d", info, entry.Caller.File, entry.Caller.Line)
		}
		stack := entry.Stack
		if stack != "" {
			info = fmt.Sprintf("%s \n 调用栈：%s", info, stack)
			// 解析最终业务调用方，并提取其方法源码
			if frame, ok := stacktrace.FindFinalCaller(stack); ok {
				fnName, fnSrc, sLine, eLine, exErr := ast.ExtractFuncSourceByPosition(frame.File, frame.Line)
				if exErr == nil {
					info = fmt.Sprintf("%s \n 最终调用方法:%s:%d (%s lines %d-%d)\n----- 产生日志的方法代码如下 -----\n%s", info, frame.File, frame.Line, fnName, sLine, eLine, fnSrc)
				} else {
					info = fmt.Sprintf("%s \n 最终调用方法:%s:%d (%s) | extract_err=%v", info, frame.File, frame.Line, fnName, exErr)
				}
			}
		}

		// 使用后台上下文，避免依赖 gin.Context
		ctx := context.Background()
		_ = system.NewService(z.db).CreateSysError(ctx, &models.SysError{
			Form:  &form,
			Info:  &info,
			Level: level,
		})
	}
	return err
}

func (z *ZapCore) Sync() error {
	return z.Core.Sync()
}
