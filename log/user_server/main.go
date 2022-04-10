package main

import "learninggo/log/logger"

func initLogger(name, logPath, logName string, level string) (err error) {
	//对map进行赋值
	m := make(map[string]string, 8)
	m["log_path"] = logPath
	m["log_name"] = logName
	m["log_level"] = level
	err = logger.InitLogger(name, m)
	if err != nil {
		panic(err)
	}
	logger.Debug("init logger success")
	return
}

func run() {
	for i := 0; i < 5; i++ {
		logger.Debug("user server is running")
	}
}
func main() {
	initLogger("console", "/Users/royd/code/", "logger", "DEBUG")
	run()
}
