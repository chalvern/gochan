package gochan

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("tesing begin...")
	resultCode := m.Run()
	fmt.Println("testing end.")
	os.Exit(resultCode)
}

func TestGoChan(t *testing.T) {
	gc := newGochan(1)
	gc.run()
	assert.NotNil(t, gc)

	t.Run("startAndEnd", func(t *testing.T) {
		a := 0
		myCh := make(chan struct{})
		gc.tasksChan <- func() error {
			a = 100
			myCh <- struct{}{}
			return nil
		}

		<-myCh
		assert.Equal(t, 100, a)

		gc.dieChan <- struct{}{}
	})
}
