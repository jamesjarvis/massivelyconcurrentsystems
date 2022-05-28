// Package pool is designed to be a simple worker pool orchestrator
package pool

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
	// not sure what do use the error for yet. Maybe rejections? :shrug:
	Put(UnitOfWork) error
	// Close gracefully shuts down the dispatcher once all work is complete.
	Close()
}
