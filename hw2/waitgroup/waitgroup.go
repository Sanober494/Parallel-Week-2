package waitgroup

// From the Go Documentation: A WaitGroup waits for a collection of
// goroutines to finish. The main goroutine calls Add to set the number
// of goroutines to wait for.  Then each of the goroutines runs and calls
// Done when finished.  At the same time, Wait can be used to block until
// all goroutines have finished.
type WaitGroup interface {
	Add(amount uint)
	Done()
	Wait()
}

// NewWaitGroup returns a instance of a waitgroup
// This instance must be a pointer and should not
// be copied after creation.
func NewWaitGroup() WaitGroup {
	//IMPLEMENT ME!
}

