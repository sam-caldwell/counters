package counters

import (
	"testing"
)

func TestSimpleCounter_Increment(t *testing.T) {
	var count Simple
	if count.value != 0 {
		t.Fatal("initial state")
	}

	if count.Increment() != 0 {
		t.Fatal("expect 0")
	}

	if count.Increment() != 1 {
		t.Fatal("expect 1")
	}

	if count.Increment() != 2 {
		t.Fatal("expect 2")
	}
}
