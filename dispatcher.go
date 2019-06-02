package gochan

// Dispatcher pool of gochan
type Dispatcher struct {
	gcNum int
	gcs   []*gochan
}

// NewDispatcher new pool
// @gochanNum the number of gochan
// @bufferNum the buffer number of chan in each gochan
func NewDispatcher(gochanNum, bufferNum int) *Dispatcher {
	logger.Infof("%d gochans and %d bufferNum chan buffer", gochanNum, bufferNum)
	p := &Dispatcher{
		gcNum: gochanNum,
		gcs:   make([]*gochan, gochanNum),
	}
	for index := range p.gcs {
		gc := newGochan(bufferNum)
		gc.setUUID(index)
		p.gcs[index] = gc
	}
	return p
}

// Dispatch dispatch task referenced by objID
func (d *Dispatcher) Dispatch(objID int, task TaskFunc) {
	index := objID % d.gcNum
	d.gcs[index].tasksChan <- task
}
