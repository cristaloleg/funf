package loglock

import (
	"sync"
)

// LogLock is a logged lock
type LogLock struct {
	mutex sync.RWMutex

	onLock    func()
	onUnlock  func()
	onRLock   func()
	onRUnlock func()
}

// Lock locks for read/write
func (ll *LogLock) Lock() {
	ll.mutex.Lock()
	if ll.onLock != nil {
		ll.onLock()
	}
}

// Unlock unlocks for read/write
func (ll *LogLock) Unlock() {
	ll.mutex.Unlock()
	if ll.onUnlock != nil {
		ll.onUnlock()
	}
}

// RLock locks for reading
func (ll *LogLock) RLock() {
	ll.mutex.RLock()
	if ll.onRLock != nil {
		ll.onRLock()
	}
}

// RUnlock unlocks reading
func (ll *LogLock) RUnlock() {
	ll.mutex.RUnlock()
	if ll.onRUnlock != nil {
		ll.onRUnlock()
	}
}

// OnLock executes on locking
func (ll *LogLock) OnLock(f func()) {
	ll.onLock = f
}

// OnUnlock executes on unlocking
func (ll *LogLock) OnUnlock(f func()) {
	ll.onUnlock = f
}

// OnRLock executes on read locking
func (ll *LogLock) OnRLock(f func()) {
	ll.onRLock = f
}

// OnRUnlock executes on read unlocking
func (ll *LogLock) OnRUnlock(f func()) {
	ll.onRUnlock = f
}
