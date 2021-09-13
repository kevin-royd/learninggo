package logger

import (
	"path"
	"runtime"
)

func GetLineInfo()(fileName string,funcName string,lineNo int){
	// pc 程序执行计数器整数类型
	// 获取的文件路径和程序路径都为绝对路径，我们只需要最后路径即可
	pc, file, line, ok := runtime.Caller(0)
	if ok{
		fileName = path.Base(file)
		funcName = path.Base(runtime.FuncForPC(pc).Name())
		lineNo = line
	}
	return

}