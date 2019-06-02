package gochan_test

import (
	"errors"
	"testing"
	"time"

	"github.com/chalvern/gochan"
	"github.com/stretchr/testify/assert"
)

type Manager struct {
	dispatcher *gochan.Dispatcher
}

func (m *Manager) Dispatch(objID int, task gochan.TaskFunc) error {
	return m.dispatcher.Dispatch(objID, task)
}

func (m *Manager) Close() {
	m.dispatcher.Close()
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
	err := manager.Dispatch(objID, task1)
	assert.Nil(t, err)

	objID = 2
	task2 := func() error {
		return errors.New("task 2")
	}
	err = manager.Dispatch(objID, task2)
	assert.Nil(t, err)

	objID = 3
	task3 := func() error {
		return errors.New("task 3")
	}
	err = manager.Dispatch(objID, task3)
	assert.Nil(t, err)

	objID = 4
	task4 := func() error {
		return errors.New("task 4")
	}
	err = manager.Dispatch(objID, task4)
	assert.Nil(t, err)

	time.Sleep(time.Second)

	manager.Close()
	err = manager.Dispatch(objID, task4)
	assert.NotNil(t, err)

}
