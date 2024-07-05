package counters

import (
	"testing"
)

func TestSimpleCounter_Value(t *testing.T) {

	var count Simple
	if count.Value() != 0 {
		t.Fatal("initial state")
	}

	if count.Increment() != 0 {
		t.Fatal("expect 0")
	}

	if count.Value() != 1 {
		t.Fatal("expect 1")
	}
}
