package waitgroup

import (
	"github.com/nivekithan/concurrent_programming_with_go/ch-5/semaphore"
)

type SemaphoreWaitGroup struct {
	semaphore *semaphore.Semaphore
}

func NewSemaphoreWaitGroup(size int) *SemaphoreWaitGroup {
	semaphore := semaphore.NewSemaphore(1 - size)

	return &SemaphoreWaitGroup{semaphore: semaphore}
}

func (wg *SemaphoreWaitGroup) Done() {
	wg.semaphore.Release(1)
}

func (wg *SemaphoreWaitGroup) Wait() {
	wg.semaphore.Acquire(1)
}
