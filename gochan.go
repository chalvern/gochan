package gochan

type pool struct {
}

type gochan struct {
	uuid      int
	tasksChan chan func()
	dieChan   chan struct{}
}

// newGochan return gochan with bufferNum tasks
func newGochan(bufferNum int) *gochan {
	return &gochan{
		uuid:      1,
		tasksChan: make(chan func(), bufferNum),
		dieChan:   make(chan struct{}),
	}
}

// start gochan's goroutine
func (gc *gochan) start() {
	defer func() {
		logger.Infof("gochan %d ending...", gc.uuid)
	}()
	logger.Infof("gochan %d starting...", gc.uuid)

	for {
		select {
		case task := <-gc.tasksChan:
			task()
		case <-gc.dieChan:
			return
		}
	}
}
