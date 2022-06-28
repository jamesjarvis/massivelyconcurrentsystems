package pool

import (
	"runtime"
	"sync"
	"time"
)

// batchConsumer consumes from the inner queue and invokes the BatchWorker.
type batchConsumer[REQ, RESP any] struct {
	*queue[UnitOfWork[REQ, RESP]]
	close     chan struct{}
	waitClose *sync.WaitGroup

	worker        BatchWorker[REQ, RESP]
	batchSize     int
	batchInterval time.Duration
}

func (c *batchConsumer[REQ, RESP]) start() {
	received := make([]UnitOfWork[REQ, RESP], 0, c.batchSize)
	watchdog := time.NewTimer(c.batchInterval)

	doWork := func() {
		watchdog.Stop()

		if len(received) > 0 {
			// TODO(jamesjarvis): Extend to support passing a context into the BatchWorker.
			err := c.worker(received)
			if err != nil {
				for _, uow := range received {
					uow.SetError(err)
				}
			}
			received = received[:0]
		}
		watchdog.Reset(c.batchInterval)
	}

	for {
		if e, ok := c.queue.dequeue(); ok {
			received = append(received, e)

			if len(received) >= c.batchSize {
				// reached max batch size.
				doWork()
				continue
			}
		}

		// nothing currently in the queue to fetch.
		select {
		case <-c.close: // kill the worker.
			doWork()
			if !watchdog.Stop() {
				<-watchdog.C
			}
			c.waitClose.Done()
			return
		case <-watchdog.C:
			doWork()
		default: // this forces a busy wait, perhaps this is too heavy?
			runtime.Gosched()
		}
	}
}
