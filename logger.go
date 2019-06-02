package gochan

import (
	"log"
)

type gochanLogger interface {
	Info(v ...interface{})
	Infof(format string, v ...interface{})
}

var logger gochanLogger

// SetLogger logger for gochan
func SetLogger(l gochanLogger) {
	logger = l
}

type defaultLogger struct {
}

func (dl *defaultLogger) Info(v ...interface{}) {
	log.Println(v...)
}
func (dl *defaultLogger) Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}
