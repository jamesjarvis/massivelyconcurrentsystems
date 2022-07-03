package pool

import (
	"context"
	"sync"
)

// WorkDispatcher controls access to the pool implementation.
type WorkDispatcher[E any] struct {
	// inner buffer
	*queue[E]
	bufferSize int

	worker       Consumer
	numConsumers int

	close     chan struct{} // notify workers to close.
	waitClose *sync.WaitGroup
}

// NewBatchDispatcher returns a WorkDispatcher parameterised from Config.
// It is specialised such that it will receive a batch of UnitOfWork's in each worker.
func NewBatchDispatcher[REQ, RESP any](worker BatchWorker[REQ, RESP], config Config) *WorkDispatcher[UnitOfWork[REQ, RESP]] {
	var waitClose sync.WaitGroup
	closeChan := make(chan struct{})

	q := newQueue[UnitOfWork[REQ, RESP]](config.bufferSize, &waitClose)
	d := &WorkDispatcher[UnitOfWork[REQ, RESP]]{
		queue:        q,
		bufferSize:   config.bufferSize,
		numConsumers: config.numConsumers,

		worker: &batchConsumer[REQ, RESP]{
			queue:         q,
			close:         closeChan,
			waitClose:     &waitClose,
			worker:        worker,
			batchSize:     config.batchSize,
			batchInterval: config.batchInterval,
		},

		close:     closeChan,
		waitClose: &waitClose,
	}
	return d
}

// NewSingleDispatcher returns a WorkDispatcher parameterised from Config.
// It is specialised such that each worker will receive one element each to work on.
func NewSingleDispatcher[E any](worker Worker[E], config Config) *WorkDispatcher[E] {
	var waitClose sync.WaitGroup
	closeChan := make(chan struct{})

	q := newQueue[E](config.bufferSize, &waitClose)
	d := &WorkDispatcher[E]{
		queue:        q,
		bufferSize:   config.bufferSize,
		numConsumers: config.numConsumers,

		worker: &singleConsumer[E]{
			queue:     q,
			close:     closeChan,
			waitClose: &waitClose,
			worker:    worker,
		},

		close:     closeChan,
		waitClose: &waitClose,
	}
	return d
}

func (d *WorkDispatcher[E]) start(numWorkers int) {
	d.waitClose.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go d.worker.Start() // start consumer
	}
}

// Start initialises the BatchDispatcher.
func (d *WorkDispatcher[E]) Start() {
	d.start(d.numConsumers)
}

// Put submits the UnitOfWork to the worker pool.
func (d *WorkDispatcher[E]) Put(ctx context.Context, e E) error {
	return d.queue.enqueue(ctx, e)
}

// Close gracefully shuts down the BatchDispatcher.
func (d *WorkDispatcher[E]) Close() error {
	d.queue.close()
	close(d.close)
	d.waitClose.Wait()
	return nil
}
