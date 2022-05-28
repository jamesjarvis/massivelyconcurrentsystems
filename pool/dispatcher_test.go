package pool

import (
	"context"
	"sync"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestBatchDispatcher(t *testing.T) {
	const defaultWork = 1000

	tests := []struct {
		name                  string
		expectedWorkProcessed int
		workerDuration        time.Duration
		opts                  []Opt
	}{
		{
			name:                  "no work",
			expectedWorkProcessed: 0,
		},
		{
			name:                  "lots of work",
			expectedWorkProcessed: defaultWork * 10,
		},
		{
			name:                  "more work than buffer",
			expectedWorkProcessed: defaultWork,
			opts:                  []Opt{SetBufferSize(10)},
		},
		{
			name:                  "more workers than work",
			expectedWorkProcessed: defaultWork,
			opts:                  []Opt{SetNumConsumers(1001)},
		},
		{
			name:                  "tiny batch size",
			expectedWorkProcessed: defaultWork,
			opts:                  []Opt{SetBatchSize(1)},
		},
		{
			name:                  "huge batch size",
			expectedWorkProcessed: defaultWork,
			opts:                  []Opt{SetBatchSize(1000)},
		},
		{
			name:                  "huge interval",
			expectedWorkProcessed: defaultWork,
			opts:                  []Opt{SetBatchInterval(time.Hour)},
		},
		{
			name:                  "teeny interval",
			expectedWorkProcessed: defaultWork,
			opts:                  []Opt{SetBatchInterval(time.Nanosecond)},
		},
		{
			name:                  "teeny interval and many workers",
			expectedWorkProcessed: defaultWork,
			opts:                  []Opt{SetBatchInterval(time.Nanosecond), SetNumConsumers(100)},
		},
		{
			name:                  "slow worker",
			expectedWorkProcessed: defaultWork,
			workerDuration:        time.Millisecond,
		},
		{
			name:                  "slow worker and huge batch size",
			expectedWorkProcessed: defaultWork,
			workerDuration:        5 * time.Millisecond,
			opts:                  []Opt{SetBatchSize(1000)},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer goleak.VerifyNone(t)

			var mu sync.Mutex
			var workProcessed int
			var calls int

			worker := func(us []UnitOfWork) error {
				time.Sleep(test.workerDuration)
				mu.Lock()
				defer mu.Unlock()
				workProcessed = workProcessed + len(us)
				calls = calls + 1
				return nil
			}
			config := NewConfig(worker, test.opts...)
			dispatcher := NewBatchDispatcher(config)

			dispatcher.Start()
			startTime := time.Now()

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
				t.Logf("workers called %d times with avg batch size of %d in %s", calls, workProcessed/calls, time.Since(startTime))
			}
		})
	}
}
