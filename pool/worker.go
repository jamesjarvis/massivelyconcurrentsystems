package pool

import (
	"sync"
	"time"
)

// batchConsumer consumes from the inner queue and invokes the BatchWorker.
type batchConsumer[REQ, RESP any] struct {
	*queue[REQ, RESP]
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

loop:
	for {
		if e, ok := c.queue.dequeue(); ok {
			received = append(received, e)

			if len(received) >= c.batchSize {
				// reached max batch size.
				doWork()
			} else {
				// otherwise, check for batchInterval timeout.
				select {
				case <-watchdog.C:
					// if batchInterval timeout reached, operate on current batch size.
					doWork()
				default:
				}
			}
		} else {
			// nothing currently in the queue to fetch.
			select {
			case <-c.close:
				doWork()
				if !watchdog.Stop() {
					<-watchdog.C
				}
				c.waitClose.Done()
				break loop
			case <-watchdog.C:
				// TODO(jamesjarvis): revisit this strategy. I think it means if we have nothing in the queue,
				// we force ourselves to wait for the whole duration of the batchInterval before picking up values again.
				doWork()
			case e := <-c.queue.ch:
				if e != nil {
					received = append(received, e)
					if len(received) >= c.batchSize {
						// reached max batch size.
						doWork()
					}
				}
			}
		}
	}
}
