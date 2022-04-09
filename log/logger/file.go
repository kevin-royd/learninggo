package logger

import (
	"fmt"
	"os"
	"time"
)

/*
定义文件类型实现日志库接口
*/

// FileLogger 日志的格式 日期 日志级别 日志打印源文件和行数， 错误信息
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

func (f FileLogger) writeLog(file *os.File,level int,format string,args ...interface{})  {
	if f.Level > level{
		return
	}
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

func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.writeLog(f.file,LogLevelDebug,format,args...)
}
func (f *FileLogger) Trace(format string, args ...interface{}) {
	f.writeLog(f.file,LogLevelTrace,format,args...)
}
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.writeLog(f.file,LogLevelInfo,format,args...)
}
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.writeLog(f.file,LogLevelWarn,format,args...)
}
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.writeLog(f.file,LogLevelError,format,args...)
}
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.writeLog(f.file,LogLevelFatal,format,args...)
}

func (f *FileLogger) Close() {
	f.warnFile.Close()
	f.file.Close()
}
