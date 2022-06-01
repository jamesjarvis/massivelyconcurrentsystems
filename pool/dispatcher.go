package pool

import (
	"context"
	"sync"
)

// BatchDispatcher controls access to the batched pool implementation.
type BatchDispatcher[REQ, RESP any] struct {
	// inner buffer
	*queue[REQ, RESP]

	*batchConsumer[REQ, RESP]
	numConsumers int

	close     chan struct{} // notify workers to close.
	waitClose *sync.WaitGroup
}

// NewBatchDispatcher returns a BatchDispatcher parameterised from Config.
func NewBatchDispatcher[REQ, RESP any](worker BatchWorker[REQ, RESP], config Config) BatchDispatcher[REQ, RESP] {
	var waitClose sync.WaitGroup
	closeChan := make(chan struct{})

	q := newQueue[REQ, RESP](config.bufferSize, &waitClose)
	d := BatchDispatcher[REQ, RESP]{
		queue:        q,
		numConsumers: config.numConsumers,

		batchConsumer: &batchConsumer[REQ, RESP]{
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

// Start initialises the BatchDispatcher.
func (d BatchDispatcher[REQ, RESP]) Start() {
	d.waitClose.Add(d.numConsumers)
	for i := 0; i < d.numConsumers; i++ {
		go d.batchConsumer.start() // start consumer
	}
}

// Put submits the UnitOfWork to the worker pool.
func (d BatchDispatcher[REQ, RESP]) Put(ctx context.Context, u UnitOfWork[REQ, RESP]) error {
	return d.queue.enqueue(ctx, u)
}

// Close gracefully shuts down the BatchDispatcher.
func (d BatchDispatcher[REQ, RESP]) Close() error {
	d.queue.close()
	close(d.close)
	d.waitClose.Wait()
	return nil
}
