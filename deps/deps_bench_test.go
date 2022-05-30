package deps

import (
	"fmt"
	"math"
	"testing"
	"time"
)

// This benchmarking suite will test the performance of dephandlerwaitgroup vs dephandlerchans.
// The scenario will be a request that spawns L layers of dependencies, where each layer has N dependencies each.
const L, N = 4, 10
const d = 1 * time.Millisecond

type dephandlerflavour int

const (
	channel dephandlerflavour = iota
	waitgroup
)

// request is used for testing purposes as an example implementation of DependencyManager usage.
type request struct {
	depmanager DependencyManager
	name       string
	sleeptime  time.Duration
}

// AddDependency adds requests as dependencies of the current request.
func (r request) AddDependency(deps ...request) {
	for _, dep := range deps {
		r.depmanager.AddDependency(dep.depmanager)
	}
}

// Run services the request.
func (r request) Run() {
	defer r.depmanager.Finish()

	// switching things around, wait for dependencies first, then sleep.
	r.depmanager.WaitForDependencies()

	// Fake doing some work.
	time.Sleep(r.sleeptime)
}

// newRequest creates a new request, using a different dependency handler depending on an enum.
func newRequest(name string, sleeptime time.Duration, flavour dephandlerflavour, deps ...request) request {
	var depManager DependencyManager
	switch flavour {
	case channel:
		depManager = NewDepHandlerChans()
	case waitgroup:
		depManager = NewDepHandlerWaitGroup()
	}
	r := request{
		name:       name,
		depmanager: depManager,
		sleeptime:  sleeptime,
	}
	r.AddDependency(deps...)
	return r
}

func createRequests(depth, depsPerReq int, flavour dephandlerflavour) (root request, all []request) {
	// break out of the loop if there is a depth of 0, returning only one request.
	if depth == 0 {
		return newRequest(fmt.Sprintf("%d_leaf", depth), d, flavour), nil
	}

	all = make([]request, 0, int(math.Pow(float64(depsPerReq), float64(depth))))

	siblingreqs := make([]request, 0, int(math.Pow(float64(depsPerReq), float64(depth))))
	for i := 0; i < depsPerReq; i++ {
		req, additionalReqs := createRequests(depth-1, depsPerReq, flavour)
		all = append(all, additionalReqs...)

		siblingreqs = append(siblingreqs, req)
	}

	root = newRequest(fmt.Sprintf("%d_root", depth), d, flavour, siblingreqs...)
	all = append(all, siblingreqs...)

	return root, all
}

// worker runs requests from a channel input. Will die when the channel is closed.
func worker(reqs <-chan request) {
	for r := range reqs {
		r.Run()
	}
}

func BenchmarkWaitManyDepsWithDeps(b *testing.B) {
	// spin up pool of workers
	// note: if the work inside each worker is significant (but non blocking),
	// then it is better to have ~10k workers. Whereas if the work is blocking, or insignificant
	// then it is better to have ~1k workers.
	const numWorkers = 10000
	reqChan := make(chan request, 20000)
	for w := 0; w < numWorkers; w++ {
		go worker(reqChan)
	}
	defer close(reqChan)

	fmt.Printf("Executing %d requests per iteration. Should take %s ideally\n", int(math.Pow(N, L)), (L+1)*d)

	tests := []struct {
		name    string
		flavour dephandlerflavour
	}{
		{
			name:    "chans",
			flavour: channel,
		},
		{
			name:    "waitgroups",
			flavour: waitgroup,
		},
	}
	for _, test := range tests {
		testLocal := test
		b.Run(test.name, func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for n := 0; n < b.N; n++ {
				b.StopTimer() // we don't care about the time it takes to create these request objects.
				root, reqs := createRequests(L, N, testLocal.flavour)
				b.StartTimer()
				// fmt.Println(len(reqs)+1, "requests")

				reqChan <- root
				for _, req := range reqs {
					reqChan <- req // send to pool of workers already waiting
				}

				root.depmanager.Wait()
			}
		})
	}
}
