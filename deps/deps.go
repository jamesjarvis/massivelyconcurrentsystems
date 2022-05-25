package deps

// DependencyManager handles
type DependencyManager interface {
	// AddDependency adds a dependency to the DependencyManager.
	AddDependency(DependencyManager)
	// WaitForDependencies waits until all dependencies have completed.
	WaitForDependencies()
	// Wait waits until the dependencymanager has completed.
	Wait()
	// Finish completes the DependencyManager, thereby unblocking Wait.
	// Finish is designed to be called once, and ideally last.
	Finish()
}
