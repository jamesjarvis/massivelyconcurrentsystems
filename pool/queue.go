package pool

import (
	"context"
	"sync"
	"time"
)

// queue handles access into and out of the internal queue.
type queue[REQ, RESP any] struct {
	ch chan UnitOfWork[REQ, RESP]

	waitClose *sync.WaitGroup
}

func (q *queue[REQ, RESP]) enqueue(ctx context.Context, e UnitOfWork[REQ, RESP]) error {
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

func (q *queue[REQ, RESP]) dequeue() (UnitOfWork[REQ, RESP], bool) {
	select {
	case v, ok := <-q.ch:
		return v, ok
	default:
		return nil, false
	}
}

func (q *queue[REQ, RESP]) size() int {
	return len(q.ch)
}

func (q *queue[REQ, RESP]) close() {
	close(q.ch)
	for {
		if q.size() == 0 {
			break
		}
		time.Sleep(time.Millisecond)
	}
	q.waitClose.Done()
}

func newQueue[REQ, RESP any](bufferSize int, waitClose *sync.WaitGroup) *queue[REQ, RESP] {
	waitClose.Add(1)
	return &queue[REQ, RESP]{
		ch:        make(chan UnitOfWork[REQ, RESP], bufferSize),
		waitClose: waitClose,
	}
}
