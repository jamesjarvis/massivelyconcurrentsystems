package pool

import (
	"context"
	"sync"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestBatchDispatcher(t *testing.T) {
	tests := []struct {
		name                  string
		expectedWorkProcessed int
		opts                  []Opt
	}{
		{
			name:                  "no work",
			expectedWorkProcessed: 0,
		},
		{
			name:                  "lots of work",
			expectedWorkProcessed: 50001,
		},
		{
			name:                  "more work than buffer",
			expectedWorkProcessed: 10000,
			opts:                  []Opt{SetBufferSize(10)},
		},
		{
			name:                  "more workers than work",
			expectedWorkProcessed: 10000,
			opts:                  []Opt{SetNumConsumers(1001)},
		},
		{
			name:                  "tiny batch size",
			expectedWorkProcessed: 10000,
			opts:                  []Opt{SetBatchSize(1)},
		},
		{
			name:                  "huge interval",
			expectedWorkProcessed: 10000,
			opts:                  []Opt{SetBatchInterval(time.Hour)},
		},
		{
			name:                  "teeny interval",
			expectedWorkProcessed: 10000,
			opts:                  []Opt{SetBatchInterval(time.Nanosecond)},
		},
		{
			name:                  "teeny interval and many workers",
			expectedWorkProcessed: 10000,
			opts:                  []Opt{SetBatchInterval(time.Nanosecond), SetNumConsumers(100)},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer goleak.VerifyNone(t)

			var mu sync.Mutex
			var workProcessed int
			var calls int

			worker := func(us []UnitOfWork) error {
				mu.Lock()
				defer mu.Unlock()
				workProcessed = workProcessed + len(us)
				calls = calls + 1
				return nil
			}
			config := NewConfig(worker)
			dispatcher := NewBatchDispatcher(config)

			dispatcher.Start()

			for i := 0; i < test.expectedWorkProcessed; i++ {
				err := dispatcher.Put(context.TODO(), &testUnitOfWork{})
				if err != nil {
					t.Errorf("dispatcher.Put error encountered! %v != nil", err)
				}
			}

			err := dispatcher.Close()
			if err != nil {
				t.Errorf("dispatcher.Close should not error! %v != nil", err)
			}

			if workProcessed != test.expectedWorkProcessed {
				t.Errorf("workProcessed invalid! %d != %d", workProcessed, test.expectedWorkProcessed)
			}

			if calls > 0 {
				t.Logf("worker called %d times with avg batch size of %d", calls, workProcessed/calls)
			}
		})
	}
}
