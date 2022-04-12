package logger

import "fmt"

//定义全局的日志化文件对象。
var log LogInterface

/*
封装日志库对外提供服务
*/
func InitLogger(name string, config map[string]string) (err error) {
	switch name {
	case "file":
		log, err = NewFileLogger(config)
	case "console":
		log, err = NewConsoleLogger(config)
	default:
		err = fmt.Errorf("unsupport logger name:%s", name)
	}
	return
}

//封装日志库对外提供的方法
func Debug(format string, args ...interface{}) {
	log.Debug(format, args...)
}
func Trace(format string, args ...interface{}) {
	log.Trace(format, args...)
}
func Info(format string, args ...interface{}) {
	log.Info(format, args...)
}
func Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}
func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}
func Fatal(format string, args ...interface{}) {
	log.Fatal(format, args...)
}
