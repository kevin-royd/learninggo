package logger

import "testing"

/*
单元测试，文件名都是以_test结尾，方法名都是以Test开头
*/
func TestFIleLogger(t *testing.T) {
	//	创建文件实例
	logger := NewFileLogger(LogLevelDebug, `/Users/royd/code/`, "logger")
	logger.Debug("the user is first from china as us")
	logger.Warn("this is warn test")
	//	io流操作都需要close
	logger.Close()
}

/*
控制台输出测试函数
当一个测试用例有多个test时
*/
func TestConsoleLogger(t *testing.T) {
	//	创建文件实例
	logger := NewConsoleLogger(LogLevelDebug)
	logger.Debug("the user is first from china as us")
	logger.Warn("this is warn test")
	//	io流操作都需要close
	logger.Close()
}

