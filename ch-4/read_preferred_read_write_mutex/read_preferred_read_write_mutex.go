package readPrefferedReadWriteMutex

import "sync"

// Implements these methods
// 1. ReadLock() - Allows multiple readers to read the data
// 2. ReadUnlock() - Releases the Read lock
// 3. WriteLock() - Allows only one writer to write the data
// 4. WriteUnlock() - Releases the Write lock
// 5. TryReadLock() - Tries to acquire the read lock (Excersie 3)
// 6. TryWriteLock() - Tries to acquire the write lock (Excersie 2)

type ReadPreferredReadWriteMutex struct {
	readersCounter int
	readersLock    sync.Mutex
	globalLock     sync.Mutex
}

func NewReadPrefferedReadWriteMutex() *ReadPreferredReadWriteMutex {
	return &ReadPreferredReadWriteMutex{
		readersCounter: 0,
	}
}

func (mu *ReadPreferredReadWriteMutex) ReadLock() {
	mu.readersLock.Lock() // Prevents any other reader from entering/leaving
	defer mu.readersLock.Unlock()

	mu.readersCounter++ // Increment the number of readers

	isFirstToEnter := mu.readersCounter == 1

	if isFirstToEnter {
		mu.globalLock.Lock() // Prevent any writers to enter
	}

}

func (mu *ReadPreferredReadWriteMutex) ReadUnlock() {
	mu.readersLock.Lock() // Prevent anyone from enter/leave
	defer mu.readersLock.Unlock()
	mu.readersCounter--

	isLastToLeave := mu.readersCounter == 0

	if isLastToLeave {
		mu.globalLock.Unlock() // Allow writers to enter
	}

}

func (mu *ReadPreferredReadWriteMutex) TryReadLock() bool {

	isReaderLockFree := mu.readersLock.TryLock()

	if !isReaderLockFree {
		return false
	}

	defer mu.readersLock.Unlock()

	isReadersPresent := mu.readersCounter >= 1

	if !isReadersPresent && !mu.globalLock.TryLock() {
		// There is no readers present and we are unable to get the global lock. It means there
		// are writers holding this lock
		return false
	}

	mu.readersCounter++

	return true
}

func (mu *ReadPreferredReadWriteMutex) WriteLock() {
	// Lock will be acquired only when there is no readers on room
	mu.globalLock.Lock()
}

func (mu *ReadPreferredReadWriteMutex) WriteUnlock() {
	mu.globalLock.Unlock()
}

func (mu *ReadPreferredReadWriteMutex) TryWriteLock() bool {
	return mu.globalLock.TryLock()
}
