package main

import (
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
	myCh := make(chan struct{})
	myNumber := 2016
	task1 := func() error {
		myNumber = 2019
		myCh <- struct{}{}
		return nil
	}
	manager.Dispatch(objID, task1)
	<-myCh
	if myNumber != 2019 {
		panic("myNumber should be 2019")
	}
}
