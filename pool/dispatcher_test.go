package pool

import (
	"context"
	"runtime"
	"sync"
	"testing"
	"time"

	"go.uber.org/goleak"
)

const defaultWork = 100

var tests = []struct {
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
		opts:                  []Opt{SetBufferSize(defaultWork / 10)},
	},
	{
		name:                  "more workers than work",
		expectedWorkProcessed: defaultWork,
		opts:                  []Opt{SetNumConsumers(defaultWork * 2)},
	},
	{
		name:                  "tiny batch size",
		expectedWorkProcessed: defaultWork,
		opts:                  []Opt{SetBatchSize(1)},
	},
	{
		name:                  "huge batch size",
		expectedWorkProcessed: defaultWork,
		opts:                  []Opt{SetBatchSize(defaultWork * 10)},
	},
	// {
	// 	name:                  "huge interval",
	// 	expectedWorkProcessed: defaultWork,
	// 	opts:                  []Opt{SetBatchInterval(time.Hour)},
	// },
	{
		name:                  "teeny interval",
		expectedWorkProcessed: defaultWork,
		opts:                  []Opt{SetBatchInterval(time.Nanosecond)},
	},
	// {
	// 	name:                  "teeny interval and many workers (V INCONSISTENT WITH MORE WORKERS THAN WORK)",
	// 	expectedWorkProcessed: defaultWork,
	// 	opts:                  []Opt{SetBatchInterval(time.Nanosecond), SetNumConsumers(defaultWork * 2)},
	// },
	{
		name:                  "slow worker",
		expectedWorkProcessed: defaultWork,
		workerDuration:        5 * time.Millisecond,
	},
	{
		name:                  "slow worker and huge batch size",
		expectedWorkProcessed: defaultWork,
		workerDuration:        5 * time.Millisecond,
		opts:                  []Opt{SetBatchSize(defaultWork)},
	},
}

func TestBatchDispatcher(t *testing.T) {
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
				t.Logf("avg latency per request of %s", time.Since(startTime)/time.Duration(workProcessed))
			}
		})
	}
}

func BenchmarkBatchDispatcher(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			var wg sync.WaitGroup

			worker := func(us []UnitOfWork) error {
				time.Sleep(test.workerDuration)
				wg.Add(-len(us))
				return nil
			}
			config := NewConfig(worker, test.opts...)
			dispatcher := NewBatchDispatcher(config)

			dispatcher.Start()

			b.ReportAllocs()
			b.ResetTimer()

			for n := 0; n < b.N; n++ {
				wg.Add(test.expectedWorkProcessed)

				for i := 0; i < test.expectedWorkProcessed; i++ {
					err := dispatcher.Put(context.TODO(), &testUnitOfWork{})
					if err != nil {
						b.Errorf("dispatcher.Put error encountered! %v != nil", err)
					}
				}

				wg.Wait()
			}

			b.StopTimer()
			err := dispatcher.Close()
			if err != nil {
				b.Errorf("dispatcher.Close should not error! %v != nil", err)
			}
		})
	}
}

func BenchmarkBatchDispatcherSingleItem(b *testing.B) {
	defaultDuration := time.Millisecond

	configTests := []struct {
		name           string
		workerDuration time.Duration
		opts           []Opt
	}{
		{
			name:           "default",
			workerDuration: defaultDuration,
		},
		{
			name:           "tiny buffer",
			workerDuration: defaultDuration,
			opts:           []Opt{SetBufferSize(1)},
		},
		{
			name:           "large buffer",
			workerDuration: defaultDuration,
			opts:           []Opt{SetBufferSize(1000)},
		},
		{
			name:           "tiny interval",
			workerDuration: defaultDuration,
			opts:           []Opt{SetBatchInterval(time.Nanosecond)},
		},
		{
			name:           "huge interval",
			workerDuration: defaultDuration,
			opts:           []Opt{SetBatchInterval(10 * time.Millisecond)},
		},
		{
			name:           "tiny batch size",
			workerDuration: defaultDuration,
			opts:           []Opt{SetBatchSize(1)},
		},
		{
			name:           "huge batch size",
			workerDuration: defaultDuration,
			opts:           []Opt{SetBatchSize(1000)},
		},
		{
			name:           "tiny number of workers",
			workerDuration: defaultDuration,
			opts:           []Opt{SetNumConsumers(1)},
		},
		{
			name:           "educated guess number of workers",
			workerDuration: defaultDuration,
			opts:           []Opt{SetNumConsumers(runtime.NumCPU() / 2)},
		},
		{
			name:           "huge number of workers",
			workerDuration: defaultDuration,
			opts:           []Opt{SetNumConsumers(1000)},
		},
		{
			name:           "guesstimate",
			workerDuration: defaultDuration,
			opts:           []Opt{SetNumConsumers(runtime.NumCPU() / 2), SetBatchInterval(time.Nanosecond), SetBatchSize(200), SetBufferSize((runtime.NumCPU() / 2) * 200)},
		},
	}
	for _, test := range configTests {
		b.Run(test.name, func(b *testing.B) {
			worker := func(us []UnitOfWork) error {
				time.Sleep(test.workerDuration)
				for _, u := range us {
					u.Done()
				}
				return nil
			}
			config := NewConfig(worker, test.opts...)
			dispatcher := NewBatchDispatcher(config)

			dispatcher.Start()

			b.ReportAllocs()
			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					wg := &sync.WaitGroup{}
					wg.Add(1)

					err := dispatcher.Put(context.TODO(), &testUnitOfWork{wg: wg})
					if err != nil {
						b.Errorf("dispatcher.Put error encountered! %v != nil", err)
					}

					wg.Wait()
				}
			})

			b.StopTimer()
			err := dispatcher.Close()
			if err != nil {
				b.Errorf("dispatcher.Close should not error! %v != nil", err)
			}
		})
	}
}
