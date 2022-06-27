package pool

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// BatchDispatcher controls access to the batched pool implementation.
type BatchDispatcher[REQ, RESP any] struct {
	// inner buffer
	*queue[REQ, RESP]
	bufferSize int

	*batchConsumer[REQ, RESP]
	numConsumers int

	close         chan struct{} // notify workers to close.
	closeWatchdog chan struct{} // notify watchdog to close.
	waitClose     *sync.WaitGroup
}

// NewBatchDispatcher returns a BatchDispatcher parameterised from Config.
func NewBatchDispatcher[REQ, RESP any](worker BatchWorker[REQ, RESP], config Config) *BatchDispatcher[REQ, RESP] {
	var waitClose sync.WaitGroup
	closeChan := make(chan struct{})
	closeWatchdogChan := make(chan struct{})

	q := newQueue[REQ, RESP](config.bufferSize, &waitClose)
	d := &BatchDispatcher[REQ, RESP]{
		queue:        q,
		bufferSize:   config.bufferSize,
		numConsumers: config.numConsumers,

		batchConsumer: &batchConsumer[REQ, RESP]{
			queue:         q,
			close:         closeChan,
			waitClose:     &waitClose,
			worker:        worker,
			batchSize:     config.batchSize,
			batchInterval: config.batchInterval,
		},

		close:         closeChan,
		closeWatchdog: closeWatchdogChan,
		waitClose:     &waitClose,
	}
	return d
}

// watchDog will periodically look at the goroutine pool, and scale up/down based on the pressure on the queue.
// Initial implementation will double the number of workers, if queue is above 80% of bufferSize.
// If queue is between 20-80% of bufferSize, do nothing.
// If queue is between 0-20% of bufferSize, remove 10% of workers.
// after each change, wait waitDuration before checking again.
func (d *BatchDispatcher[REQ, RESP]) watchDog() {
	const tik = 500 * time.Millisecond
	tkr := time.NewTicker(tik)
	var sizePerc float64
	for {
		select {
		case <-d.closeWatchdog:
			tkr.Stop()
			return
		case <-tkr.C:
			fmt.Printf("Size of worker pool: %d, size of queue: %d/%d\n", d.numConsumers, d.queue.size(), d.bufferSize)
			sizePerc = float64(d.queue.size() / d.bufferSize)
			switch {
			case sizePerc > 0.8:
				// Add workers to pool.
				newWorkers := int(d.numConsumers)
				d.start(newWorkers)
				d.numConsumers += newWorkers
			case sizePerc < 0.2:
				// Remove workers from pool.
				for i := 0; i < int(float64(d.numConsumers)*0.1); i++ {
					d.close <- struct{}{}
					d.numConsumers -= 1
				}
			}
			tkr.Reset(tik)
		}
	}
}

func (d *BatchDispatcher[REQ, RESP]) start(numWorkers int) {
	d.waitClose.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go d.batchConsumer.start() // start consumer
	}
}

// Start initialises the BatchDispatcher.
func (d *BatchDispatcher[REQ, RESP]) Start() {
	d.start(d.numConsumers)
	go d.watchDog()
}

// Put submits the UnitOfWork to the worker pool.
func (d *BatchDispatcher[REQ, RESP]) Put(ctx context.Context, u UnitOfWork[REQ, RESP]) error {
	return d.queue.enqueue(ctx, u)
}

// Close gracefully shuts down the BatchDispatcher.
func (d *BatchDispatcher[REQ, RESP]) Close() error {
	d.queue.close()
	close(d.closeWatchdog)
	close(d.close)
	d.waitClose.Wait()
	return nil
}
