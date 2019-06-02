package gochan_test

import (
	"errors"
	"testing"
	"time"

	"github.com/chalvern/gochan"
)

type Manager struct {
	dispatcher *gochan.Dispatcher
}

func (m *Manager) Dispatch(objID int, task gochan.TaskFunc) {
	m.dispatcher.Dispatch(objID, task)
}

func TestDispatcher(t *testing.T) {
	gochanNum := 3
	bufferNum := 10
	manager := Manager{
		dispatcher: gochan.NewDispatcher(gochanNum, bufferNum),
	}

	objID := 1
	task1 := func() error {
		return errors.New("task 1")
	}
	manager.Dispatch(objID, task1)

	objID = 2
	task2 := func() error {
		return errors.New("task 2")
	}
	manager.Dispatch(objID, task2)

	objID = 3
	task3 := func() error {
		return errors.New("task 3")
	}
	manager.Dispatch(objID, task3)

	objID = 4
	task4 := func() error {
		return errors.New("task 4")
	}
	manager.Dispatch(objID, task4)

	time.Sleep(time.Second)
}
