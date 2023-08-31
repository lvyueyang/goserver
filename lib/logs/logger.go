package logs

import (
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"selfserver/config"
	"selfserver/lib/console"
)

var fileLogger zerolog.Logger
var stdLogger zerolog.Logger

func InitLogger() {
	fileLogger = createFileLog()
	stdLogger = createStdLog()
}

func getLogWriter() *lumberjack.Logger {
	fileName := path.Join(config.GetLoggerOutPutPath(), "log.log")
	console.Success("日志输出路径：", fileName)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    1,     // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 100,   // 保留旧文件的最大个数
		MaxAge:     30,    // 保留旧文件的最大天数
		Compress:   false, // 是否压缩/归档旧文件
	}
	return lumberJackLogger
}

// 文件日志
func createFileLog() zerolog.Logger {
	logger := zerolog.New(getLogWriter())
	return logger
}

// 控制台日志
func createStdLog() zerolog.Logger {
	logger := zerolog.New(os.Stderr)
	return logger
}

// Debug 打印调试信息
// 仅在开发环境的控制台输出调试信息
func Debug() *zerolog.Event {
	return stdLogger.Debug()
}

func Info() *zerolog.Event {
	return fileLogger.Info()
}

// Warn 打印警告信息在文件中输出
func Warn() *zerolog.Event {
	return fileLogger.Warn()
}

// Err 打印警告信息在文件中输出
func Err() *zerolog.Event {
	return fileLogger.Error()
}
