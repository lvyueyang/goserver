package logger

import (
	"github.com/gookit/color"
	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"server/config"
)

func New() {
	var w io.Writer = os.Stdout
	if config.Config.IsProd {
		w = getLogWriter()
	}

	logger := slog.New(slog.NewTextHandler(w, nil))
	slog.SetDefault(logger)
}

func getLogWriter() *lumberjack.Logger {
	fileName := path.Join(config.GetLoggerOutPutPath(), "log.log")
	slog.Debug("日志输出路径：", fileName)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    1,     // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 100,   // 保留旧文件的最大个数
		MaxAge:     30,    // 保留旧文件的最大天数
		Compress:   false, // 是否压缩/归档旧文件
	}
	return lumberJackLogger
}

func Success(msg string, a ...any) {
	slog.Info(SuccessColor(msg), a...)
}

func Info(msg string, a ...any) {
	slog.Info(InfoColor(msg), a...)
}

func Warn(msg string, a ...any) {
	slog.Warn(WarnColor(msg), a...)
}

func Error(msg string, a ...any) {
	slog.Error(ErrorColor(msg), a...)
}

func SuccessColor(a string) string {
	if config.Config.IsProd {
		return a
	}
	return color.Green.Sprint(a)
}

func WarnColor(a string) string {
	if config.Config.IsProd {
		return a
	}
	return color.Yellow.Sprint(a)
}

func InfoColor(a string) string {
	if config.Config.IsProd {
		return a
	}
	return color.Blue.Sprint(a)
}
func ErrorColor(a string) string {
	if config.Config.IsProd {
		return a
	}
	return color.Red.Sprint(a)
}
