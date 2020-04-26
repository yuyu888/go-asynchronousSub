package log



import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var InfoLogger *zap.Logger
var DebugLogger *zap.Logger

func Init() {

    InfoLogger = NewLogger("./info.log", zapcore.InfoLevel, 128, 30, 7, true, "Info")
    DebugLogger = NewLogger("./debug.log", zapcore.DebugLevel, 128, 30, 7, true, "Debug")
}
