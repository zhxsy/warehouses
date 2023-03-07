package mq

import (
	"go.uber.org/atomic"
	"testing"
)

func TestAtomicBool(t *testing.T) {
	var a = atomic.NewBool(false)
	t.Logf("%v\n", a)
	ok := a.Swap(true)
	t.Logf("%v, %v\n", a, ok)
	ok = a.Swap(true)
	t.Logf("%v, %v\n", a, ok)
	ok = a.Swap(false)
	t.Logf("%v, %v\n", a, ok)
	ok = a.Swap(false)
	t.Logf("%v, %v\n", a, ok)
}
