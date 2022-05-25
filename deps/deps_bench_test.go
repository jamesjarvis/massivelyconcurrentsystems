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
const d = 10 * time.Millisecond

// request is used for testing purposes as an example implementation of DependencyManager usage.
type request struct {
	depmanager DependencyManager
	name       string
	sleeptime  time.Duration
}

// AddDependency adds requests as dependencies of the current request.
func (r *request) AddDependency(deps ...*request) {
	for _, dep := range deps {
		r.depmanager.AddDependency(dep.depmanager)
	}
}

// Run services the request.
func (r *request) Run() {
	defer r.depmanager.Finish()

	// Fake doing some work.
	time.Sleep(r.sleeptime)

	// Cannot finish until it's dependencies have finished.
	r.depmanager.WaitForDependencies()
}

func newRequest(name string, sleeptime time.Duration, deps ...*request) *request {
	depManager := &DepHandlerChans{
		doneChan: make(chan struct{}),
	}
	r := &request{
		name:       name,
		depmanager: depManager,
		sleeptime:  sleeptime,
	}
	r.AddDependency(deps...)
	return r
}

func createRequests(depth, depsPerReq int) (root *request, all []*request) {
	// break out of the loop if there is a depth of 0, returning only one request.
	if depth == 0 {
		return newRequest(fmt.Sprintf("%d_leaf", depth), d), nil
	}

	siblingreqs := make([]*request, 0, int(math.Pow(float64(depsPerReq), float64(depth))))
	for i := 0; i < depsPerReq; i++ {
		req, additionalReqs := createRequests(depth-1, depsPerReq)
		all = append(all, additionalReqs...)

		siblingreqs = append(siblingreqs, req)
	}

	root = newRequest(fmt.Sprintf("%d_root", depth), d, siblingreqs...)
	all = append(all, siblingreqs...)

	return root, all
}

func startRequests(reqs []*request) {
	for _, req := range reqs {
		go req.Run()
	}
}

func BenchmarkWaitManyDepsWithDeps(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		root, reqs := createRequests(L, N)
		// fmt.Println(len(reqs)+1, "requests")

		startRequests(reqs)
		go root.Run()

		root.depmanager.Wait()
	}
}
