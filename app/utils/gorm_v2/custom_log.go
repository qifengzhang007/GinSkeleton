package gorm_v2

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"goskeleton/app/global/variable"
	"time"
)

func createCustomeGormLog() gormLog.Interface {
	var (
		infoStr      = "[info] %s\n "
		warnStr      = "[warn] %s\n"
		errStr       = "[error] %s\n"
		traceStr     = "[traceStr] %s [%.3fms] [rows:%v] %s\n"
		traceWarnStr = "[traceWarn] %s %s [%.3fms] [rows:%v] %s\n"
		traceErrStr  = "[traceErr] %s %s [%.3fms] [rows:%v] %s\n"
	)
	logConf := gormLog.Config{
		SlowThreshold: time.Millisecond * 200, // 慢 SQL 阈值(sql执行时间超过此时间单位，就会触发gorm自带的日志输出)
		LogLevel:      gormLog.Warn,           // Log level
		Colorful:      false,
	}
	return &logger{
		Writer:       logOutPut{},
		Config:       logConf,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

// 对 gorm v2 自带的日志进行接管，使用本项目统一的日志工具 zaplog 管理
type logOutPut struct{}

func (l logOutPut) Printf(strFormat string, args ...interface{}) {
	logRes := fmt.Sprintf(strFormat, args...)
	variable.ZapLog.Info("gorm_v2 日志:", zap.String("详情：", logRes))
}

type logger struct {
	gormLog.Writer
	gormLog.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode log mode
func (l *logger) LogMode(level gormLog.LogLevel) gormLog.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLog.Info {
		l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLog.Warn {
		l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLog.Error {
		l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (l logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= gormLog.Error:
			sql, rows := fc()
			if rows == -1 {
				l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-1", sql)
			} else {
				l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormLog.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-1", sql)
			} else {
				l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.LogLevel >= gormLog.Info:
			sql, rows := fc()
			if rows == -1 {
				l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-1", sql)
			} else {
				l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}
