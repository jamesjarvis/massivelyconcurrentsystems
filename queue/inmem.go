package queue

import (
	"context"
	"sync"
	"time"
)

// InMemoryQueue handles access into and out of an in memory queue.
type InMemoryQueue[E any] struct {
	ch chan E

	waitClose *sync.WaitGroup
}

// Check InMemoryQueue satisfies Queue interface.
var _ Queue[int] = &InMemoryQueue[int]{}

// NewInMemoryQueue returns a new queue backed by volatile memory.
func NewInMemoryQueue[E any](bufferSize int, waitClose *sync.WaitGroup) *InMemoryQueue[E] {
	waitClose.Add(1)
	return &InMemoryQueue[E]{
		ch:        make(chan E, bufferSize),
		waitClose: waitClose,
	}
}

// Enqueue appends the element to the end of the queue, returns an error
// if the context expires whilst waiting to append.
func (q *InMemoryQueue[E]) Enqueue(ctx context.Context, e E) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case q.ch <- e:
		return nil
	}
}

// Dequeue retrieves an element from the queue, and a bool indicating whether
// or not there were elements to return.
func (q *InMemoryQueue[E]) Dequeue() (E, bool) {
	select {
	case v, ok := <-q.ch:
		return v, ok
	default:
		var v E
		return v, false
	}
}

// DequeueBlocking returns a channel that can be used to retrieve an element
// from the queue.
func (q *InMemoryQueue[E]) DequeueBlocking() <-chan E {
	return q.ch
}

// Size returns the current size of the queue.
func (q *InMemoryQueue[E]) Size() int {
	return len(q.ch)
}

// Close closes the queue, preventing any further Enqueue operations.
// Blocks until the queue is empty.
func (q *InMemoryQueue[E]) Close() error {
	close(q.ch)
	for {
		if q.Size() == 0 {
			break
		}
		time.Sleep(time.Millisecond)
	}
	q.waitClose.Done()
	return nil
}
