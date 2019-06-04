package gochan_test

import (
	"errors"
	"strconv"
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

	for i := range make([]int, 6) {
		objID := i
		task1 := func() error {
			return errors.New("task " + strconv.Itoa(objID))
		}
		err := manager.Dispatch(objID, task1)
		assert.Nil(t, err)
	}
	for _ = range make([]int, 20) {
		objID := -1
		taskM3 := func() error {
			return errors.New("task " + strconv.Itoa(objID))
		}
		err := manager.Dispatch(objID, taskM3)
		assert.Nil(t, err)
	}

	time.Sleep(time.Second)

	manager.Close()
	err := manager.Dispatch(1, func() error { return nil })
	assert.NotNil(t, err)

}
