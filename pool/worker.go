package pool

import (
	"sync"
	"time"
)

// batchConsumer consumes from the inner queue and invokes the BatchWorker.
type batchConsumer struct {
	*queue
	close     chan struct{}
	waitClose *sync.WaitGroup

	worker        BatchWorker
	batchSize     int
	batchInterval time.Duration
}

func (c *batchConsumer) start() {
	received := make([]UnitOfWork, 0, c.batchSize)
	watchdog := time.NewTimer(c.batchInterval)

	doWork := func() {
		watchdog.Stop()

		if len(received) > 0 {
			_ = c.worker(received) // TODO(jamesjarvis): Extend with the ability to append the error to all items in the batch.
			// TODO(jamesjarvis): Extend to support passing a context into the BatchWorker.
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
			}
		}
	}
}
