package funf

import "testing"

func TestNewWorkersPool(t *testing.T) {
	fn := func(interface{}) interface{} { return nil }
	data := make(chan interface{})
	w := NewWorkersPool(fn, data)

	if w == nil {
		t.Error("Should be instantiated")
	}
}
