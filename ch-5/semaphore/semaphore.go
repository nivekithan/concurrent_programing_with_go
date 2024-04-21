package semaphore

import (
	"sync"
)

type Semaphore struct {
	permits int
	cond    *sync.Cond
}

func NewSemaphore(initalPermits int) *Semaphore {
	return &Semaphore{
		permits: initalPermits,
		cond:    sync.NewCond(&sync.Mutex{}),
	}
}

func (s *Semaphore) Release(permit int) {
	s.cond.L.Lock()

	defer s.cond.L.Unlock()

	s.permits += permit
	s.cond.Signal()
}

func (s *Semaphore) Acquire(permit int) {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	for s.permits >= permit {
		s.cond.Wait()
	}

	s.permits -= permit
}
