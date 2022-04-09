package logger

import (
	"path"
	"runtime"
)

func GetLineInfo()(fileName string,funcName string,lineNo int){
	// pc 程序执行计数器整数类型
	// 获取的文件路径和程序路径都为绝对路径，我们只需要最后路径即可
	// Caller为栈的深度 实例化对象调用debug级别日志为第0层，程序函数为第1层
	pc, file, line, ok := runtime.Caller(1)
	if ok{
		fileName = path.Base(file)
		funcName = path.Base(runtime.FuncForPC(pc).Name())
		lineNo = line
	}
	return

}