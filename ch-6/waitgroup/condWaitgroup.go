package waitgroup

import "sync"

type CondWaitGroup struct {
	cond  *sync.Cond
	tasks int
}

func NewCondWaitGroup() *CondWaitGroup {
	return &CondWaitGroup{
		cond:  sync.NewCond(&sync.Mutex{}),
		tasks: 0,
	}
}

func (wg *CondWaitGroup) Add(size int) {
	wg.cond.L.Lock()
	defer wg.cond.L.Unlock()

	wg.tasks++
}

func (wg *CondWaitGroup) Done(size int) {
	wg.cond.L.Lock()
	defer wg.cond.L.Unlock()

	wg.tasks--
	// If there are multiple go routines calling Wait. It will awaken
	// all those goroutines
	if wg.tasks == 0 {
		wg.cond.Broadcast()
	}
}

func (wg *CondWaitGroup) Wait() {
	wg.cond.L.Lock()
	defer wg.cond.L.Unlock()

	for wg.tasks > 0 {
		wg.Wait()
	}
}
