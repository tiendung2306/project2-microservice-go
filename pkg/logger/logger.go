package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Level định nghĩa mức độ của log
type Level int

const (
	// DEBUG level log
	DEBUG Level = iota
	// INFO level log
	INFO
	// WARNING level log
	WARNING
	// ERROR level log
	ERROR
	// FATAL level log
	FATAL
)

var levelNames = []string{
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
	"FATAL",
}

// Config chứa các tùy chọn cấu hình cho logger
type Config struct {
	Level      Level
	TimeFormat string
	Output     io.Writer
}

// Logger định nghĩa một interface logger
type Logger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
	Fatal(format string, v ...interface{})
}

type logger struct {
	level      Level
	timeFormat string
	log        *log.Logger
}

// NewLogger tạo một instance mới của Logger
func NewLogger(config Config) Logger {
	if config.Output == nil {
		config.Output = os.Stdout
	}
	if config.TimeFormat == "" {
		config.TimeFormat = "2006-01-02 15:04:05"
	}

	return &logger{
		level:      config.Level,
		timeFormat: config.TimeFormat,
		log:        log.New(config.Output, "", 0),
	}
}

func (l *logger) log0(level Level, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	now := time.Now().Format(l.timeFormat)
	msg := fmt.Sprintf(format, v...)

	// Lấy thông tin file và line number
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	// Chỉ lấy tên file, không phải đường dẫn đầy đủ
	file = filepath.Base(file)

	l.log.Printf("%s [%s] %s:%d: %s", now, levelNames[level], file, line, msg)

	// Nếu là lỗi Fatal, kết thúc chương trình
	if level == FATAL {
		os.Exit(1)
	}
}

func (l *logger) Debug(format string, v ...interface{}) {
	l.log0(DEBUG, format, v...)
}

func (l *logger) Info(format string, v ...interface{}) {
	l.log0(INFO, format, v...)
}

func (l *logger) Warning(format string, v ...interface{}) {
	l.log0(WARNING, format, v...)
}

func (l *logger) Error(format string, v ...interface{}) {
	l.log0(ERROR, format, v...)
}

func (l *logger) Fatal(format string, v ...interface{}) {
	l.log0(FATAL, format, v...)
}
