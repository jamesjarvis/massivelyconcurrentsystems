package queue

import (
	"context"
)

// Queue is a FIFO structure for elements of type E.
type Queue[E any] interface {
	Enqueue(context.Context, E) error
	Dequeue() (E, bool)
	DequeueBlocking() <-chan E
	Size() int
	Close()
}
