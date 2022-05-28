// Package pool is designed to be a simple worker pool orchestrator,
// primarily focused on workloads that require batched execution of work
// due to I/O limitations.
// e.g.: batching access to a underlying database connection.
// Heavily influenced by github.com/jiacai2050/prosumer
package pool

import "context"

// UnitOfWork is the object designed to be passed into, and then operated on, by the pool.
type UnitOfWork interface {
	GetRequest() any
	GetResponse() any
	SetResponse(any)
	GetError() error
	SetError(error)
	Done()
}

// BatchWorker is a function that will be used to operate on a batch of requests from the pool.
type BatchWorker func([]UnitOfWork) error

// Dispatcher controls interactions with the pool.
type Dispatcher interface {
	// Start initialises the dispatcher.
	Start()
	// Put places the UnitOfWork into the pool.
	// if the context expires before it can be enqueued, error will be ctx.Err().
	Put(context.Context, UnitOfWork) error
	// Close gracefully shuts down the dispatcher once all work is complete.
	Close()
}
