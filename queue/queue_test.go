package queue

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestInMemoryQueueEnqueue(t *testing.T) {
	const bufferSize = 10

	t.Run("sends successfully", func(t *testing.T) {
		var wg sync.WaitGroup
		q := NewInMemoryQueue[int](bufferSize, &wg)

		if err := q.Enqueue(context.TODO(), 10); err != nil {
			t.Errorf("queue.enqueue() error = %v, wantErr %v", err, nil)
		}
	})
	t.Run("returns context error on deadline exceeded", func(t *testing.T) {
		var wg sync.WaitGroup
		q := NewInMemoryQueue[int](bufferSize, &wg)

		deadCtx, cancelFunc := context.WithDeadline(context.TODO(), time.Now().Add(-time.Second))
		defer cancelFunc()
		<-deadCtx.Done()

		if err := q.Enqueue(deadCtx, 10); err == nil {
			t.Fatalf("queue.enqueue() error = %v, wantErr %v", err, context.Canceled)
		}
	})

}

func TestQueueGracefulClose(t *testing.T) {
	bufferSize := 2

	t.Run("empty queue closes instantly", func(t *testing.T) {
		var wg sync.WaitGroup
		q := NewInMemoryQueue[string](bufferSize, &wg)
		q.Close()
	})
	t.Run("queue with items waits for items to leave", func(t *testing.T) {
		var wg sync.WaitGroup
		q := NewInMemoryQueue[string](bufferSize, &wg)
		err := q.Enqueue(context.TODO(), "hello world")
		if err != nil {
			t.Error(err)
		}
		err = q.Enqueue(context.TODO(), "hello world")
		if err != nil {
			t.Error(err)
		}

		var itemsRetrieved int
		go func() {
			for {
				u, ok := q.Dequeue()
				if !ok || u == "" {
					break
				}
				itemsRetrieved++
			}
		}()

		q.Close()
		if itemsRetrieved != 2 {
			t.Errorf("Incorrect items received while closing queue")
		}
	})
}
