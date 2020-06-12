package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InfoLogger info日志类型
var InfoLogger *zap.Logger

// DebugLogger debug日志类型
var DebugLogger *zap.Logger

// Init 初始化方法
func Init() {

	InfoLogger = NewLogger("./info.log", zapcore.InfoLevel, 128, 30, 7, true, "Info")
	DebugLogger = NewLogger("./debug.log", zapcore.DebugLevel, 128, 30, 7, true, "Debug")
}
