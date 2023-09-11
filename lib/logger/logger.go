package logger

import (
	"github.com/gookit/color"
	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"server/config"
)

func New() {
	var handler slog.Handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(0),
	})

	if config.Config.IsProd {
		handler = slog.NewJSONHandler(getLogWriter(), nil)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func getLogWriter() *lumberjack.Logger {
	fileName := path.Join(config.GetLoggerOutPutPath(), "log.log")
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

func SuccessColor(a string) string {
	if config.Config.IsProd {
		return a
	}
	return color.Green.Sprint(a)
}
