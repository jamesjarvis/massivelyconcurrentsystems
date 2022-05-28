package pool

import (
	"context"
	"sync"
)

// BatchDispatcher controls access to the batched pool implementation.
type BatchDispatcher struct {
	// inner buffer
	*queue

	*batchConsumer
	numConsumers int

	close     chan struct{} // notify workers to close.
	waitClose *sync.WaitGroup
}

// NewBatchDispatcher returns a BatchDispatcher parameterised from Config.
func NewBatchDispatcher(config Config) BatchDispatcher {
	var waitClose sync.WaitGroup
	closeChan := make(chan struct{})

	q := newQueue(config.bufferSize, &waitClose)
	d := BatchDispatcher{
		queue:        q,
		numConsumers: config.numConsumers,

		batchConsumer: &batchConsumer{
			queue:         q,
			close:         closeChan,
			waitClose:     &waitClose,
			worker:        config.worker,
			batchSize:     config.batchSize,
			batchInterval: config.batchInterval,
		},

		close:     closeChan,
		waitClose: &waitClose,
	}
	return d
}

// Start initialises the BatchDispatcher.
func (d BatchDispatcher) Start() {
	d.waitClose.Add(d.numConsumers)
	for i := 0; i < d.numConsumers; i++ {
		go d.batchConsumer.start() // start consumer
	}
}

// Put submits the UnitOfWork to the worker pool.
func (d BatchDispatcher) Put(ctx context.Context, u UnitOfWork) error {
	return d.enqueue(ctx, u)
}

// Close gracefully shuts down the BatchDispatcher.
func (d BatchDispatcher) Close() error {
	d.queue.close()
	close(d.close)
	d.waitClose.Wait()
	return nil
}
