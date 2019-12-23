package utils

import (
	"fmt"
	"os"
)

// 定义错误级别
const (
	// LevelError 异常
	LevelPlanc = iota
	// LevelError 错误
	LevelError
	// LevelWarning 警告
	LevelWarning
	// LevelInformational 提示
	LevelInformational
	// LevelDebug 排查
	LevelDebug
)

var logger *Logger

var errorLevel = LevelDebug

type Logger struct {
	level int
	Msg   []string
}

// 将错误级别转换成文字
func (l *Logger) levelToM() (msg string) {
	msg = "未定义的错误级别"
	switch l.level {
	case LevelPlanc:
		msg = "异常"
	case LevelError:
		msg = "错误"
	case LevelWarning:
		msg = "警告"
	case LevelInformational:
		msg = "提示"
	case LevelDebug:
		msg = "排查"
	}
	return
}

// 普通打印
func (l *Logger) PrintLn(msg string) {
	msg = fmt.Sprintf("[%s][message]%s", l.levelToM(), msg)
	l.Msg = append(l.Msg, msg)
}

// 异常打印
func (l *Logger) Planc(format string, v ...interface{}) {
	if errorLevel < LevelPlanc {
		return
	}
	l.level = LevelPlanc
	msg := ""
	if v == nil {
		msg = format
	} else {
		msg = fmt.Sprintf(format, v...)
	}
	l.PrintLn(msg)
	os.Exit(0)
}

// 错误打印
func (l *Logger) Error(format string, v ...interface{}) {
	if errorLevel < LevelError {
		return
	}
	l.level = LevelError
	msg := ""
	if v == nil {
		msg = format
	} else {
		msg = fmt.Sprintf(format, v...)
	}
	l.PrintLn(msg)
}

// 警告打印
func (l *Logger) Warning(format string, v ...interface{}) {
	if errorLevel < LevelWarning {
		return
	}
	l.level = LevelWarning
	msg := ""
	if v == nil {
		msg = format
	} else {
		msg = fmt.Sprintf(format, v...)
	}
	l.PrintLn(msg)
}

// 提示打印
func (l *Logger) Info(format string, v ...interface{}) {
	if errorLevel < LevelInformational {
		return
	}
	l.level = LevelInformational
	msg := ""
	if v == nil {
		msg = format
	} else {
		msg = fmt.Sprintf(format, v...)
	}
	l.PrintLn(msg)
}

// 排错打印
func (l *Logger) Debug(format string, v ...interface{}) {
	if errorLevel < LevelDebug {
		return
	}
	l.level = LevelDebug
	msg := ""
	if v == nil {
		msg = format
	} else {
		msg = fmt.Sprintf(format, v...)
	}
	l.PrintLn(msg)
}

func Log() *Logger {

	if logger == nil {
		logger = &Logger{level: -1}
	}

	return logger
}

func SetErrorLevel(level string) {

	var ErrorLevelIto = LevelDebug

	switch level {
	case "debug":
		ErrorLevelIto = LevelDebug
	case "informational":
		ErrorLevelIto = LevelInformational
	case "warning":
		ErrorLevelIto = LevelWarning
	case "error":
		ErrorLevelIto = LevelError
	case "planc":
		ErrorLevelIto = LevelPlanc

	}
	errorLevel = ErrorLevelIto
}
