package log

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// error logger

var Log *zap.SugaredLogger



type Stack struct {
	zapcore.LevelEnabler
}

func (a Stack) Enabled(Level zapcore.Level) bool {
	return Level >= zapcore.ErrorLevel
}

func Debug(args ...interface{}) {
	Log.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	Log.Debugf(template, args...)
}

func Info(args ...interface{}) {
	Log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	Log.Infof(template, args...)
}

func Warn(args ...interface{}) {
	Log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	Log.Warnf(template, args...)
}

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	Log.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	Log.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	Log.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	Log.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	Log.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	Log.Fatalf(template, args...)

}

func WithCtx(c *gin.Context) *zap.SugaredLogger {

	uri := c.Value("uri").(string)
	return Log.With(
		zap.String("traceId", c.Value("traceId").(string)),
		zap.String("uri",uri))
}

func With(args interface{}) *zap.SugaredLogger {
	return Log.With(args)
}
