package writePrefferedReadWriteMutex

import "sync"

type WritePreferredReadWriteMutex struct {
	readersCounter int
	writersWaiting int
	writerActive   bool
	cond           *sync.Cond
}

func NewWritePreferredReadWriteMutex() *WritePreferredReadWriteMutex {

	return &WritePreferredReadWriteMutex{
		writerActive:   false,
		readersCounter: 0,
		writersWaiting: 0,
		cond:           sync.NewCond(&sync.Mutex{}),
	}
}

func (rw *WritePreferredReadWriteMutex) WriteUnlock() {
	rw.cond.L.Lock()
	defer rw.cond.L.Unlock()

	rw.writerActive = false

	rw.cond.Broadcast()
}

func (rw *WritePreferredReadWriteMutex) WriteLock() {
	rw.cond.L.Lock()
	defer rw.cond.L.Unlock()

	rw.writersWaiting++

	for rw.readersCounter != 0 || rw.writerActive {
		rw.cond.Wait()
	}

	rw.writersWaiting--

	rw.writerActive = true
}

func (rw *WritePreferredReadWriteMutex) ReadUnlock() {
	rw.cond.L.Lock()
	defer rw.cond.L.Unlock()

	rw.readersCounter--

	if rw.readersCounter == 0 {
		rw.cond.Broadcast()
	}
}

func (rw *WritePreferredReadWriteMutex) ReadLock() {
	rw.cond.L.Lock()
	defer rw.cond.L.Unlock()

	for rw.writersWaiting != 0 || rw.writerActive {
		rw.cond.Wait()
	}

	rw.readersCounter++
}
