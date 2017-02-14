package funf

import "sync"

type WorkersPool struct {
	wg         sync.WaitGroup
	atomicSize int
	fn         func(interface{}) interface{}
	data       <-chan interface{}
	result     chan interface{}
}

func NewWorkersPool(fn func(interface{}) interface{}, data <-chan interface{}) *WorkersPool {
	return &WorkersPool{
		fn:     fn,
		data:   data,
		result: make(chan interface{}),
	}
}

func (w *WorkersPool) Results() <-chan interface{} {
	return w.result
}

func (w *WorkersPool) Add(delta int) {
	if delta <= 0 {
		return
	}
	size := w.atomicSize
	for i := 1; i < delta; i++ {
		w.invoke(i + size)
	}
}

func (w *WorkersPool) Rem(delta int) {
	if delta <= 0 {
		return
	}
	w.wg.Add(delta)
	w.atomicSize -= delta
}

func (w *WorkersPool) Wait() {
	w.wg.Wait()
}

func (w *WorkersPool) invoke(id int) {
	w.wg.Add(1)
	w.atomicSize++
	go func(id int) {
		for d := range w.data {
			w.result <- w.fn(d)

			if w.atomicSize < id {
				w.wg.Done()
			}
		}
	}(id)
}
