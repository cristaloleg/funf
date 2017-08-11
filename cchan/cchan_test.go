package cchan

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	in := make(chan int, 10)
	out := make(chan float32, 10)

	for i := 0; i < 10; i++ {
		in <- i
	}
	for i := 0; i < 10; i++ {
		out <- float32(math.Sqrt(float64(i)))
	}

	f := func(i int) float32 { return float32(math.Sqrt(float64(i))) }
	tmp, err := Convert(in, f)

	if err != nil {
		t.Error("err must be nil")
	}

	ch, ok := tmp.Interface().(chan float32)
	if !ok {
		fmt.Printf("%+v\n", reflect.ValueOf(tmp))
		fmt.Printf("%#v\n", reflect.ValueOf(out))
		t.Error("cannot convert")
		t.Error(reflect.ValueOf(tmp))
	}

	if len(ch) != 0 {
		t.Error("size isn't 10", len(ch))
	}

	for i := 0; i < 10; i++ {
		expectd, ok := <-out
		actual, ok2 := <-ch

		if !ok || !ok2 {
			t.Error("empty")
		}
		if expectd != actual {
			t.Error("not equal", expectd, actual)
		}
	}
}
