# gochan

[中文说明文档](./README-zh.md)

## background

In general, a buffered channel variable is defined with a dedicated 
goroutine to execute the specific events in channel.

If there are so many events that a single gorotine is not enough, 
may more goroutines are created to execute events. However, it is
not okey if some events need to be executed sequentially.

Two events, for example, one order payment and goods delivery of 
that order, need to be executed sequentially. At the same time, 
events of different orders can be executed concurrently. So that 
we can push events of the same order to the same buffer channel, 
and a specific goroutine to execute the events in the queue. 

Concurrency is achieved by expanding multiple similar combinations 
(a buffer channel and its specific goroutine) in parallel.

### Usual design

a buffered channel variable is defined with a dedicated 
goroutine to execute the specific events in channel

```bash

event ->
        |
event -> buffer-channel -> goroutine
        |
event ->

```

### State-independent concurrent design

There is no state dependency between events, so you can 
simply extend goroutine to speed up event execution.

```bash

event ->                 ->goroutine
        |               |
event -> buffer-channel -> goroutine
        |               |
event ->                 ->goroutine

```

### State-dependent concurrent design (as gochan does)

Introduce a layer of dispatcher to distribute events to 
the corresponding buffer-channel according to a feature 
such as uuid.

```bash

event ->              -> buffer-channel -> goroutine
        |            |
event --> dispatcher -> buffer-channel -> goroutine
        |            |
event ->              -> buffer-channel -> goroutine

```

## example

you can also find example in [examples](./examples).

```go

package main

import (
	"errors"
	"math/rand"
	"time"

	"github.com/chalvern/gochan"
)

type Manager struct {
	gochanNum  int
	bufferNum  int
	dispatcher *gochan.Dispatcher
}

func (m *Manager) Dispatch(objID int, task gochan.TaskFunc) error {
	if objID < 0 {
		objID = rand.Intn(m.gochanNum)
	}
	return m.dispatcher.Dispatch(objID, task)
}

func (m *Manager) Close() {
	m.dispatcher.Close()
}

func main() {
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

```