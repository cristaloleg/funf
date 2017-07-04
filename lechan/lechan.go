package lechan

import (
	"time"

	"go.uber.org/atomic"
)

type Channel interface {
	Put(interface{})
	TryPut(interface{}) bool
	TryTimedPut(interface{}, time.Duration) bool
	TryCountedPut(interface{}, time.Duration, int) bool
	TryFallbackPut(interface{}, func(interface{})) bool
	Get() (interface{}, bool)
}

type Lechan struct {
	size  uint32
	items atomic.Uint32
	queue chan interface{}
}

func New(size uint32) *Lechan {
	ch := &Lechan{
		size:  size,
		queue: make(chan interface{}, size),
	}
	return ch
}

func (ch *Lechan) Put(value interface{}) {
	ch.queue <- value
	ch.items.Inc()
}

func (ch *Lechan) TryPut(value interface{}) bool {
	if ch.items.Load() == ch.size {
		return false
	}
	ch.queue <- value
	ch.items.Inc()
	return true
}

func (ch *Lechan) TryTimedPut(value interface{}, duration time.Duration) bool {
	for {
		select {
		case <-time.After(duration):
			return false
		default:
			if ch.TryPut(value) {
				return true
			}
		}
	}
}

func (ch *Lechan) TryCountedPut(value interface{}, pause time.Duration, count int) bool {
	for i := 0; i < count; i++ {
		if ch.TryPut(value) {
			return true
		}
	}
	return false
}

func (ch *Lechan) TryFallbackPut(value interface{}, f func(interface{})) bool {
	if ch.TryPut(value) {
		return true
	}
	f(value)
	return false
}

func (ch *Lechan) Get() (value interface{}, ok bool) {
	if ch.items.Load() == 0 {
		return nil, false
	}
	ch.items.Dec()
	return <-ch.queue, true
}
