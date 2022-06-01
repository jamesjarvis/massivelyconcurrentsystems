// Package pool is designed to be a simple worker pool orchestrator,
// primarily focused on workloads that require batched execution of work
// due to I/O limitations.
// e.g.: batching access to a underlying database connection.
// Heavily influenced by github.com/jiacai2050/prosumer
package pool

import (
	"context"
)

// UnitOfWork is the object designed to be passed into, and then operated on, by the pool.
type UnitOfWork[REQ any, RESP any] interface {
	GetRequest() REQ
	GetResponse() RESP
	SetResponse(RESP)
	GetError() error
	SetError(error)
	Done()
}

// BatchWorker is a function that will be used to operate on a batch of requests from the pool.
type BatchWorker[REQ, RESP any] func([]UnitOfWork[REQ, RESP]) error

// Dispatcher controls interactions with the pool.
type Dispatcher[E any] interface {
	// Start initialises the dispatcher.
	Start()
	// Put places the UnitOfWork into the pool.
	// if the context expires before it can be enqueued, error will be ctx.Err().
	Put(context.Context, E) error
	// Close gracefully shuts down the dispatcher once all work is complete.
	Close()
}
