package logger

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
	level int
}

// NewConsoleLogger 定义构造函数
func NewConsoleLogger(config map[string]string) (log LogInterface, err error) {
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found log_level")
	}
	level := getLogLevel(logLevel)
	log = &ConsoleLogger{
		level: level,
	}
	return
}

func (c *ConsoleLogger) Init() {

}

// SetLevel 设置日志级别
func (c *ConsoleLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		// 因为不输出到日志文件中所有只有一个日志级别
		level = LogLevelDebug
	}
	c.level = level
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	//做日志级别校验
	if c.level > LogLevelDebug {
		return
	}
	data := writeLog(LogLevelDebug, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s:%d) %s\n", data.TimeStr, data.LevelStr, data.Filename, data.Filename, data.LineNo, data.Message)

}
func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	if c.level > LogLevelTrace {
		return
	}
	data := writeLog(LogLevelTrace, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s:%d) %s\n", data.TimeStr, data.LevelStr, data.Filename, data.Filename, data.LineNo, data.Message)
}
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > LogLevelInfo {
		return
	}
	data := writeLog(LogLevelInfo, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s:%d) %s\n", data.TimeStr, data.LevelStr, data.Filename, data.Filename, data.LineNo, data.Message)
}
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > LogLevelWarn {
		return
	}
	data := writeLog(LogLevelWarn, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s:%d) %s\n", data.TimeStr, data.LevelStr, data.Filename, data.Filename, data.LineNo, data.Message)
}
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > LogLevelError {
		return
	}
	data := writeLog(LogLevelError, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s:%d) %s\n", data.TimeStr, data.LevelStr, data.Filename, data.Filename, data.LineNo, data.Message)
}
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > LogLevelFatal {
		return
	}
	data := writeLog(LogLevelFatal, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s:%d) %s\n", data.TimeStr, data.LevelStr, data.Filename, data.Filename, data.LineNo, data.Message)
}

func (c ConsoleLogger) Close() {

}
