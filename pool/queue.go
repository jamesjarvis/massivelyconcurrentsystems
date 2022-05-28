package pool

import (
	"context"
	"sync"
	"time"
)

// queue handles access into and out of the internal queue.
type queue struct {
	ch chan UnitOfWork

	waitClose *sync.WaitGroup
}

func (q *queue) enqueue(ctx context.Context, e UnitOfWork) error {
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

func (q *queue) dequeue() (UnitOfWork, bool) {
	select {
	case v, ok := <-q.ch:
		return v, ok
	default:
		return nil, false
	}
}

func (q *queue) size() int {
	return len(q.ch)
}

func (q *queue) cap() int {
	return cap(q.ch)
}

func (q *queue) close() {
	close(q.ch)
	for {
		if q.size() == 0 {
			break
		}
		time.Sleep(time.Millisecond)
	}
	q.waitClose.Done()
}

func newQueue(bufferSize int, waitClose *sync.WaitGroup) *queue {
	waitClose.Add(1)
	return &queue{
		ch:        make(chan UnitOfWork, bufferSize),
		waitClose: waitClose,
	}
}
