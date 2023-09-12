package EventBus

import "log"

// 独立的日志子模块

var logger iLogger

type iLogger interface {
	Error(content ...interface{})
	Warn(content ...interface{})
	Info(content ...interface{})
	Debug(content ...interface{})
	Log(content ...interface{})
}

func RegLog(log iLogger) {
	logger = log
}

type LoggerModk struct {
}

func (LoggerModk) Error(content ...interface{}) {
	log.Fatal(content...)
}

func (LoggerModk) Warn(content ...interface{}) {
	log.Print(content...)
}

func (LoggerModk) Info(content ...interface{}) {
	log.Print(content...)
}

func (LoggerModk) Debug(content ...interface{}) {
	log.Print(content...)
}

func (LoggerModk) Log(content ...interface{}) {
	log.Print(content...)
}
