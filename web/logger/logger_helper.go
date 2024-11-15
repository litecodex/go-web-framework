package logger

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	RequestUtil "github.com/litecodex/go-web-framework/web/utils/request"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// 全局静态变量
var logger *zap.Logger = NewConsoleLogger()

func Info(ctx *gin.Context, msg string, fields ...zap.Field) {
	if ctx != nil {
		fields = append(fields, TraceId(ctx))
	}
	logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func Warn(ctx *gin.Context, msg string, fields ...zap.Field) {
	if ctx != nil {
		fields = append(fields, TraceId(ctx))
	}
	logger.WithOptions(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

func Error(ctx *gin.Context, msg string, fields ...zap.Field) {
	if ctx != nil {
		fields = append(fields, TraceId(ctx))
	}
	logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}

func Debug(ctx *gin.Context, msg string, fields ...zap.Field) {
	if ctx != nil {
		fields = append(fields, TraceId(ctx))
	}
	logger.WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

func SetLogger(log *zap.Logger) {
	logger = log
}

func GetLogger() *zap.Logger {
	return logger
}

func TraceId(ctx *gin.Context) zap.Field {
	return zap.String("traceId", RequestUtil.GetTraceId(ctx))
}

func NewConsoleLogger() *zap.Logger {
	// 配置日志编码
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 设置日志级别
	level := zap.NewAtomicLevel()
	level.SetLevel(zap.DebugLevel)
	// 创建多个输出的Core
	var loggerCores []zapcore.Core = make([]zapcore.Core, 0)
	loggerCores = append(loggerCores, consoleCore(encoderConfig, level)) // 输出到控制台
	core := zapcore.NewTee(loggerCores...)
	// 创建Logger
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func consoleCore(encoderConfig zapcore.EncoderConfig, level zap.AtomicLevel) zapcore.Core {
	// 创建编码器
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	return zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)
}

func fileCore(logPath string, encoderConfig zapcore.EncoderConfig, level zap.AtomicLevel) zapcore.Core {
	if logPath == "" {
		logPath = "./_log/app.log"
	}
	// 设置日志文件轮转
	fileEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	fileWriter, _ := rotatelogs.New(
		logPath+".%Y%m%d", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(time.Hour*24*30),    // 保存30天
		rotatelogs.WithRotationTime(time.Hour*24), //切割频率 24小时
	)
	return zapcore.NewCore(fileEncoder, zapcore.AddSync(fileWriter), level)
}
