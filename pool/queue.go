package pool

import (
	"context"
	"sync"
	"time"
)

// queue handles access into and out of the internal queue.
type queue[E any] struct {
	ch chan E

	waitClose *sync.WaitGroup
}

func (q *queue[E]) enqueue(ctx context.Context, e E) error {
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

func (q *queue[E]) dequeue() (E, bool) {
	select {
	case v, ok := <-q.ch:
		return v, ok
	default:
		var v E
		return v, false
	}
}

func (q *queue[E]) size() int {
	return len(q.ch)
}

func (q *queue[E]) close() {
	close(q.ch)
	for {
		if q.size() == 0 {
			break
		}
		time.Sleep(time.Millisecond)
	}
	q.waitClose.Done()
}

func newQueue[E any](bufferSize int, waitClose *sync.WaitGroup) *queue[E] {
	waitClose.Add(1)
	return &queue[E]{
		ch:        make(chan E, bufferSize),
		waitClose: waitClose,
	}
}
