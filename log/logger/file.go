package logger

import (
	"fmt"
	"os"
	"time"
)

/*
定义文件类型实现日志库接口
*/

// 日志的格式 日期 日志级别 日志打印源文件和行数， 错误信息
type FileLogger struct {
	Level    int
	LogPath  string
	LogName  string
	file     *os.File //错误文件日志
	warnFile *os.File //警告文件日志
}

// NewFileLogger 生成构造函数 返回接口类型
func NewFileLogger(level int, logPath string, logName string) LogInterface {
	logger := &FileLogger{
		Level:   level,
		LogPath: logPath,
		LogName: logName,
	}
	//	文件初始化
	logger.init()
	return logger
}

func (f *FileLogger) init() {
	//	 构造错误文件，文件名
	filename := fmt.Sprintf("%s%s.log", f.LogPath, f.LogName)
	//打开文件，第一个os为如果文件不存在则创建，第二个将数据写入附加文件中，第三个为只以写的方法打开，第三个参数为文件权限
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0775)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err:%v", filename, err))
	}
	// 如果文件打开没问题就将文件进行赋值
	f.file = file

	filename = fmt.Sprintf("%s%s.log", f.LogPath, f.LogName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0775)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err:%v", filename, err))
	}
	f.warnFile = file

}

/*
实现接口方法
*/

func (f *FileLogger) SetLevel(level int) {
	//判断日志级别是否合法
	if level < 0 || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.Level = level
	fmt.Println("implement me")
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.Level > LogLevelDebug {
		return
	}
	//	获取时间 如果获取时间戳则为time.now().unix()
	now := time.Now()                          //返回当前本地时间，time
	s := now.Format("2006-01-02 15:04:05.999") //注意括号中的时间是不能进行更改的但格式可以更改
	//获取日志级别
	levelStr := getLevelText(LogLevelDebug)
	fileName, funcName, lineNo := GetLineInfo()
	// 用户传入的日志格式，并将可变参数进行格式化
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(f.file, "%s %s %s:%d %s %s\n", s, levelStr, fileName, lineNo, funcName, msg)
	//fmt.Fprintf(f.file, s)
	////根据格式化吸入,第一个为输出文件,第二个为数据格式，第三个为参数
	//fmt.Fprintf(f.file, format, args...)
	//fmt.Fprintln(f.file) //起换行的作用

}
func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.Level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.file, format, args...)
	fmt.Fprintln(f.file) //起换行的作用
}
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.Level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.file, format, args...)
	fmt.Fprintln(f.file) //起换行的作用
}
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.Level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.warnFile, format, args...)
	fmt.Fprintln(f.warnFile) //起换行的作用
}
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.Level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.warnFile, format, args...)
	fmt.Fprintln(f.warnFile) //起换行的作用
}
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.Level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.warnFile, format, args...)
	fmt.Fprintln(f.warnFile) //起换行的作用
}

func (f *FileLogger) Close() {
	f.warnFile.Close()
	f.file.Close()
}
