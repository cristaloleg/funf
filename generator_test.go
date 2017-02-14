package funf

import "testing"

func TestNewGenerator(t *testing.T) {
	fn := func(int64) interface{} { return nil }
	g := NewGenerator(fn)

	if g == nil {
		t.Error("Should be instantiated")
	}
}

func TestNewBoundedGenerator(t *testing.T) {
	size := int64(10)
	fn := func(x int64) interface{} { return x }
	g := NewBoundedGenerator(size, fn)

	if g == nil {
		t.Error("Should be instantiated")
	}

	for i := int64(0); i < size; i++ {
		value := <-g.Next()
		if value != i {
			t.Error("Should equal to index, was ", value, " instead of ", i)
		}
	}
}
