package barrier

import "sync"

type Barrier struct {
	taskSize       int
	completedTasks int
	cond           *sync.Cond
}

func NewBarrier(taskSize int) *Barrier {
	return &Barrier{
		taskSize: taskSize, completedTasks: 0,
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (b *Barrier) Wait() {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	b.completedTasks++

	if b.completedTasks == b.taskSize {
		b.completedTasks = 0
		b.cond.Broadcast()
	}

	b.cond.Wait()
}
