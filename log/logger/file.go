package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
定义文件类型实现日志库接口
*/

// FileLogger 日志的格式 日期 日志级别 日志打印源文件和行数， 错误信息
type FileLogger struct {
	Level         int
	LogPath       string
	LogName       string
	file          *os.File //错误文件日志
	warnFile      *os.File //警告文件日志
	LogDataChan   chan *LogData
	LogSplitType  int
	LogSplitSize  int64
	LastSplitHour int //上次切分的小时数
}

// NewFileLogger 生成构造函数 返回接口类型
func NewFileLogger(config map[string]string) (log LogInterface, err error) {
	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("not found log_path")
		return
	}
	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("not found log_name")
		return
	}
	logLevel, ok := config["level"]
	if !ok {
		err = fmt.Errorf("not found level")
		return
	}
	logChanSize, ok := config["chan_size"]
	if !ok {
		// 如果没有设置channel的长度设置默认5万
		logChanSize = "50000"
	}
	//将字符串转换为int
	ChanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		ChanSize = 50000
	}
	// 定义接受日志切割的变量
	var logSplitType = LogSplitTypeHour
	var logSplitSize int64
	logSplitStr, ok := config["log_split_type"]
	if !ok {
		logSplitStr = "hour"
	} else {
		if logSplitStr == "size" {
			// 获取按日志切分的大小 如果没有就设置默认切割大小
			logSplitSizeStr, ok := config["log_split_size"]
			if !ok {
				logSplitSizeStr = "104857600"
			}
			logSplitSize, err = strconv.ParseInt(logSplitSizeStr, 10, 64)
			if err != nil {
				logSplitSize = 104857600
			}
			logSplitType = LogSplitTypeSize
		} else {
			logSplitType = LogSplitTypeHour
		}
	}

	level := getLogLevel(logLevel)
	log = &FileLogger{
		Level:         level,
		LogPath:       logPath,
		LogName:       logName,
		LogDataChan:   make(chan *LogData, ChanSize),
		LogSplitSize:  logSplitSize,
		LogSplitType:  logSplitType,
		LastSplitHour: time.Now().Hour(),
	}
	log.Init()
	return
}

func (f *FileLogger) Init() {
	//	 构造错误文件，文件名
	filename := fmt.Sprintf("%s%s.log", f.LogPath, f.LogName)
	//打开文件，第一个os为如果文件不存在则创建，第二个将数据写入附加文件中，第三个为只以写的方法打开，第三个参数为文件权限
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0775)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err:%v", filename, err))
	}
	// 如果文件打开没问题就将文件进行赋值
	f.file = file

	filename = fmt.Sprintf("%s%s-err.log", f.LogPath, f.LogName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0775)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err:%v", filename, err))
	}
	f.warnFile = file

	//开启后台进程
	go f.writeLogBackGroup()

}

func (f FileLogger) checkSplitHour(warnFile bool) {
	// 按时间切分检查上次切分的时间
	//如果上次时间和当前时间不同 则切分一个新的并且把当前时间赋值给变量
	now := time.Now()
	if f.LastSplitHour == now.Hour() {
		return
	}
	//定义新日志文件名
	var backupFileName string
	var fileName string
	//判断日志文件的日志级别
	file := f.file
	if warnFile {
		backupFileName = fmt.Sprintf("%s%s_err.log_%04d%02d%02d", f.LogPath, f.LogName, now.Year(), now.Month(), f.LastSplitHour)
		fileName = fmt.Sprintf("%s%s.log_err", f.LogPath, f.LogName)
	} else {
		backupFileName = fmt.Sprintf("%s%s_%04d%02d%02d", f.LogPath, f.LogName, now.Year(), now.Month(), f.LastSplitHour)
		fileName = fmt.Sprintf("%s%s.log", f.LogPath, f.LogName)
	}
	// 文件关闭进行改名
	file.Close()
	os.Rename(fileName, backupFileName)
	//重新将日志文件打开
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0775)
	if err != nil {
		return
	}
	//重新对文件进行赋值
	if warnFile {
		f.warnFile = file
	} else {
		f.file = file
	}
}

func (f FileLogger) checkSplitSize(warnFile bool) {
	file := f.file
	now := time.Now()
	if warnFile {
		file = f.warnFile
	}
	//判断文件的大小
	info, err := file.Stat()
	if err != nil {
		return
	}
	logSize := info.Size()
	if logSize <= f.LogSplitSize {
		return
	}
	var backupFileName string
	var fileName string
	if warnFile {
		backupFileName = fmt.Sprintf("%s%s_err.log_%04d%02d%02d%02d", f.LogPath, f.LogName, now.Year(), now.Month(), now.Hour(), now.Minute())
		fileName = fmt.Sprintf("%s%s.log_err", f.LogPath, f.LogName)
	} else {
		backupFileName = fmt.Sprintf("%s%s_err.log_%04d%02d%02d%02d", f.LogPath, f.LogName, now.Year(), now.Month(), now.Hour(), now.Minute())
		fileName = fmt.Sprintf("%s%s.log_err", f.LogPath, f.LogName)
	}
	file.Close()
	os.Rename(fileName, backupFileName)
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0775)
	if err != nil {
		return
	}
	if warnFile {
		f.warnFile = file
	} else {
		f.file = file
	}
}

func (f *FileLogger) checkSplitFile(warnFile bool) {
	// 检查日志切分的类型
	if f.LogSplitType == LogSplitTypeHour {
		f.checkSplitHour(warnFile)
		return
	} else {
		f.checkSplitSize(warnFile)
	}
}

// 实现取出channel中的值写入日志文件
func (f *FileLogger) writeLogBackGroup() {
	//channel为空会阻塞，但因为是子线程所有没有影响
	for data := range f.LogDataChan {
		var file *os.File = f.file
		//判断写入日志级别文件
		if data.WarnAndFatal {
			file = f.warnFile
		}
		//日志写入文件之前检查日志文件切分
		f.checkSplitFile(data.WarnAndFatal)
		fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", data.TimeStr, data.LevelStr, data.Filename, data.Filename, data.LineNo, data.Message)

	}
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
	//同样对日志级别进行校验
	if f.Level > LogLevelDebug {
		return
	}
	data := writeLog(LogLevelDebug, format, args...)
	// 对channel的长度进行判断，为超过放入，超过抛去，否则阻塞后严重影响性能
	select {
	case f.LogDataChan <- data:
	default:

	}
}
func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.Level > LogLevelTrace {
		return
	}
	data := writeLog(LogLevelTrace, format, args...)
	select {
	case f.LogDataChan <- data:
	default:

	}
}
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.Level > LogLevelInfo {
		return
	}
	data := writeLog(LogLevelInfo, format, args...)
	select {
	case f.LogDataChan <- data:
	default:

	}
}
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.Level > LogLevelWarn {
		return
	}
	data := writeLog(LogLevelWarn, format, args...)
	select {
	case f.LogDataChan <- data:
	default:

	}
}
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.Level > LogLevelError {
		return
	}
	data := writeLog(LogLevelError, format, args...)
	select {
	case f.LogDataChan <- data:
	default:

	}
}
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.Level > LogLevelFatal {
		return
	}
	data := writeLog(LogLevelFatal, format, args...)
	select {
	case f.LogDataChan <- data:
	default:

	}
}

func (f *FileLogger) Close() {
	f.warnFile.Close()
	f.file.Close()
}
