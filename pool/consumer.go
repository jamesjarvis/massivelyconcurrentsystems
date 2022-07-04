package pool

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/jamesjarvis/massivelyconcurrentsystems/queue"
)

// Consumer is an internal consumer from a queue, calling go Start() spawns a new threadsafe consumer.
type Consumer interface {
	Start()
}

// batchWorkConsumer consumes from the inner queue and invokes the BatchWorker.
type batchWorkConsumer[REQ, RESP any] struct {
	queue     queue.Queue[UnitOfWork[REQ, RESP]]
	close     chan struct{}
	waitClose *sync.WaitGroup

	worker        BatchWorker[UnitOfWork[REQ, RESP]]
	batchSize     int
	batchInterval time.Duration
}

func (c *batchWorkConsumer[REQ, RESP]) Start() {
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
		if e, ok := c.queue.Dequeue(); ok {
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

// batchConsumer consumes from the inner queue and invokes the BatchWorker.
type batchConsumer[E any] struct {
	queue     queue.Queue[E]
	close     chan struct{}
	waitClose *sync.WaitGroup

	worker        BatchWorker[E]
	batchSize     int
	batchInterval time.Duration
}

func (c *batchConsumer[E]) Start() {
	received := make([]E, 0, c.batchSize)
	watchdog := time.NewTimer(c.batchInterval)

	doWork := func() {
		watchdog.Stop()

		if len(received) > 0 {
			// TODO(jamesjarvis): Extend to support passing a context into the BatchWorker.
			err := c.worker(received)
			if err != nil {
				// TODO(jamesjarvis): Do something with this error I guess?
				log.Println("Error encountered!", err)
			}
			received = received[:0]
		}
		watchdog.Reset(c.batchInterval)
	}

	for {
		if e, ok := c.queue.Dequeue(); ok {
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

// singleConsumer consumes from the inner queue and invokes the Worker.
type singleConsumer[E any] struct {
	queue     queue.Queue[E]
	close     chan struct{}
	waitClose *sync.WaitGroup

	worker Worker[E]
}

func (c *singleConsumer[E]) Start() {
	for {
		select {
		case e := <-c.queue.DequeueBlocking():
			c.worker(e)
		case <-c.close: // kill the worker.
			c.waitClose.Done()
			return
		}
	}
}
