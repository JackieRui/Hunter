package log

import (
	"context"

	"go.uber.org/zap"
)

// L zap提供了两个全局的Logger
// 1. *zap.Logger 调用zap.L()获得
// 2. *zap.SugaredLogger 调用zap.S()获得
// 全局的Logger默认不会记录日志 可以使用ReplaceGlobals(logger *Logger) 将logger设置为全局Logger
// 此处仅用于记录日志
func L(ctx context.Context) *zap.Logger {
	// TODO 功能扩展
	return zap.L()
}
