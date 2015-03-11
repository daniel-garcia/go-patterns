package example

import (
	"sync"
)

type Example struct {
	name    string
	closing chan struct{}
	wg      sync.WaitGroup
}

// New creates a new example with the given name.
func New(name string) *Example {
	s := &Example{
		name:    name,
		closing: make(chan struct{}),
		wg:      sync.WaitGroup{},
	}
	s.wg.Add(1)
	go s.run()
	return s
}

// Name is the name of the of the example.
func (s *Example) Name() string {
	return s.name
}

// Close shuts down the example. This method will panic if
// called twice.
func (s *Example) Close() {
	close(s.closing)
	s.wg.Wait()
}

// run the the body of the example. It is responsible to reacting
// to the changes in the application, model, and infrastructure.
func (s *Example) run() {
	defer s.wg.Done()
	select {
	case <-s.closing:
		return
	}
}
