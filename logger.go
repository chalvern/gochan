package gochan

import (
	"log"
)

type gochanLogger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}

var logger gochanLogger = &defaultLogger{}

// SetLogger logger for gochan
func SetLogger(l gochanLogger) {
	logger = l
}

type defaultLogger struct {
}

func (dl *defaultLogger) Debug(v ...interface{}) {
	log.Println(v...)
}

func (dl *defaultLogger) Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (dl *defaultLogger) Info(v ...interface{}) {
	log.Println(v...)
}
func (dl *defaultLogger) Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func (dl *defaultLogger) Error(v ...interface{}) {
	log.Println(v...)
}
func (dl *defaultLogger) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
