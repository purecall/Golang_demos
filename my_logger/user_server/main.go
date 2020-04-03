package main

import (
	"time"

	"../logger"
)

func initLogger(name, logPath, logName, level string) (err error) {
	config := make(map[string]string, 8)
	config["log_path"] = logPath
	config["log_name"] = "user_server"
	config["log_level"] = level
	config["log_chan_size"] = "32768"
	config["log_split_type"] = "size"
	err = logger.InitLogger(name, config)
	if err != nil {
		return
	}
	logger.Debug("init logger success")
	return
}

func Run() {
	for {
		logger.Debug("user server is running")
		time.Sleep(time.Second)
	}
}

func main() {
	initLogger("file", "c:/logs/", "user_server", "debug")
	Run()
	return
}
