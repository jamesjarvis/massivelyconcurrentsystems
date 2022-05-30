package pool

import (
	"context"
	"sync"
	"testing"
	"time"
)

type testUnitOfWork struct {
	wg *sync.WaitGroup
}

func (u *testUnitOfWork) GetRequest() any {
	panic("not implemented") // TODO: Implement
}

func (u *testUnitOfWork) GetResponse() any {
	panic("not implemented") // TODO: Implement
}

func (u *testUnitOfWork) SetResponse(_ any) {
	panic("not implemented") // TODO: Implement
}

func (u *testUnitOfWork) GetError() error {
	panic("not implemented") // TODO: Implement
}

func (u *testUnitOfWork) SetError(_ error) {
	panic("not implemented") // TODO: Implement
}

func (u *testUnitOfWork) Done() {
	u.wg.Done()
}

func Test_queue_enqueue(t *testing.T) {
	const bufferSize = 10

	t.Run("sends successfully", func(t *testing.T) {
		var wg sync.WaitGroup
		q := newQueue(bufferSize, &wg)

		if err := q.enqueue(context.TODO(), &testUnitOfWork{}); err != nil {
			t.Errorf("queue.enqueue() error = %v, wantErr %v", err, nil)
		}
	})
	t.Run("returns context error on deadline exceeded", func(t *testing.T) {
		var wg sync.WaitGroup
		q := newQueue(bufferSize, &wg)

		deadCtx, cancelFunc := context.WithDeadline(context.TODO(), time.Now().Add(-time.Second))
		defer cancelFunc()
		<-deadCtx.Done()

		if err := q.enqueue(deadCtx, &testUnitOfWork{}); err == nil {
			t.Fatalf("queue.enqueue() error = %v, wantErr %v", err, context.Canceled)
		}
	})

}

func TestQueueGracefulClose(t *testing.T) {
	bufferSize := 2

	t.Run("empty queue closes instantly", func(t *testing.T) {
		var wg sync.WaitGroup
		q := newQueue(bufferSize, &wg)
		q.close()
	})
	t.Run("queue with items waits for items to leave", func(t *testing.T) {
		var wg sync.WaitGroup
		q := newQueue(bufferSize, &wg)
		err := q.enqueue(context.TODO(), &testUnitOfWork{})
		if err != nil {
			t.Error(err)
		}
		err = q.enqueue(context.TODO(), &testUnitOfWork{})
		if err != nil {
			t.Error(err)
		}

		var itemsRetrieved int
		go func() {
			for {
				u, ok := q.dequeue()
				if !ok || u == nil {
					break
				}
				itemsRetrieved++
			}
		}()

		q.close()
		if itemsRetrieved != 2 {
			t.Errorf("Incorrect items received while closing queue")
		}
	})
}
