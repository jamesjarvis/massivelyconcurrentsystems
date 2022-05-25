package deps

import "sync"

type DepHandlerChans struct {
	deps     []DependencyManager
	depsMu   sync.RWMutex
	doneChan chan struct{}
}

func (r *DepHandlerChans) AddDependency(dep DependencyManager) {
	r.depsMu.Lock()
	r.deps = append(r.deps, dep)
	r.depsMu.Unlock()
}

func (r *DepHandlerChans) WaitForDependencies() {
	r.depsMu.RLock()
	for _, dep := range r.deps {
		dep.Wait()
	}
	r.depsMu.RUnlock()
}

func (r *DepHandlerChans) Wait() {
	r.WaitForDependencies()
	<-r.doneChan
}

func (r *DepHandlerChans) Finish() {
	close(r.doneChan)
}
