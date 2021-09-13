package log_file

import (
	"fmt"
	"learninggo/log/log_level/log_interface"
)

//封装log日志对象
type FileLog struct {
}

func NewFileLog(file string) log_interface.LogInterFace {
	return &FileLog{}
}

func (f *FileLog) LogDebug(msg string) {
	fmt.Println(msg)
}
func (f *FileLog) LogWarn(msg string) {
	fmt.Println(msg)
}
