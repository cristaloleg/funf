package lechan_test

import (
	"testing"

	"github.com/cristaloleg/funf/lechan"
)

func TestLechan_New(t *testing.T) {
	ch := lechan.New(10)
	if ch == nil {
		t.Fatal("cannot instantiate Lechan")
	}
}
