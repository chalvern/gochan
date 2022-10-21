package gochan

// TaskFunc task
type TaskFunc func() error

type gochan struct {
	uuid      int
	tasksChan chan TaskFunc
	dieChan   chan struct{}
}

// newGochan return gochan with bufferNum tasks
func newGochan(bufferNum int) *gochan {
	gc := &gochan{
		uuid:      defualtUUID(),
		tasksChan: make(chan TaskFunc, bufferNum),
		dieChan:   make(chan struct{}),
	}
	return gc
}

func (gc *gochan) setUUID(uuid int) {
	gc.uuid = uuid
}

// run make gochan running
func (gc *gochan) run() {
	go gc.start()
}

// start gochan's goroutine
func (gc *gochan) start() {
	defer func() {
		logger.Infof("gochan %d ending...", gc.uuid)
	}()
	logger.Debugf("gochan %d starting...", gc.uuid)

	for {
		select {
		case task := <-gc.tasksChan:
			err := task()
			if err != nil {
				logger.Errorf("task in gochan %d error: %s", gc.uuid, err.Error())
			}
		case <-gc.dieChan:
			return
		}
	}
}
