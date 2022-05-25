package deps

import "sync"

type DepHandlerWaitGroup struct {
	deps   []DependencyManager
	depsMu sync.RWMutex
	wg     sync.WaitGroup
}

func NewDepHandlerWaitGroup() *DepHandlerWaitGroup {
	depHandler := &DepHandlerWaitGroup{
		depsMu: sync.RWMutex{},
		wg:     sync.WaitGroup{},
	}
	depHandler.wg.Add(1)
	return depHandler
}

func (r *DepHandlerWaitGroup) AddDependency(dep DependencyManager) {
	r.depsMu.Lock()
	r.deps = append(r.deps, dep)
	r.depsMu.Unlock()
}

func (r *DepHandlerWaitGroup) WaitForDependencies() {
	r.depsMu.RLock()
	for _, dep := range r.deps {
		dep.Wait()
	}
	r.depsMu.RUnlock()
}

func (r *DepHandlerWaitGroup) Wait() {
	r.WaitForDependencies()
	r.wg.Wait()
}

func (r *DepHandlerWaitGroup) Finish() {
	r.wg.Done()
}
