package gochan

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("tesing begin...")
	SetLogger(&defaultLogger{})
	resultCode := m.Run()
	fmt.Println("testing end.")
	os.Exit(resultCode)
}

func TestGoChan(t *testing.T) {
	gc := newGochan(1)
	assert.NotNil(t, gc)

	t.Run("startAndEnd", func(t *testing.T) {
		go gc.start()
		a := 0
		gc.tasksChan <- func() {
			a = 100
		}

		time.Sleep(time.Second)
		assert.Equal(t, 100, a)

		gc.dieChan <- struct{}{}
	})
}
