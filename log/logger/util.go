package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

func GetLineInfo()(fileName string,funcName string,lineNo int){
	// pc 程序执行计数器整数类型
	// 获取的文件路径和程序路径都为绝对路径，我们只需要最后路径即可
	// Caller为栈的深度 实例化对象调用debug级别日志为第0层，初始化函数第1层，调用函数格式化日志第二层，实现接口数据写人第3层
	pc, file, line, ok := runtime.Caller(3)
	if ok{
		fileName = path.Base(file)
		funcName = path.Base(runtime.FuncForPC(pc).Name())
		lineNo = line
	}
	return
}

func writeLog(file *os.File,level int,format string,args ...interface{})  {
	//	获取时间 如果获取时间戳则为time.now().unix()
	now := time.Now()                                //返回当前本地时间，time
	nowData := now.Format("2006-01-02 15:04:05.999") //注意括号中的时间是不能进行更改的但格式可以更改
	//获取日志级别
	levelStr := getLevelText(level)
	// fileName调用程序的文件名和函数名，写入的行数
	fileName, funcName, lineNo := GetLineInfo()
	// 用户传入的日志格式，并将可变参数进行格式化
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", nowData, levelStr, fileName, funcName, lineNo, msg)

}
