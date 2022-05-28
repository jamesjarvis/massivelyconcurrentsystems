package pool

import (
	"context"
	"sync"
	"testing"
)

type testUnitOfWork struct {
	request  int
	response int
	err      error
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
	panic("not implemented") // TODO: Implement
}

func Test_queue_enqueue(t *testing.T) {
	var bufferSize int = 10
	deadCtx, cancelFunc := context.WithCancel(context.TODO())
	cancelFunc()

	type args struct {
		ctx context.Context
		e   UnitOfWork
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sends successfully",
			args: args{
				ctx: context.TODO(),
				e:   &testUnitOfWork{},
			},
			wantErr: false,
		},
		{
			name: "returns context error on deadline exceeded",
			args: args{
				ctx: deadCtx,
				e:   &testUnitOfWork{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			q := newQueue(bufferSize, &wg)

			if err := q.enqueue(tt.args.ctx, tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("queue.enqueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
		q.enqueue(context.TODO(), &testUnitOfWork{})
		q.enqueue(context.TODO(), &testUnitOfWork{})

		var itemsRetrieved int
		go func() {
			for {
				u, ok := q.dequeue()
				if !ok || u == nil {
					break
				}
				itemsRetrieved = itemsRetrieved + 1
			}
		}()

		q.close()
		if itemsRetrieved != 2 {
			t.Fatal("Incorrect items received while closing queue")
		}
	})
}
