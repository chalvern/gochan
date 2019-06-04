package main

import (
	"errors"
	"time"

	"github.com/chalvern/gochan"
	"go.uber.org/zap"
)

type Manager struct {
	gochanNum  int // number of goroutine-channel pair
	bufferNum  int // number of channel buffer in each goroutine-channel pair
	dispatcher *gochan.Dispatcher
}

func (m *Manager) Dispatch(objID int, task gochan.TaskFunc) error {
	return m.dispatcher.Dispatch(objID, task)
}

func (m *Manager) Close() {
	m.dispatcher.Close()
}

func main() {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()
	gochan.SetLogger(sugar)

	gochanNum := 3
	bufferNum := 10
	manager := Manager{
		gochanNum:  gochanNum,
		bufferNum:  bufferNum,
		dispatcher: gochan.NewDispatcher(gochanNum, bufferNum),
	}

	objID := 1
	task1 := func() error {
		return errors.New("task 1")
	}
	manager.Dispatch(objID, task1)
	time.Sleep(time.Second)
}
