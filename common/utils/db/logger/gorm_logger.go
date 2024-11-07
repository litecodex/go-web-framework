package logger

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type GormLogger struct {
	Log *zap.Logger
}

func NewGormLogger(log *zap.Logger) *GormLogger {
	return &GormLogger{
		Log: log,
	}
}

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	//日志级别由外部logger配置
	return g
}

func (g GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	g.Log.Info(msg, zap.Any("data", data))
}

func (g GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	g.Log.Warn(msg, zap.Any("data", data))
}

func (g GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	g.Log.Error(msg, zap.Any("data", data))
}

func (g GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	var fields []zap.Field
	fields = append(fields, zap.Int64("rows", rows), zap.Float64("elapsed", float64(elapsed.Nanoseconds())/1e6))
	if err != nil && err.Error() != "record not found" {
		fields = append(fields, zap.String("lineNum", utils.FileWithLineNum()), zap.Error(err))
		g.Log.Error(sql, fields...)
	} else {
		g.Log.Debug(sql, fields...)
	}
}
