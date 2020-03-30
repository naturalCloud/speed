package log

import (
	"github.com/gin-gonic/gin"
	"github.com/syyongx/php2go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// error logger

var Log *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

type stack struct {
	zapcore.LevelEnabler
}

func (a stack) Enabled(Level zapcore.Level) bool {
	return Level >= zapcore.ErrorLevel
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func InitLog() *zap.Logger {
	s, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	appPath := php2go.Substr(s, 0, strings.Index(s, "speed")+5)

	fileName := appPath + "/storage/logs/zap.log"
	level := getLoggerLevel("info")
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(i time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(i.Format("2006-01-02 15:04:05.000"))
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(stack{}))
	Log = logger.Sugar()
	return logger
	//Log1 = logger //嵌套结构化
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
