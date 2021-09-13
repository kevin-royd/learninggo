package main

import (
	"learninggo/log/log_level/log_file"
)

func main() {
	log := log_file.NewFileLog("c:/a")
	log.LogDebug("aaaa")
}
